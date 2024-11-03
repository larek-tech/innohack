import uuid

from qdrant_client import QdrantClient
from qdrant_client.http.models import Distance, VectorParams, PointStruct

from rag.utils.bi_encode import str_to_vec
from rag.utils.chunks import file_to_chunks


class QdrantBase:

    COLL_NAME = "mts_test"

    def __init__(
            self, 
            host: str,
            port: int):

        self.qdrant_client = QdrantClient(host, port=port)


    # Помещаем чанки и доп. информаицю в векторую БД
    def save_chunks(self, bi_encoder, chunks, file_name):
        # Конвертируем чанки в векитора
        chunk_embeddings = str_to_vec(bi_encoder, chunks)

        # Содаем объект(ы) для БД
        points = []
        for i in range(len(chunk_embeddings)):
            point = PointStruct(
                id=str(uuid.uuid4()), # генерируем GUID
                vector = chunk_embeddings[i], 
                payload={'file': file_name, 'chunk': chunks[i]}
            )
            points.append(point)
        
        # Сохраняем вектора в БД
        operation_info = self.qdrant_client.upsert(
            collection_name = self.COLL_NAME,
            wait = True,
            points = points
        )
        
        return operation_info

    def files_to_vecdb(self, 
                       files, 
                       bi_encoder, 
                       vec_size, 
                       sep, 
                       chunk_size, 
                       chunk_overlap
    ):

        if self.COLL_NAME not in self.qdrant_client.get_collections():
            self.qdrant_client.delete_collection(collection_name=self.COLL_NAME)
            self.qdrant_client.create_collection(
                collection_name = self.COLL_NAME,
                vectors_config = VectorParams(size=vec_size, distance=Distance.COSINE),
            )
        

        for file_name in files:
            chunks = file_to_chunks(file_name, sep, chunk_size, chunk_overlap)
            # помещаем чанки в векторную БД
            self.save_chunks(bi_encoder, chunks, file_name)


    def vec_search(self, bi_encoder, query, n_top_cos):
        # Кодируем запрос в вектор
        query_emb = str_to_vec(bi_encoder, query)

        # Поиск в БД
        search_result = self.qdrant_client.search(
            collection_name = self.COLL_NAME,
            query_vector = query_emb,
            limit = n_top_cos,
            with_vectors = False
        )
        
        top_chunks = [x.payload['chunk'] for x in search_result]
        top_files = list(set([x.payload['file'] for x in search_result]))
        
        return top_chunks, top_files
