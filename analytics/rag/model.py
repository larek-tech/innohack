import requests
import json
import ollama

import random

from loguru import logger

from rag.utils.bi_encode import get_bi_encoder
from rag.db import vec_search, define_question_topic


class LLMClient:
    def __init__(self, model="llama3.1"):  # meta-llama/Llama-3.2-11B-Vision-Instruct
        self.model = model
        # self.api_url = (
        #     "https://mts-aidocprocessing-case.olymp.innopolis.university/generate"
        # )
        self.bi_encoder, self.vect_dim = get_bi_encoder("cointegrated/LaBSE-en-ru")

        self.n_top_cos = 2
        self.n_top_cos_question = 7

    def get_response(self, prompt):

        # classificator_response = define_question_topic(prompt)
        # if "нет" in classificator_response.lower():
        #     return random.choice()

        top_chunks, top_files = vec_search(
            self.bi_encoder, prompt, self.n_top_cos, self.n_top_cos_question
        )
        top_chunks_join = "\n".join(top_chunks)

        logger.info(top_chunks)

        content = f"""
            Используй только следующий контекст, чтобы ответить на вопрос.
            Не пытайся выдумывать ответ.
            Не отвечай на вопросы, не связанные с финансами.
            Не отвечай на вопросы, не связанные с финансами.
            Контекст:
            ===========
            {top_chunks_join}
            ===========
            Вопрос:
            ===========
            {prompt}
            """

        system_prompt = """ 
            Ты — помощник по анализу финансовых отчетов. Твоя задача — предоставлять
            точные и полезные ответы на вопросы, связанные с финансовыми данными, отчетами и анализом.
            Не отвечай на вопросы, не связанные с финансами и бухгалтерией.
        """

        max_tokens = 512
        temperature = 0.8

        response = ollama.chat(
            model="llama3.1",
            messages=[
                {
                    "role": "user",
                    "content": content,
                },
            ],
            options={
                "system_prompt": system_prompt,
                "max_tokens": max_tokens,
                "temperature": temperature,
            },
        )

        return response["message"]["content"]


# Пример использования
if __name__ == "__main__":
    client = LLMClient()

    # prompt = "Какой был резервный капитал в 2012 и 2013 годах?"
    # prompt = "Напиши функцию на Python, которая складывает два числа"
    prompt = "Какие были активы компании в 2023 году?"

    """
     ̈ КакиепродуктыестьвэкосистемеМТС  ̈ СколькостоитотправкаСМСвсетиМТС  ̈ КакиеестьтарифывМТС?

    """
    try:
        response = client.get_response(prompt)
        print()
        print("Ответ LLM:", response)
    except Exception as e:
        print(e)
