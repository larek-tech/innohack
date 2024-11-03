import uuid
import requests
import json
from typing import List

from qdrant_client import QdrantClient
from qdrant_client.http.models import Distance, VectorParams, PointStruct

from utils.bi_encode import str_to_vec
from utils.chunks import file_to_chunks

from loguru import logger

# Создаем подключение к векторной БД
qdrant_client = QdrantClient("https://qdrant.larek.tech", port=443)

COLL_NAME = "test_chuncks"
COLL_QUESTION_NAME = "test_questions_chuncks"


def get_questions_for_chunk(chunk_text: str) -> str:

    url = "https://mts-aidocprocessing-case.olymp.innopolis.university/generate"
    data = {
        "prompt": chunk_text,
        "apply_chat_template": True,
        "system_prompt": "Представь, что ты финансовый аналитик. Тебе подаётся чанк текста из финансового документа. \
              Напиши 7 вопросов, связанных с финансами и бухгалтерией на которые можно было бы найти ответ, исходя из информации в чанке. \
                Вопросы про количественные значения (показатели) приветствуются. Формулируй вопросы на русском языке",
        "max_tokens": 512,
        "n": 1,
        "temperature": 0.7,
    }

    headers = {"Content-Type": "application/json"}

    response = requests.post(url, data=json.dumps(data), headers=headers)

    if response.status_code == 200:
        logger.info(response.json())
        return response.json()
    else:
        return f"Error: {response.status_code} - {response.text}"


# Помещаем чанки и доп. информаицю в векторую БД
def save_chunks(
    bi_encoder, chunks: List[str], file_name: str, questions_for_chunk: List[str]
):
    # Конвертируем чанки в векитора
    chunk_embeddings = str_to_vec(bi_encoder, chunks)
    questions_for_chunk_embeddings = str_to_vec(bi_encoder, questions_for_chunk)

    # Содаем объект(ы) для БД
    points = []
    points_question = []
    chunk_uuid = 0  # генерируем GUID
    for i in range(len(chunk_embeddings)):

        chunk_uuid = str(uuid.uuid4())

        point = PointStruct(
            id=chunk_uuid,
            vector=chunk_embeddings[i],
            payload={"file": file_name, "chunk": chunks[i]},
        )
        points.append(point)

        point = PointStruct(
            id=chunk_uuid,
            vector=questions_for_chunk_embeddings[i],
            payload={"file": file_name, "chunk": chunks[i]},
        )
        points_question.append(point)

    # Сохраняем вектора в БД
    operation_info = qdrant_client.upsert(
        collection_name=COLL_NAME, wait=True, points=points
    )
    operation_info = qdrant_client.upsert(
        collection_name=COLL_QUESTION_NAME, wait=True, points=points_question
    )


    def files_to_vecdb(self, 
                       files, 
                       bi_encoder, 
                       vec_size, 
                       sep, 
                       chunk_size, 
                       chunk_overlap
    ):

    # Коллекция для чанков
    qdrant_client.delete_collection(collection_name=COLL_NAME)
    qdrant_client.create_collection(
        collection_name=COLL_NAME,
        vectors_config=VectorParams(size=vec_size, distance=Distance.COSINE),
    )

    # Коллекция для вопросов к чанкам
    qdrant_client.delete_collection(collection_name=COLL_QUESTION_NAME)
    qdrant_client.create_collection(
        collection_name=COLL_QUESTION_NAME,
        vectors_config=VectorParams(size=vec_size, distance=Distance.COSINE),
    )

    logger.info("Collections created successfully")

    for file_name in files:
        chunks = file_to_chunks(file_name, sep, chunk_size, chunk_overlap)

        questions_for_chunk = []
        for chunk in chunks:
            questions_for_chunk += [get_questions_for_chunk(chunk)]
            logger.info("Questions was created successfully")

        # помещаем чанки в векторную БД
        save_chunks(bi_encoder, chunks, file_name, questions_for_chunk)
        logger.info("chunks saved successfully")


def vec_search(bi_encoder, query, n_top_cos, n_top_cos_question):
    # Кодируем запрос в вектор
    query_emb = str_to_vec(bi_encoder, query)

    # Поиск в БД по документам
    search_result = qdrant_client.search(
        collection_name=COLL_NAME,
        query_vector=query_emb,
        limit=n_top_cos,
        with_vectors=False,
    )

    # Поиск в БД по вопросам и мэтчинг соотвествующих чанков
    search_questions_result = qdrant_client.search(
        collection_name=COLL_QUESTION_NAME,
        query_vector=query_emb,
        limit=n_top_cos_question,
        with_vectors=False,
    )

    top_chunks = [x.payload["chunk"] for x in search_result]
    top_question_chunks = [x.payload["chunk"] for x in search_questions_result]
    top_files = list(set([x.payload["file"] for x in search_result]))
    top_question_files = list(set([x.payload["file"] for x in search_questions_result]))

    return top_chunks + top_question_chunks, top_files + top_question_files
