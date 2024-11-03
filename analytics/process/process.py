from minio import Minio
import typing as tp
import pandas as pd
import numpy as np
import bson
from pymongo import MongoClient

from process.compute_business_metrics import read_excel, code_column, get_liquidity, profitability_of_sales, profitability_of_assets, coefficients, coefficients_3years
from process.const import CODE_NAME, MULTYPLIER_NAME


mongo = MongoClient("mongodb://46.138.243.191:27017/data", timeoutMS=30000**2)
records_col = mongo.get_database("data").get_collection("records")
multipliers_col = mongo.get_database("data").get_collection("multipliers")

s3 = Minio(
    "s3.larek.tech",
    access_key="I3gAX8ygZF1pnXuSSo00",
    secret_key="ah2d805PcxU0PxYj0uhLzrasDnsgPi2xFGycpDm8",
    secure=True,
)
BUCKET_NAME = "innohack"
PERIODS = ["1", "2", "3"]
REVERSED_MULTYPLIER_NAME = {v.lower(): k for k, v in MULTYPLIER_NAME.items()}

def list_files() -> tp.List[str]:
    paths = []
    files  = s3.list_objects(BUCKET_NAME, "mts/excel_data", recursive=True)
    for file in files:
        path = f"./data/{file.object_name}"
        s3.fget_object(BUCKET_NAME, file.object_name, path)
        paths.append(path)
    return paths

def get_value_by_code(df: pd.DataFrame, code: int, period: str):
    try:
        return df[df[code_column] == code][period].values[0]
    except:
        return None

def add_one_param_records(records: dict, df: pd.DataFrame, code: int, year: int) -> dict:
    if str(code) not in records:
        records[str(code)] = {}

    for period in PERIODS:
        value = get_value_by_code(df, code, period)
        substraction_value = int(period) - 1
        year_key = str(year - substraction_value)
        if value and value != np.nan and value is not None:
            if year_key not in records[str(code)]:
                records[str(code)][year_key] = float(value)
            else:
                records[str(code)][year_key] = float(value)
    return records

def parse_df_to_dict(records: dict, df: pd.DataFrame, year: str):
    codes = list(CODE_NAME.keys())
    for code in codes:
        records = add_one_param_records(records, df, code, year)
    
    return records

def parse_multy_to_dict(records: dict, df: pd.DataFrame) -> dict:
    for _, row in df.iterrows():
        metric_name = row['metric_name']
        value = row['matric_value']
        year = row['year']
        
        if metric_name in REVERSED_MULTYPLIER_NAME:
            abbrev = REVERSED_MULTYPLIER_NAME[metric_name]
            
            if abbrev not in records:
                records[abbrev] = {}
            if value and value != np.nan and value is not None:
                records[abbrev][str(year)] = value
    
    return records

def save_data(records: dict, multipliers: dict):
    records, multipliers = preprocess_xlsx()

    has_records = len([r for r in records_col.find({})]) > 0
    has_multipliers = len([m for m in multipliers_col.find({})]) > 0
    
    if not has_records:
        records_col.insert_one(records)
    if not has_multipliers:
        multipliers_col.insert_one(multipliers)

def preprocess_xlsx():
    excel_paths = list_files()
    records = {}
    multipliers = {}
    for file in excel_paths:
        dfs, year = read_excel(file)
        for df in dfs:
            records = parse_df_to_dict(records, df, year)

        metrics_for_chart = pd.concat(
            [
                get_liquidity(dfs, year),
                profitability_of_sales(dfs, year),
                profitability_of_assets(dfs, year),
                coefficients(dfs, year),
                coefficients_3years(dfs, year),
            ]
        )
        multipliers = parse_multy_to_dict(multipliers, metrics_for_chart)

    save_data(records, multipliers)


