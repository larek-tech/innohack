from compute_business_metrics import read_excel, profitability_of_sales, code_column

from dataclasses import dataclass
import pandas as pd

@dataclass
class Record:
    x: str
    y: float

@dataclass
class Chart:
    title: str
    records: list[Record]
    type: str
    description: str

periods = ["1", "2", "3"]


def get_value_by_code(df: pd.DataFrame, code: int, period: str):
    try:
        return df[df[code_column] == code][period].values[0]
    except:
        return None

def get_one_param_records(df: pd.DataFrame, code: int, year: int) -> list[Record]:
    records = []
    for period in periods:
        value = get_value_by_code(df, code, period)
        substraction_value = int(period) - 1
        if value:
            records.append(
                Record(
                    x=str(year - substraction_value),
                    y=value
                )
            )
    return records

def count_records_percentage(records: list[Record]) -> list[Record]:
    percent_records = []
    whole_sum = 0
    for record in records:
        whole_sum += record.y

    for record in records:
        percent_records.append(
            Record(
                x=record.x,
                y=(record.y/whole_sum)*100
            )
        )
    return percent_records

def form_pie_chart(df: pd.DataFrame, title: str, codes: list[int]) -> Chart:
    pass

def form_graphs_for_one_xlsx(xlsx_path: str) -> list[Chart]:
    graphs: list[Chart] = []
    dfs, year = read_excel(xlsx_path)
    year = int(year)

    # single bar charts
    graphs.append(Chart(
        title="Выручка",
        records=get_one_param_records(dfs[1], 2110, year),
        type="bar chart",
        description=""
    ))
    graphs.append(Chart(
        title="Себестоимость",
        records=get_one_param_records(dfs[1], 2120, year),
        type="bar chart",
        description=""
    ))
    graphs.append(Chart(
        title="Прибыль",
        records=get_one_param_records(dfs[1], 2200, year),
        type="bar chart",
        description=""
    ))
    # graphs.append(Chart(
    #     title="EBITDA",
    #     records=get_one_param_records(dfs[1], 2200, year),
    #     type="bar chart",
    #     description=""
    # ))

    # single pie chart
    graphs.append(Chart(
        title="Прибыль",
        records=get_one_param_records(dfs[1], 2200, year),
        type="pie chart",
        description=""
    ))
    return graphs

if __name__ == "__main__":
    graphs = form_graphs_for_one_xlsx("/home/hope/Hope/Innohack/excel_data/МТС_2011.xlsx")
    print(graphs)
    