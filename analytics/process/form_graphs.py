import datetime as dt
from analytics import analytics_pb2, analytics_pb2_grpc
from process.const import CODE_NAME, MULTYPLIER_NAME
import math
from process.get_report_summary import form_report_description
from process.load_json import load_json

valid_min_year = 2009
valid_max_year = 2023
availableColors = [
    "#2563eb", # Blue
    "#60a5fa", # Light Blue
    "#DA3832", # Red
    "#34D399", # Green
    "#FBBF24", # Yellow
    "#A78BFA", # Purple
    "#F87171", # Light Red
    "#10B981", # Emerald
    "#F59E0B", # Amber
    "#8B5CF6", # Violet
]

def get_one_param_records(
    records: dict, code: int, start_date: str, end_date: str, return_year=True
) -> list[analytics_pb2.Record]:
    record_points = []
    start_date = int(start_date)
    end_date = int(end_date)
    for year in range(start_date, end_date + 1, 1):
        value = records[str(code)].get(str(year), None)
        if value is not None and not math.isnan(value):
            if return_year:
                x = str(year)
            else:
                x = str(CODE_NAME[int(code)])
            record_points.append(analytics_pb2.Record(x=x, y=value))
    return record_points


def get_one_multipl_records(
    records: dict, multy_key: str, start_date: str, end_date: str, return_year=True
) -> list[analytics_pb2.Record]:
    record_points = []
    start_date = int(start_date)
    end_date = int(end_date)
    for year in range(start_date, end_date + 1, 1):
        value = records[str(multy_key)].get(str(year), None)
        if value is not None and not math.isnan(value):
            if return_year:
                x = str(year)
            else:
                x = str(MULTYPLIER_NAME[multy_key])
            record_points.append(analytics_pb2.Record(x=x, y=value))
    return record_points


def count_records_percentage(
    records: list[analytics_pb2.Record],
) -> list[analytics_pb2.Record]:
    valid_records = [rec for rec in records if not math.isnan(rec.y)]
    percent_records = []
    whole_sum = 0
    for record in valid_records:
        whole_sum += record.y

    for record in valid_records:
        percent_records.append(
            analytics_pb2.Record(x=record.x, y=(record.y / whole_sum) * 100)
        )
    return percent_records

def form_group_chart(records: dict, codes: list[int], start_date: int, end_date: int, 
                    chart_type: analytics_pb2.ChartType) -> analytics_pb2.ListChartsLegend:
    legend_element = {}
    charts: list[analytics_pb2.Chart] = []
    for i, code in enumerate(codes):
        color = availableColors[i % len(availableColors)]
        recs = get_one_param_records(records, code, start_date, end_date)
        if len(recs) > 0:
            if chart_type == analytics_pb2.PIE_CHART:
                new_chart = analytics_pb2.Chart(
                    color=color,
                    type=chart_type,
                    records=count_records_percentage(recs)
                )
            else:
                new_chart = analytics_pb2.Chart(
                    color=color,
                    type=chart_type,
                    records=recs
                )
            legend_element[color] = CODE_NAME[int(code)]
            charts.append(new_chart)

    chart_list = analytics_pb2.ListChartsLegend(
        charts=charts,
        legend=legend_element,
    )
    return chart_list

def form_group_chart_multy(records: dict, m_keys: list[str], start_date: int, end_date: int, 
                    chart_type: analytics_pb2.ChartType) -> analytics_pb2.ListChartsLegend:
    legend_element = {}
    charts: list[analytics_pb2.Chart] = []
    for i, m_key in enumerate(m_keys):
        color = availableColors[i % len(availableColors)]
        recs = get_one_multipl_records(records, m_key, start_date, end_date)
        if len(recs) > 0:
            if chart_type == analytics_pb2.PIE_CHART:
                recs = [rec[0] for rec in recs if len(rec) > 0]
                new_chart = analytics_pb2.Chart(
                    color=color,
                    type=chart_type,
                    records=count_records_percentage(recs)
                )
            else:
                new_chart = analytics_pb2.Chart(
                    color=color,
                    type=chart_type,
                    records=recs
                )
            legend_element[color] = MULTYPLIER_NAME[m_key]
            charts.append(new_chart)

    chart_list = analytics_pb2.ListChartsLegend(
        charts=charts,
        legend=legend_element,
    )
    return chart_list

def form_graph_info(
    records: dict, multipliers: dict, request: analytics_pb2.Filter
) -> dict[str, analytics_pb2.ListChartsLegend]:
    chart_map: dict = {}
    start_date = request.start_date
    end_date = request.end_date
    if start_date == end_date:
        # new group chart 1
        chart_map["Сравнение показателей"] = form_group_chart(
            records, [2110, 2120, 2200, 2400], start_date, end_date,
            analytics_pb2.BAR_CHART
        )
        chart_map["Соотношение собственных и заемных средств"] = form_group_chart(
            records, [1300, 1400, 1500, 1600], start_date, end_date,
            analytics_pb2.BAR_CHART
        )
        chart_map["АКТИВ: Внеоборотные активы"] = form_group_chart(
            records, [1110, 1120, 1130, 1140, 1150, 1160, 1170], start_date, end_date,
            analytics_pb2.PIE_CHART
        )
        chart_map["АКТИВ: Оборотные активы"] = form_group_chart(
            records, [1210, 1220, 1230, 1240, 1250, 1260], start_date, end_date,
            analytics_pb2.PIE_CHART
        )
        chart_map["ПАССИВ: Капитал и резервы"] = form_group_chart(
            records, [1310, 1320, 1340, 1350, 1360, 1370], start_date, end_date,
            analytics_pb2.PIE_CHART
        )
        chart_map["ПАССИВ: Долгосрочные обязательства"] = form_group_chart(
            records, [1410, 1420, 1430, 1450, 1460], start_date, end_date,
            analytics_pb2.PIE_CHART
        )
        chart_map["ПАССИВ: Краткосрочные обязательства"] = form_group_chart(
            records, [1510, 1520, 1530, 1535, 1540, 1550], start_date, end_date,
            analytics_pb2.PIE_CHART
        )
        return chart_map
    else:
        chart_map["Динамика рентабельности собственного капитала"] = form_group_chart(
            records, ["ROE"], start_date, end_date,
            analytics_pb2.BAR_CHART
        )
        return chart_map
    #     # profitability
    #     # single line
    #     recs = get_one_multipl_records(multipliers, "ROE", start_date, end_date)
    #     if len(recs) > 0:
    #         charts.append(
    #             analytics_pb2.Chart(
    #                 title="Динамика рентабельности собственного капитала",
    #                 records=recs,
    #                 type=analytics_pb2.BAR_CHART,
    #                 description="",
    #             )
    #         )
    #     # group line chart 1
    #     recs = get_one_multipl_records(multipliers, "GP_Margin", start_date, end_date)
    #     if len(recs) > 0:
    #         charts.append(
    #             analytics_pb2.Chart(
    #                 title="Динамика рентабельности продаж",
    #                 records=recs,
    #                 type=analytics_pb2.BAR_CHART,
    #                 description=MULTYPLIER_NAME["GP_Margin"],
    #             )
    #         )
    #     recs = get_one_multipl_records(multipliers, "OP_Margin", start_date, end_date)
    #     if len(recs) > 0:
    #         charts.append(
    #             analytics_pb2.Chart(
    #                 title="Динамика рентабельности продаж",
    #                 records=recs,
    #                 type=analytics_pb2.BAR_CHART,
    #                 description=MULTYPLIER_NAME["OP_Margin"],
    #             )
    #         )
    #     recs = get_one_multipl_records(multipliers, "NP_Margin", start_date, end_date)
    #     if len(recs) > 0:
    #         charts.append(
    #             analytics_pb2.Chart(
    #                 title="Динамика рентабельности продаж",
    #                 records=recs,
    #                 type=analytics_pb2.BAR_CHART,
    #                 description=MULTYPLIER_NAME["NP_Margin"],
    #             )
    #         )
    #     # group line chart 2
    #     recs = get_one_multipl_records(multipliers, "ROA", start_date, end_date)
    #     if len(recs) > 0:
    #         charts.append(
    #             analytics_pb2.Chart(
    #                 title="Динамика рентабельности активов",
    #                 records=recs,
    #                 type=analytics_pb2.BAR_CHART,
    #                 description=MULTYPLIER_NAME["ROA"],
    #             )
    #         )
    #     recs = get_one_multipl_records(multipliers, "ROCA", start_date, end_date)
    #     if len(recs) > 0:
    #         charts.append(
    #             analytics_pb2.Chart(
    #                 title="Динамика рентабельности активов",
    #                 records=recs,
    #                 type=analytics_pb2.BAR_CHART,
    #                 description=MULTYPLIER_NAME["ROCA"],
    #             )
    #         )
    #     recs = get_one_multipl_records(multipliers, "RONCA", start_date, end_date)
    #     if len(recs) > 0:
    #         charts.append(
    #             analytics_pb2.Chart(
    #                 title="Динамика рентабельности активов",
    #                 records=recs,
    #                 type=analytics_pb2.BAR_CHART,
    #                 description=MULTYPLIER_NAME["RONCA"],
    #             )
    #         )

    #     # liquidity
    #     # group line
    #     recs = get_one_multipl_records(multipliers, "CR", start_date, end_date)
    #     if len(recs) > 0:
    #         charts.append(
    #             analytics_pb2.Chart(
    #                 title="Динамика ликвидности",
    #                 records=recs,
    #                 type=analytics_pb2.BAR_CHART,
    #                 description=MULTYPLIER_NAME["CR"],
    #             )
    #         )
    #     recs = get_one_multipl_records(multipliers, "QR", start_date, end_date)
    #     if len(recs) > 0:
    #         charts.append(
    #             analytics_pb2.Chart(
    #                 title="Динамика ликвидности",
    #                 records=recs,
    #                 type=analytics_pb2.BAR_CHART,
    #                 description=MULTYPLIER_NAME["QR"],
    #             )
    #         )
    #     recs = get_one_multipl_records(multipliers, "AR", start_date, end_date)
    #     if len(recs) > 0:
    #         charts.append(
    #             analytics_pb2.Chart(
    #                 title="Динамика ликвидности",
    #                 records=recs,
    #                 type=analytics_pb2.BAR_CHART,
    #                 description=MULTYPLIER_NAME["AR"],
    #             )
    #         )
    #     # coefs
    #     # group line
    #     recs = get_one_multipl_records(
    #         multipliers, "Autonomy_Ratio", start_date, end_date
    #     )
    #     if len(recs) > 0:
    #         charts.append(
    #             analytics_pb2.Chart(
    #                 title="Динамика показателей финансовой устойчивости",
    #                 records=recs,
    #                 type=analytics_pb2.BAR_CHART,
    #                 description=MULTYPLIER_NAME["Autonomy_Ratio"],
    #             )
    #         )
    #     recs = get_one_multipl_records(
    #         multipliers, "Capitalization_Ratio", start_date, end_date
    #     )
    #     if len(recs) > 0:
    #         charts.append(
    #             analytics_pb2.Chart(
    #                 title="Динамика показателей финансовой устойчивости",
    #                 records=recs,
    #                 type=analytics_pb2.BAR_CHART,
    #                 description=MULTYPLIER_NAME["Capitalization_Ratio"],
    #             )
    #         )
    #     recs = get_one_multipl_records(
    #         multipliers, "Investment_Coverage_Ratio", start_date, end_date
    #     )
    #     if len(recs) > 0:
    #         charts.append(
    #             analytics_pb2.Chart(
    #                 title="Динамика показателей финансовой устойчивости",
    #                 records=recs,
    #                 type=analytics_pb2.BAR_CHART,
    #                 description=MULTYPLIER_NAME["Investment_Coverage_Ratio"],
    #             )
    #         )
    #     recs = get_one_multipl_records(
    #         multipliers, "Inventory_Security_Ratio", start_date, end_date
    #     )
    #     if len(recs) > 0:
    #         charts.append(
    #             analytics_pb2.Chart(
    #                 title="Динамика показателей финансовой устойчивости",
    #                 records=recs,
    #                 type=analytics_pb2.BAR_CHART,
    #                 description=MULTYPLIER_NAME["Inventory_Security_Ratio"],
    #             )
    #         )
    #     recs = get_one_multipl_records(
    #         multipliers, "Financial_Dependency_Ratio", start_date, end_date
    #     )
    #     if len(recs) > 0:
    #         charts.append(
    #             analytics_pb2.Chart(
    #                 title="Динамика показателей финансовой устойчивости",
    #                 records=recs,
    #                 type=analytics_pb2.BAR_CHART,
    #                 description=MULTYPLIER_NAME["Financial_Dependency_Ratio"],
    #             )
    #         )
    #     recs = get_one_multipl_records(
    #         multipliers, "Financial_Leverage_Ratio", start_date, end_date
    #     )
    #     if len(recs) > 0:
    #         charts.append(
    #             analytics_pb2.Chart(
    #                 title="Динамика показателей финансовой устойчивости",
    #                 records=recs,
    #                 type=analytics_pb2.BAR_CHART,
    #                 description=MULTYPLIER_NAME["Financial_Leverage_Ratio"],
    #             )
    #         )
    #     # table params
    #     # group bar chart
    #     recs = get_one_param_records(records, 2110, start_date, end_date)
    #     if len(recs) > 0:
    #         charts.append(
    #             analytics_pb2.Chart(
    #                 title="Динамика экономических показателей",
    #                 records=recs,
    #                 type=analytics_pb2.BAR_CHART,
    #                 description=CODE_NAME[int(2110)],
    #             )
    #         )
    #     recs = get_one_param_records(records, 2120, start_date, end_date)
    #     if len(recs) > 0:
    #         charts.append(
    #             analytics_pb2.Chart(
    #                 title="Динамика экономических показателей",
    #                 records=recs,
    #                 type=analytics_pb2.BAR_CHART,
    #                 description=CODE_NAME[int(2120)],
    #             )
    #         )
    #     recs = get_one_param_records(records, 2200, start_date, end_date)
    #     if len(recs) > 0:
    #         charts.append(
    #             analytics_pb2.Chart(
    #                 title="Динамика экономических показателей",
    #                 records=recs,
    #                 type=analytics_pb2.BAR_CHART,
    #                 description=CODE_NAME[int(2200)],
    #             )
    #         )
    #     recs = get_one_param_records(records, 2400, start_date, end_date)
    #     if len(recs) > 0:
    #         charts.append(
    #             analytics_pb2.Chart(
    #                 title="Динамика экономических показателей",
    #                 records=recs,
    #                 type=analytics_pb2.BAR_CHART,
    #                 description=CODE_NAME[int(2400)],
    #             )
    #         )
    #     return charts


def get_analitics_report(request: analytics_pb2.Filter) -> analytics_pb2.ChartReport:
    records, multipliers, report_summary = load_json()

    start_date = int(request.start_date)
    end_date = int(request.end_date)
    if start_date < valid_min_year or start_date > valid_max_year:
        request.start_date = valid_min_year
    if end_date > valid_max_year or end_date < valid_min_year:
        request.end_date = valid_max_year

    charts: dict[str, analytics_pb2.ListChartsLegend] = form_graph_info(records, multipliers, request)
    return_multy = []
    if request.start_date == request.end_date:
        for k, v in multipliers.items():
            if k == "_id":
                continue
            year = str(request.start_date)
            value = v.get(str(year), None)
            if value is not None and not math.isnan(value):
                return_multy.append(
                    analytics_pb2.Multiplier(key=MULTYPLIER_NAME[k], value=value)
                )

    dates = sorted([request.start_date, request.end_date])
    report = analytics_pb2.ChartReport(
        summary=report_summary[str(dates[0])][str(dates[1])],
        info=charts,
        multipliers=return_multy if len(return_multy) > 0 else None,
    )
    return report

def main():
    request = analytics_pb2.Filter(
        start_date=2022,
        end_date=2022
    )
    resp = get_analitics_report(request)
    print(resp)

main()