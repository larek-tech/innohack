import json
from analytics import analytics_pb2, analytics_pb2_grpc
from process.const import CODE_NAME, MULTYPLIER_NAME
import math
from pymongo import MongoClient


mongo = MongoClient("mongodb://46.138.243.191:27017/data", timeoutMS=30000**2)
records_col = mongo.get_database("data").get_collection("records")
multipliers_col = mongo.get_database("data").get_collection("multipliers")


def load_json():
    records = [r for r in records_col.find({})][0]
    multipliers = [m for m in multipliers_col.find({})][0]

    return records, multipliers


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


def form_graph_info(
    records: dict, multipliers: dict, request: analytics_pb2.Params
) -> list[analytics_pb2.Chart]:
    charts = []
    start_date = request.start_date
    end_date = request.end_date
    if start_date == end_date:
        # group bar chart 1
        recs = get_one_param_records(records, 2110, start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Сравнение показателей",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=CODE_NAME[int(2110)],
                )
            )

        recs = get_one_param_records(records, 2120, start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Сравнение показателей",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=CODE_NAME[int(2120)],
                )
            )

        recs = get_one_param_records(records, 2200, start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Сравнение показателей",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=CODE_NAME[int(2200)],
                )
            )

        recs = get_one_param_records(records, 2400, start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Сравнение показателей",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=CODE_NAME[int(2400)],
                )
            )
        # group bar chart 2
        recs = get_one_param_records(records, 1300, start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Соотношение собственных и заемных средств",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description="Капитал и резервы",
                )
            )
        recs = get_one_param_records(records, 1400, start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Соотношение собственных и заемных средств",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description="Долгострочные обязательства",
                )
            )
        recs = get_one_param_records(records, 1500, start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Соотношение собственных и заемных средств",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description="Краткосрочные обязательства",
                )
            )

        # single pie
        recs = [
            get_one_param_records(
                records, 1110, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1120, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1130, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1140, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1150, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1160, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1170, start_date, end_date, return_year=False
            ),
        ]
        recs = [rec[0] for rec in recs if len(rec) > 0]
        percent_recs = count_records_percentage(recs)
        if len(percent_recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="АКТИВ: Внеоборотные активы",
                    records=percent_recs,
                    type=analytics_pb2.PIE_CHART,
                    description="",
                )
            )

        recs = [
            get_one_param_records(
                records, 1210, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1220, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1230, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1240, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1250, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1260, start_date, end_date, return_year=False
            ),
        ]
        recs = [rec[0] for rec in recs if len(rec) > 0]
        percent_recs = count_records_percentage(recs)
        if len(percent_recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="АКТИВ: Оборотные активы",
                    records=percent_recs,
                    type=analytics_pb2.PIE_CHART,
                    description="",
                )
            )

        recs = [
            get_one_param_records(
                records, 1310, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1320, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1340, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1350, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1360, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1370, start_date, end_date, return_year=False
            ),
        ]
        recs = [rec[0] for rec in recs if len(rec) > 0]
        percent_recs = count_records_percentage(recs)
        if len(percent_recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="ПАССИВ: Капитал и резервы",
                    records=percent_recs,
                    type=analytics_pb2.PIE_CHART,
                    description="",
                )
            )

        recs = [
            get_one_param_records(
                records, 1410, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1420, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1430, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1450, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1460, start_date, end_date, return_year=False
            ),
        ]
        recs = [rec[0] for rec in recs if len(rec) > 0]
        percent_recs = count_records_percentage(recs)
        if len(percent_recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="ПАССИВ: Долгосрочные обязательства",
                    records=percent_recs,
                    type=analytics_pb2.PIE_CHART,
                    description="",
                )
            )

        recs = [
            get_one_param_records(
                records, 1510, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1520, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1530, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1535, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1540, start_date, end_date, return_year=False
            ),
            get_one_param_records(
                records, 1550, start_date, end_date, return_year=False
            ),
        ]
        recs = [rec[0] for rec in recs if len(rec) > 0]
        percent_recs = count_records_percentage(recs)
        if len(percent_recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="ПАССИВ: Краткосрочные обязательства",
                    records=percent_recs,
                    type=analytics_pb2.PIE_CHART,
                    description="",
                )
            )
        return charts
    else:
        # profitability
        # single line
        recs = get_one_multipl_records(multipliers, "ROE", start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Динамика рентабельности собственного капитала",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description="",
                )
            )
        # group line chart 1
        recs = get_one_multipl_records(multipliers, "GP_Margin", start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Динамика рентабельности продаж",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=MULTYPLIER_NAME["GP_Margin"],
                )
            )
        recs = get_one_multipl_records(multipliers, "OP_Margin", start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Динамика рентабельности продаж",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=MULTYPLIER_NAME["OP_Margin"],
                )
            )
        recs = get_one_multipl_records(multipliers, "NP_Margin", start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Динамика рентабельности продаж",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=MULTYPLIER_NAME["NP_Margin"],
                )
            )
        # group line chart 2
        recs = get_one_multipl_records(multipliers, "ROA", start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Динамика рентабельности активов",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=MULTYPLIER_NAME["ROA"],
                )
            )
        recs = get_one_multipl_records(multipliers, "ROCA", start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Динамика рентабельности активов",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=MULTYPLIER_NAME["ROCA"],
                )
            )
        recs = get_one_multipl_records(multipliers, "RONCA", start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Динамика рентабельности активов",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=MULTYPLIER_NAME["RONCA"],
                )
            )

        # liquidity
        # group line
        recs = get_one_multipl_records(multipliers, "CR", start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Динамика ликвидности",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=MULTYPLIER_NAME["CR"],
                )
            )
        recs = get_one_multipl_records(multipliers, "QR", start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Динамика ликвидности",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=MULTYPLIER_NAME["QR"],
                )
            )
        recs = get_one_multipl_records(multipliers, "AR", start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Динамика ликвидности",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=MULTYPLIER_NAME["AR"],
                )
            )
        # coefs
        # group line
        recs = get_one_multipl_records(
            multipliers, "Autonomy_Ratio", start_date, end_date
        )
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Динамика показателей финансовой устойчивости",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=MULTYPLIER_NAME["Autonomy_Ratio"],
                )
            )
        recs = get_one_multipl_records(
            multipliers, "Capitalization_Ratio", start_date, end_date
        )
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Динамика показателей финансовой устойчивости",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=MULTYPLIER_NAME["Capitalization_Ratio"],
                )
            )
        recs = get_one_multipl_records(
            multipliers, "Investment_Coverage_Ratio", start_date, end_date
        )
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Динамика показателей финансовой устойчивости",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=MULTYPLIER_NAME["Investment_Coverage_Ratio"],
                )
            )
        recs = get_one_multipl_records(
            multipliers, "Inventory_Security_Ratio", start_date, end_date
        )
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Динамика показателей финансовой устойчивости",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=MULTYPLIER_NAME["Inventory_Security_Ratio"],
                )
            )
        recs = get_one_multipl_records(
            multipliers, "Financial_Dependency_Ratio", start_date, end_date
        )
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Динамика показателей финансовой устойчивости",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=MULTYPLIER_NAME["Financial_Dependency_Ratio"],
                )
            )
        recs = get_one_multipl_records(
            multipliers, "Financial_Leverage_Ratio", start_date, end_date
        )
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Динамика показателей финансовой устойчивости",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=MULTYPLIER_NAME["Financial_Leverage_Ratio"],
                )
            )
        # table params
        # group bar chart
        recs = get_one_param_records(records, 2110, start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Динамика экономических показателей",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=CODE_NAME[int(2110)],
                )
            )
        recs = get_one_param_records(records, 2120, start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Динамика экономических показателей",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=CODE_NAME[int(2120)],
                )
            )
        recs = get_one_param_records(records, 2200, start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Динамика экономических показателей",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=CODE_NAME[int(2200)],
                )
            )
        recs = get_one_param_records(records, 2400, start_date, end_date)
        if len(recs) > 0:
            charts.append(
                analytics_pb2.Chart(
                    title="Динамика экономических показателей",
                    records=recs,
                    type=analytics_pb2.BAR_CHART,
                    description=CODE_NAME[int(2400)],
                )
            )
        return charts


def get_analitics_report(request: analytics_pb2.Filter) -> analytics_pb2.ChartReport:
    records, multipliers = load_json()
    charts: list[analytics_pb2.Chart] = form_graph_info(records, multipliers, request)
    return_multy = []
    if request.start_date == request.end_date:
        for k, v in multipliers.items():
            year = str(request.start_date)
            value = v.get(str(year), None)
            if value is not None and not math.isnan(value):
                return_multy.append(
                    analytics_pb2.Multiplier(key=MULTYPLIER_NAME[k], value=value)
                )
    report = analytics_pb2.ChartReport(
        charts=charts,
        multipliers=return_multy if len(return_multy) > 0 else None,
    )
    return report
