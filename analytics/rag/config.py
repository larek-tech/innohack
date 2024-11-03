import os

from dotenv import load_dotenv

load_dotenv()

MINIO_HOST = os.getenv("MINIO_HOST")
MINIO_ACCESS_KEY = os.getenv("MINIO_ACCESS_KEY")
MINIO_SECRET_KEY= os.getenv("MINIO_SECRET_KEY")

QDRANT_HOST = os.getenv("QDRANT_HOST")
QDRANT_PORT = int(os.getenv("QDRANT_PORT"))

BI_ENCODE_NAME = os.getenv("BI_ENCODE_NAME")

CHUNK_SIZE = int(os.getenv("CHUNK_SIZE"))
CHUNK_OVERLAP = int(os.getenv("CHUNK_OVERLAP"))