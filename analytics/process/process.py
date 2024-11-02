from minio import Minio
import typing as tp
import pandas as pd

s3 = Minio(
    "s3.larek.tech",
    access_key="I3gAX8ygZF1pnXuSSo00",
    secret_key="ah2d805PcxU0PxYj0uhLzrasDnsgPi2xFGycpDm8",
    secure=True,
)
BUCKET_NAME = "innohack"

def list_files() -> tp.List[str]:
    files  = s3.list_objects(BUCKET_NAME, "mts/excel_data", recursive=True)
    for file in files:
        path = f"./data/{file.object_name}"
        s3.fget_object(BUCKET_NAME, file.object_name, path)
        excel = pd.read_excel(path)
        print(excel.head())

list_files()

def preprocess_xlsx():
    pass