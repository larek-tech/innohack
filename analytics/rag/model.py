import requests
import json
from ollama import Client

import random

from loguru import logger

from rag.utils.bi_encode import get_bi_encoder
from rag.db import vec_search, define_question_topic, vec_search_qwery
from rag.multi_qwery import get_qwestions
from rag.reranking import rerank_documents
from rag.utils.variants_of_answer import vars


class LLMClient:
    def __init__(self, model="llama3.2"):  # meta-llama/Llama-3.2-11B-Vision-Instruct
        self.model = model
        self.api_url = (
            "https://mts-aidocprocessing-case.olymp.innopolis.university/generate"
        )
        self.bi_encoder, self.vect_dim = get_bi_encoder("cointegrated/LaBSE-en-ru")

        self.n_top_cos = 2
        self.n_top_cos_question = 3

    def get_response(self, prompt):

        # Classification stage, checking for compliance with the topic
        classificator_response = define_question_topic(prompt)
        logger.info(classificator_response)
        if "нет" in classificator_response.lower():
            return random.choice(vars)

        # MultiQwestion stage
        # questions = get_qwestions(prompt)
        questions = [prompt]

        # Find chunks stage
        top_chunks = vec_search(
            self.bi_encoder,
            prompt,
            self.n_top_cos,
        )
        # logger.info(top_chunks)
        top_query_chank = []
        for question in questions:
            top_query_chank += vec_search_qwery(
                self.bi_encoder,
                question,
                self.n_top_cos_question,
            )
        logger.info(top_query_chank)

        # ReRanking stage
        top_merge_chunks = top_chunks + top_query_chank

        # logger.info(top_merge_chunks)

        # reranked_chunks = rerank_documents(
        #     prompt,
        #     top_merge_chunks,
        # )

        reranked_chunks_joint = "\n".join(top_merge_chunks)
        # Generate Stage
        content = f"""
            Используй только следующий контекст, чтобы ответить на вопрос.
            Не пытайся выдумывать ответ.
            Не отвечай на вопросы, не связанные с финансами. Численные ответы пиши в тысячах рублей.
            Контекст:
            ===========
            {reranked_chunks_joint}
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

        # client = Client(host="http://10.92.9.223:11434")

        # response = client.chat(
        #     model="hf.co/IlyaGusev/saiga_nemo_12b_gguf:Q8_0",
        #     messages=[
        #         {
        #             "role": "user",
        #             "content": content,
        #         },
        #     ],
        #     options={
        #         "system_prompt": system_prompt,
        #         "max_tokens": max_tokens,
        #         "temperature": temperature,
        #     },
        # )

        # return response["message"]["content"]

        data = {
            "prompt": content,
            "apply_chat_template": True,
            "system_prompt": system_prompt,
            "max_tokens": 512,
            "n": 1,
            "temperature": 0.4,
        }

        headers = {"Content-Type": "application/json"}

        response = requests.post(self.api_url, data=json.dumps(data), headers=headers)

        if response.status_code == 200:
            logger.info(response.json())
            return response.json()
        else:
            return f"Error: {response.status_code} - {response.text}"


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
