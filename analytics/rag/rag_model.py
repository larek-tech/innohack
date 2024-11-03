from pathlib import Path

from loguru import logger

from rag.minio_client import MinioClient
from rag.db import qdrant_client, files_to_vecdb
from rag.utils.bi_encode import get_bi_encoder
from rag.model import LLMClient
from rag.convert import format_data
from rag.config import (
    MINIO_HOST,
    MINIO_ACCESS_KEY,
    MINIO_SECRET_KEY,
    QDRANT_HOST,
    QDRANT_PORT,
    BI_ENCODE_NAME,
    CHUNK_OVERLAP,
    CHUNK_SIZE,
)


from process.form_graphs import load_json


class RagClient:

    llm_client = LLMClient()

    def __init__(self):

        # minio_client = MinioClient(MINIO_HOST, MINIO_ACCESS_KEY, MINIO_SECRET_KEY)
        self.qdrant_client = qdrant_client

        # Загркзка json из Mongo
        records, multipliers = load_json()

        logger.info(records)

        new_records = format_data(records)
        new_multipliers = format_data(multipliers)
        dir_path = str(Path(__file__).parent / "data")
        with open(dir_path + "/new_records.txt", "w", encoding="utf-8") as file:
            file.write(new_records)

        with open(dir_path + "/multipliers.txt", "w", encoding="utf-8") as file:
            file.write(new_multipliers)

        bi_encoder, vec_dim = get_bi_encoder(BI_ENCODE_NAME)

        # Скачивание данных с S3
        # minio_client.download_all_files("innohack", "data")

        dir_path = Path(__file__).parent / "data"
        file_paths = list(dir_path.glob("*"))

        # Вектаризация данных
        files_to_vecdb(
            files=file_paths,
            bi_encoder=bi_encoder,
            vec_size=vec_dim,
            sep="\n",
            chunk_size=CHUNK_SIZE,
            chunk_overlap=CHUNK_OVERLAP,
        )
