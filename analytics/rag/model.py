import requests
import ollama
from loguru import logger

from rag.utils.bi_encode import get_bi_encoder
from rag.db import QdrantBase
from rag.config import BI_ENCODE_NAME, QDRANT_HOST, QDRANT_PORT


def get_streaming_response(response: requests.Response):
    for chunk in response.iter_lines(
        chunk_size=8192, decode_unicode=False, delimiter=b"\0"
    ):
        if chunk:
            yield chunk.decode("utf-8")


def clear_line(n: int = 1) -> None:
    LINE_UP = "\033[1A"
    LINE_CLEAR = "\x1b[2K"
    for _ in range(n):
        print(LINE_UP, end=LINE_CLEAR, flush=True)


class LLMClient:
    def __init__(self, model="meta-llama/Llama-3.2-11B-Vision-Instruct"):
        self.model = model
        self.ollama = ollama.Client("localhost:11434")
        self.bi_encoder, self.vect_dim = get_bi_encoder(BI_ENCODE_NAME)
        self.qdrant = QdrantBase(QDRANT_HOST, QDRANT_PORT)
        self.n_top_cos = 8

    def get_response(self, prompt):

        top_chunks, top_files = self.qdrant.vec_search(
            self.bi_encoder, prompt, self.n_top_cos
        )
        top_chunks_join = "\n".join(top_chunks)
        logger.info(top_chunks)

        data = {
            "prompt": f"""
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
            "stream": True,
        }

        return self.ollama.chat(
            "llama3.2",
            messages=[
                {"role": "system", "content": data["system_prompt"]},
                {"role": "user", "content": data["prompt"]},
            ],
            stream=True,
        )
