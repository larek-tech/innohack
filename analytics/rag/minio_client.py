import os

from minio import Minio
from minio.error import S3Error
from loguru import logger


class MinioClient:
    def __init__(self, endpoint, access_key, secret_key):
        self.client = Minio(endpoint, access_key=access_key, secret_key=secret_key)

    def create_bucket(self, bucket_name):
        try:
            self.client.make_bucket(bucket_name)
            print(f"Bucket '{bucket_name}' created successfully.")
        except S3Error as e:
            if e.code == "BucketAlreadyOwnedByYou":
                print(f"Bucket '{bucket_name}' already exists.")
            else:
                print(f"Error occurred: {e}")

    def upload_file(self, bucket_name, file_path, object_name):
        try:
            self.client.fput_object(bucket_name, object_name, file_path)
            print(
                f"File '{file_path}' uploaded to bucket '{bucket_name}' as '{object_name}'."
            )
        except S3Error as e:
            print(f"Error occurred: {e}")

    def download_file(self, bucket_name, object_name, file_path):
        try:
            self.client.fget_object(bucket_name, object_name, file_path)
            print(
                f"File '{object_name}' downloaded from bucket '{bucket_name}' to '{file_path}'."
            )
        except S3Error as e:
            print(f"Error occurred: {e}")

    def download_all_files(self, bucket_name, download_dir):
        try:
            # Создаем директорию, если она не существует
            os.makedirs(download_dir, exist_ok=True)

            # Получаем список объектов в бакете
            objects = self.client.list_objects(
                bucket_name, prefix="mts/moskva_mts_ru", recursive=True
            )
            logger.info(f"found {len(objects)} in s3")
            for obj in objects:
                logger.info(obj.object_name)
                file_path = os.path.join(download_dir, obj.object_name)
                self.client.fget_object(bucket_name, obj.object_name, file_path)
                print(f"File '{obj.object_name}' downloaded to '{file_path}'.")
        except S3Error as e:
            print(f"Error occurred: {e}")


if __name__ == "__main__":
    minio_client = MinioClient(
        "s3.larek.tech",
        "I3gAX8ygZF1pnXuSSo00",
        "ah2d805PcxU0PxYj0uhLzrasDnsgPi2xFGycpDm8",
    )
    minio_client.download_all_files("innohack", "moskva_mts_ru")
