from pathlib import Path

from rag.minio_client import MinioClient
from rag.db import QdrantBase
from rag.utils.bi_encode import get_bi_encoder
from rag.model import LLMClient

from rag.config import (
    MINIO_HOST, 
    MINIO_ACCESS_KEY, 
    MINIO_SECRET_KEY,
    QDRANT_HOST,
    QDRANT_PORT,
    BI_ENCODE_NAME,
    CHUNK_OVERLAP,
    CHUNK_SIZE
)

class RagClient:

    llm_client = LLMClient()

    def __init__(self):

        # minio_client = MinioClient(MINIO_HOST, MINIO_ACCESS_KEY, MINIO_SECRET_KEY)
        qdrant_client = QdrantBase(QDRANT_HOST, QDRANT_PORT)

        bi_encoder, vec_dim = get_bi_encoder(BI_ENCODE_NAME)


        # Скачивание данных с S3
        # minio_client.download_all_files("innohack", "data")

        dir_path = Path(__file__) / "data"
        file_paths = list(dir_path.glob('*'))

        # Вектаризация данных
        qdrant_client.files_to_vecdb(
            files=file_paths,
            bi_encoder=bi_encoder,
            vec_size=vec_dim,
            sep="\n",
            chunk_size=CHUNK_SIZE,
            chunk_overlap=CHUNK_OVERLAP,
        )
