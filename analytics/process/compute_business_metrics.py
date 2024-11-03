import pandas as pd
import numpy as np
import os
from typing import List, Tuple

code_column = "Код"


"""
ликвидность done
Рентабельность оборотных средств done
Рентабельность продаж done
Рентабельность активов done
Рентабельность собственного капитала done
P/E нужна мосбиржа (цена акции)
P/BV 
EV/S
EV/EBITDA  
Долг/EBITDA
ROE done
Рентабельность продаж done
"""


def read_excel(path_to_excel: str) -> Tuple[List[pd.DataFrame], int]:

    sheets = pd.ExcelFile(path_to_excel).sheet_names[:2]
    xls = []
    n_start_row = pd.read_excel(path_to_excel)
    year = int(n_start_row["Unnamed: 3"][4])
    n_start_row = n_start_row[
        (n_start_row["Unnamed: 0"] == "АКТИВ")
        | (n_start_row["Unnamed: 0"] == "Наименование показателя")
    ].index.min()

    for sheet in sheets:
        df = pd.read_excel(
            path_to_excel, sheet_name=sheet, skiprows=n_start_row + 1
        )  # skiprows=9

        # Переименование колонок
        new_column_names = df.columns.tolist()

        new_column_names[1] = "Код"
        if len(new_column_names) > 2:
            new_column_names[2] = "1"
        if len(new_column_names) > 3:
            new_column_names[3] = "2"
        if len(new_column_names) > 4:
            new_column_names[4] = "3"

        df.columns = new_column_names

        df.iloc[:, 1:] = df.iloc[:, 1:].applymap(
            lambda x: np.nan if isinstance(x, str) and not x[0].isdigit() else x
        )
        df.iloc[:, 1:] = df.iloc[:, 1:].astype(float)
        xls += [df]
    return xls, year


# Ликвидность текущая, быстрая, абсолютная
def get_liquidity(df: List[pd.DataFrame], year: int) -> pd.DataFrame:
    """
    Текущая ликвидность = стр. 1200 / стр. 1500 где:
    Стр. 1200 — номер строки итога раздела II «Оборотные активы» бухгалтерского баланса;
    Стр. 1500 — номер строки итога раздела V «Краткосрочные обязательства» бухгалтерского баланса.
    Чем показатель больше, тем лучше платежеспособность предприятия.

    Быстрая ликвидность = (стр. 1230 стр. 1240 стр. 1250) / (стр. 1510 стр. 1520 стр. 1550)

    Абсолютная ликвидность = (стр. 1250 стр. 1240) / (стр. 1510 стр. 1520 стр. 1550)

    """

    data = df[0]

    def compute(year: str) -> List[float]:
        return [
            data[data[code_column] == 1200][year].values[0]
            / data[data[code_column] == 1500][year].values[0],
            (
                data[data[code_column] == 1230][year].values[0]
                + data[data[code_column] == 1240][year].values[0]
                + data[data[code_column] == 1250][year].values[0
            )
            / (
                data[data[code_column] == 1510][year].values[0]
                + data[data[code_column] == 1520][year].values[0]
                + data[data[code_column] == 1550][year].values[0]
            ),
            (
                data[data[code_column] == 1250][year].values[0]
                + data[data[code_column] == 1240][year].values[0]
            )
            / (
                data[data[code_column] == 1510][year].values[0]
                + data[data[code_column] == 1520][year].values[0]
                + data[data[code_column] == 1550][year].values[0]
            ),
        ]

    chart_df = pd.DataFrame(
        {
            "metric_name": [
                "текущая ликвидность",
                "быстрая ликвидность",
                "абсолютная ликвидность",
            ]
            * 3,
            "matric_value": compute("1") + compute("2") + compute("3"),
            "year": [year] * 3 + [year - 1] * 3 + [year - 2] * 3,
        }
    )

    return chart_df


# Рентабельность продаж, Рентабельность собственного капитала
def profitability_of_sales(df: List[pd.DataFrame], year: int) -> pd.DataFrame:
    """
    Рентабельность продаж по валовой прибыли = строка 2100 / строка 2110 × 100
    Рентабельность продаж по операционной прибыли = (строка 2300 строка 2330) / строка 2110 × 100
    Рентабельность продаж по чистой прибыли = строка 2400 / строка 2110 × 100

    Рентабельность собственного капитала = стр. 2400/ стр. 1300 × 100. где:
    Стр. 2400 -строка отчета о финансовых результатах (чистая прибыль компании);
    Стр. 1300 — строка бухгалтерского баланса (итоговая строка раздела III «Капитал и резервы»).

    """

    data_1, data_2 = df[0], df[1]

    def compute(year: str) -> List[float]:
        return [
            data_2[data_2[code_column] == 2100][year].values[0]
            / data_2[data_2[code_column] == 2110][year].values[0],
            (
                data_2[data_2[code_column] == 2300][year].values[0]
                + data_2[data_2[code_column] == 2330][year].values[0]
            )
            / data_2[data_2[code_column] == 2110][year].values[0],
            data_2[data_2[code_column] == 2400][year].values[0]
            / data_2[data_2[code_column] == 2110][year].values[0],
            data_2[data_2[code_column] == 2400][year].values[0]
            / data_1[data_1[code_column] == 1300][year].values[0],
        ]

    chart_df = pd.DataFrame(
        {
            "metric_name": [
                "рентабельность продаж по валовой прибыли",
                "рентабельность продаж по операционной прибыли",
                "рентабельность продаж по чистой прибыли",
                "рентабельность собственного капитала",
            ]
            * 2,
            "matric_value": compute("1") + compute("2"),
            "year": [year] * 4 + [year - 1] * 4,
        }
    )

    return chart_df


# ROE (Рентабельность активов), Рентабельность оборотных активов, Рентабельность внеоборотных активов, Рентабельность собственного капитала
def profitability_of_assets(
    df: List[pd.DataFrame], year: int, profit_type="sales"
) -> pd.DataFrame:
    """
    Рентабельности активов (ROE) = прибыль за период / средняя величина активов за период х 100%
    Показатели прибыли для числителя формулы рентабельности активов нужно взять из отчета о финансовых результатах:
    прибыль от продаж — из строки 2200;
    чистую прибыль — из строки 2400.
    Средняя величина активов = (Сальдо баланса на начало периода + Сальдо баланса на конец периода) / 2, где сальдо баланса берется из строки 1600.

    Рентабельность оборотных активов = (2400 или 2200) / средняя величина оборотных активов * 100%
    Средняя величина оборотных активов = (Итог раздела II актива баланса на начало периода + Итог раздела II актива баланса на конец периода) / 2, берется из строки 1200.

    Рентабельность внеоборотных активов = (2400 или 2200) / Средняя величина внеоборотных активов
    Средняя величина внеоборотных активов = (Итог раздела I актива баланса на начало периода + Итог раздела I актива баланса на конец периода) / 2, берется из строки 1100.

    """

    # В числителе используется прибыль от продаж

    if profit_type == "sales":
        profit_code = 2200
    elif profit_type == "net_profit":
        profit_code = 2400
    else:
        raise ValueError("incorrect type of profit! Choose sales or net_profit")

    data_1, data_2 = df[0], df[1]

    def compute(year: str) -> List[float]:

        year_prev = str(int(year) + 1)

        return [
            data_2[data_2[code_column] == profit_code]["1"].values[0]
            / (
                (
                    data_1[data_1[code_column] == 1600][year].values[0]
                    + data_1[data_1[code_column] == 1600][year_prev].values[0]
                )
                / 2
            ),
            data_2[data_2[code_column] == profit_code][year].values[0]
            / (
                (
                    data_1[data_1[code_column] == 1200][year].values[0]
                    + data_1[data_1[code_column] == 1200][year_prev].values[0]
                )
                / 2
            ),
            data_2[data_2[code_column] == profit_code][year].values[0]
            / (
                (
                    data_1[data_1[code_column] == 1100][year].values[0]
                    + data_1[data_1[code_column] == 1100][year_prev].values[0]
                )
                / 2
            ),
        ]

    chart_df = pd.DataFrame(
        {
            "metric_name": [
                "рентабельность всех активов",
                "рентабельность оборотных активов",
                "рентабельность внеоборотных активов",
            ]
            * 2,
            "matric_value": compute("1") + compute("2"),
            "year": [year] * 3 + [year - 1] * 3,
        }
    )

    return chart_df


def coefficients(df: List[pd.DataFrame], year: int) -> pd.DataFrame:
    """
    Коэффициент автономии = Собственный капитал / Активы = 2400 / (1100 + 1200)

    Коэффициент капитализации = Долгосрочные обязательства / (Долгосрочные обязательства + Собственный капитал)
    = 1400 / (1400 + 2400)

    Коэффициент покрытия инвестиций = (Собственный капитал+F1[1400])/F1[1600], где
    Собственный капитал = сумме раздела Баланса "Капитал и резервы" плюс задолженность учредителей по взносам в уставный капитал;
    F1[1400] – строка баланса "Итого долгосрочные обязательства";
    F1[1600] – итого Баланс (т.е. общая сумма активов организации).
    = (2400 + 1400) / 1600

    """

    data_1, data_2 = df[0], df[1]

    def compute(year: str) -> List[float]:
        return [
            data_2[data_2[code_column] == 2400][year].values[0]
            / (
                data_1[data_1[code_column] == 1100][year].values[0]
                + data_1[data_1[code_column] == 1200][year].values[0]
            ),
            data_1[data_1[code_column] == 1400][year].values[0]
            / (
                data_1[data_1[code_column] == 1400][year].values[0]
                + data_2[data_2[code_column] == 2400][year].values[0]
            ),
            (
                data_1[data_1[code_column] == 1400][year].values[0]
                + data_2[data_2[code_column] == 2400][year].values[0]
            )
            / data_1[data_1[code_column] == 1600][year].values[0],
        ]

    chart_df = pd.DataFrame(
        {
            "metric_name": [
                "коэффициент автономии",
                "коэффициент капитализации",
                "коэффициент покрытия инвестиций",
            ]
            * 2,
            "matric_value": compute("1") + compute("2"),
            "year": [year] * 3 + [year - 1] * 3,
        }
    )

    return chart_df


def coefficients_3years(df: List[pd.DataFrame], year: int) -> pd.DataFrame:
    """
    Коэффициент обеспеченности материальных запасов = (стр. 1300 – стр. 1100 ) / стр. 1210

    Коэффициент финансовой зависимости = Обязательства / Активы = 1400 / 1600

    Коэффициент финансового левериджа = Обязательства / Собственный капитал = (1400 + 1500) / 1300

    """

    data_1 = df[0]

    def compute(year: str) -> List[float]:
        return [
            (
                data_1[data_1[code_column] == 1300][year].values[0]
                - data_1[data_1[code_column] == 1100][year].values[0]
            )
            / data_1[data_1[code_column] == 1210][year].values[0],
            data_1[data_1[code_column] == 1400][year].values[0]
            / data_1[data_1[code_column] == 1600][year].values[0],
            (
                data_1[data_1[code_column] == 1400][year].values[0]
                - data_1[data_1[code_column] == 1500][year].values[0]
            )
            / data_1[data_1[code_column] == 1300][year].values[0],
        ]

    chart_df = pd.DataFrame(
        {
            "metric_name": [
                "коэффициент обеспеченности материальных запасов",
                "коэффициент финансовой зависимости",
                "коэффициент финансового левериджа",
            ]
            * 3,
            "matric_value": compute("1") + compute("2") + compute("3"),
            "year": [year] * 3 + [year - 1] * 3 + [year - 2] * 3,
        }
    )

    return chart_df


def main():
    for file in os.listdir("excel_data"):
        df, year = read_excel("excel_data/" + file)
        metrics_for_chart = pd.concat(
            [
                get_liquidity(df, year),
                profitability_of_sales(df, year),
                profitability_of_assets(df, year),
                coefficients(df, year),
                coefficients_3years(df, year),
            ]
        )

        print(metrics_for_chart)


if __name__ == "__main__":
    main()
