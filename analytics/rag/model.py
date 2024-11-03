import requests
import json

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
        self.api_url = (
            "https://mts-aidocprocessing-case.olymp.innopolis.university/generate"
        )
        self.bi_encoder, self.vect_dim = get_bi_encoder(BI_ENCODE_NAME)
        self.qdrant = QdrantBase(QDRANT_HOST, QDRANT_PORT)
        self.n_top_cos = 8

    def clear_line(self, n: int = 1) -> None:
        LINE_UP = "\033[1A"
        LINE_CLEAR = "\x1b[2K"
        for _ in range(n):
            print(LINE_UP, end=LINE_CLEAR, flush=True)

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

        headers = {"Content-Type": "application/json"}

        # response = requests.post(self.api_url, data=json.dumps(data), headers=headers)

        session = requests.Session()
        response = session.post(
            self.api_url, data=json.dumps(data), headers=headers, stream=True
        )

        num_printed_lines = 0
        for h in get_streaming_response(response):
            clear_line(num_printed_lines)
            num_printed_lines = 0
            for i, line in enumerate(h):
                num_printed_lines += 1
                yield f"{line!r}"
                print(f"Beam candidate {i}: {line!r}", flush=True)

        # if response.status_code == 200:
        #     return response
        # else:
        #     return f"Error: {response.status_code} - {response.text}"


# from typing import Iterable, List


# def post_http_request(
#     prompt: str, api_url: str, n: int = 1, stream: bool = False
# ) -> requests.Response:
#     # headers = {"User-Agent": "Test Client"}
#     pload = {
#         "prompt": prompt,
#         # "n": n,
#         # "use_beam_search": True,
#         "temperature": 0.0,
#         "max_tokens": 16,
#         "stream": stream,
#     }
#     # headers=headers,
#     response = requests.post(api_url, json=pload, stream=stream)
#     return response


# def get_streaming_response(response: requests.Response) -> Iterable[List[str]]:
#     for chunk in response.iter_lines(
#         chunk_size=8192, decode_unicode=False, delimiter=b"\0"
#     ):
#         if chunk:
#             yield chunk.decode("utf-8")


# def get_response(response: requests.Response) -> List[str]:
#     data = json.loads(response.content)
#     output = data["text"]
#     return output


# if __name__ == "__main__":
#     parser = argparse.ArgumentParser()
#     parser.add_argument(
#         "--host",
#         type=str,
#         default="mts-aidocprocessing-case.olymp.innopolis.university",
#     )
#     parser.add_argument("--port", type=int, default=443)
#     parser.add_argument("--n", type=int, default=4)
#     parser.add_argument("--prompt", type=str, default="San Francisco is a")
#     parser.add_argument("--stream", action="store_true")
#     args = parser.parse_args()
#     prompt = "Привет!"
#     api_url = f"https://mts-aidocprocessing-case.olymp.innopolis.university/generate"
#     n = args.n
#     stream = args.stream

#     print(f"Prompt: {prompt!r}\n", flush=True)
#     response = post_http_request(prompt, api_url, n, stream)
#     stream = True
#     if stream:
#         num_printed_lines = 0
#         for h in get_streaming_response(response):
#             clear_line(num_printed_lines)
#             num_printed_lines = 0
#             for i, line in enumerate(h):
#                 num_printed_lines += 1
#                 print(f"Beam candidate {i}: {line!r}", flush=True)
#     else:
#         output = get_response(response)
#         for i, line in enumerate(output):
#             print(f"Beam candidate {i}: {line!r}", flush=True)
