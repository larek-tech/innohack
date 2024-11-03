
import requests
import json

from loguru import logger

from rag.utils.bi_encode import get_bi_encoder 
from rag.db import QdrantBase
from rag.config import BI_ENCODE_NAME, QDRANT_HOST, QDRANT_PORT

class LLMClient:
    def __init__(self, model="meta-llama/Llama-3.2-11B-Vision-Instruct"):
        self.model = model
        self.api_url = "https://mts-aidocprocessing-case.olymp.innopolis.university/generate"
        self.bi_encoder , self.vect_dim = get_bi_encoder(BI_ENCODE_NAME)
        self.qdrant = QdrantBase(QDRANT_HOST, QDRANT_PORT)
        self.n_top_cos = 8

    def get_response(self, prompt):

        top_chunks, top_files = self.qdrant.vec_search(self.bi_encoder, prompt, self.n_top_cos)
        top_chunks_join = '\n'.join(top_chunks)
        logger.info(top_chunks)

        data = {
            "prompt": 
            f"""
            Используй только следующий контекст, чтобы очень кратко ответить на вопрос в конце.
            Не пытайся выдумывать ответ.
            Контекст:
            ===========
            {top_chunks_join}
            ===========
            Вопрос:
            ===========
            {prompt}
            """,
            "apply_chat_template": True,
            "system_prompt": """ 
            Ты — помощник по анализу финансовых отчетов. Твоя задача — предоставлять
              точные и полезные ответы на вопросы, связанные с финансовыми данными, отчетами и анализом.""",
            "max_tokens": 2048,
            "n": 1,
            "temperature": 0,
            "stream" : True
        }

        headers = {"Content-Type": "application/json"}

        # response = requests.post(self.api_url, data=json.dumps(data), headers=headers)

        response = requests.post(self.api_url, data=json.dumps(data), headers=headers, stream=True)

        if response.status_code == 200:
            for line in response.iter_lines():
                if line:
                    decoded_line = line.decode('utf-8')
                    yield decoded_line
        else:
            return f"Error: {response.status_code} - {response.text}"

        # if response.status_code == 200:
        #     return response.json()
        # else:
        #     return f"Error: {response.status_code} - {response.text}"

# Пример использования
if __name__ == "__main__":
    client = LLMClient()
    
    prompt = "Резервный капитал 2011 год?"
    try:
        # Получаем генератор ответов
        response_generator = client.get_response(prompt)
        
        # Итерируемся по результатам
        for response in response_generator:
            print("Ответ LLM:", response)
    except Exception as e:
        print("Произошла ошибка:", e)
