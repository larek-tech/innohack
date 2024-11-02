import pandas as pd
import numpy as np
from typing import List, Tuple

code_column = "Код стр."

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

    sheets = pd.ExcelFile(path_to_excel).sheet_names
    xls = []
    n_start_row = pd.read_excel(path_to_excel)
    year = int(n_start_row["Unnamed: 3"][4])
    n_start_row = n_start_row[
        (n_start_row["Unnamed: 0"] == "АКТИВ")
        | (n_start_row["Unnamed: 0"] == "Наименование показателя")
    ].index[0]

    for sheet in sheets:
        df = pd.read_excel(
            path_to_excel, sheet_name=sheet, skiprows=n_start_row + 1
        )  # skiprows=9

        # Переименование колонок
        new_column_names = df.columns.tolist()

        if len(new_column_names) > 2:
            new_column_names[2] = "1"
        if len(new_column_names) > 3:
            new_column_names[3] = "2"
        if len(new_column_names) > 4:
            new_column_names[4] = "3"

        df.columns = new_column_names

        df.replace("-", np.nan, inplace=True)
        df.iloc[:, 1:] = df.iloc[:, 1:].astype(float)
        xls += [df]
    return xls, year


# Ликвидность текущая, быстрая, абсолютная
def get_liquidity(df: List[pd.DataFrame]) -> List[float]:
    """
    Текущая ликвидность = стр. 1200 / стр. 1500 где:
    Стр. 1200 — номер строки итога раздела II «Оборотные активы» бухгалтерского баланса;
    Стр. 1500 — номер строки итога раздела V «Краткосрочные обязательства» бухгалтерского баланса.
    Чем показатель больше, тем лучше платежеспособность предприятия.

    Быстрая ликвидность = (стр. 1230 стр. 1240 стр. 1250) / (стр. 1510 стр. 1520 стр. 1550)

    Абсолютная ликвидность = (стр. 1250 стр. 1240) / (стр. 1510 стр. 1520 стр. 1550)

    """

    data = df[1]
    return (
        data[data[code_column] == 1200]["1"].values[0]
        / data[data[code_column] == 1500]["1"].values[0],
        (
            data[data[code_column] == 1230]["1"].values[0]
            + data[data[code_column] == 1240]["1"].values[0]
            + data[data[code_column] == 1250]["1"].values[0]
        )
        / (
            data[data[code_column] == 1510]["1"].values[0]
            + data[data[code_column] == 1520]["1"].values[0]
            + data[data[code_column] == 1550]["1"].values[0]
        ),
        (
            data[data[code_column] == 1250]["1"].values[0]
            + data[data[code_column] == 1240]["1"].values[0]
        )
        / (
            data[data[code_column] == 1510]["1"].values[0]
            + data[data[code_column] == 1520]["1"].values[0]
            + data[data[code_column] == 1550]["1"].values[0]
        ),
    )


# Рентабельность продаж
def profitability_of_sales(df: List[pd.DataFrame]) -> List[float]:
    """
    Рентабельность продаж по валовой прибыли = строка 2100 / строка 2110 × 100
    Рентабельность продаж по операционной прибыли = (строка 2300 строка 2330) / строка 2110 × 100
    Рентабельность продаж по чистой прибыли = строка 2400 / строка 2110 × 100
    """

    data = df[1]
    return [
        data[data[code_column] == 2100]["1"].values[0]
        / data[data[code_column] == 2110]["1"].values[0]
        * 100,
        (
            data[data[code_column] == 2300]["1"].values[0]
            + data[data[code_column] == 2330]["1"].values[0]
        )
        / data[data[code_column] == 2110]["1"].values[0]
        * 100,
        data[data[code_column] == 2400]["1"].values[0]
        / data[data[code_column] == 2110]["1"].values[0]
        * 100,
    ]


# ROA, Рентабельность оборотных активов, Рентабельность внеоборотных активов, Рентабельность собственного капитала
def profitability_of_assets(df: List[pd.DataFrame], profit_type="sales") -> List[float]:
    """
    Рентабельности активов = прибыль за период / средняя величина активов за период х 100%
    Показатели прибыли для числителя формулы рентабельности активов нужно взять из отчета о финансовых результатах:
    прибыль от продаж — из строки 2200;
    чистую прибыль — из строки 2400.

    Рентабельность собственного капитала = стр. 2400/ стр. 1300 × 100.

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
    return [
        data_2[data_2[code_column] == profit_code]["1"].values[0]
        / (
            (
                data_1[data_1[code_column] == 1600]["1"].values[0]
                + data_1[data_1[code_column] == 1600]["2"].values[0]
            )
            / 2
        )
        * 100,
        data_2[data_2[code_column] == profit_code]["1"].values[0]
        / (
            (
                data_1[data_1[code_column] == 1200]["1"].values[0]
                + data_1[data_1[code_column] == 1200]["2"].values[0]
            )
            / 2
        )
        * 100,
        data_2[data_2[code_column] == profit_code]["1"].values[0]
        / (
            (
                data_1[data_1[code_column] == 1100]["1"].values[0]
                + data_1[data_1[code_column] == 1100]["2"].values[0]
            )
            / 2
        )
        * 100,
    ]


def main():
    df, year = read_excel("excel_data/МТС_2011.xlsx")
    # res1 = profitability_of_sales(df)
    # res2 = profitability_of_assets(df)
    # print(res1)
    # print(res2)


if __name__ == "__main__":
    main()
