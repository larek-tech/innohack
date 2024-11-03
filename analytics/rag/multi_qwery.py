import json
import requests
import ollama
from loguru import logger


def get_qwestions(prompt: str) -> list[str]:
    
    system_prompt = """Ты — помощник по переформулированию вопросов. 
    Твоя задача — взять вопрос пользователя и переформулировать его 3 разa так, 
    чтобы он звучал более четко и понятно, сохраняя при этом исходный смысл.
    Формат вывода: вопросы разделённые строчкой. Не добавляй никакой текст кроме вопросов.
    """

    content = f""""Пожалуйста, переформулируй следующий вопрос 5 раз так, 
    чтобы он звучал более ясно и лаконично, сохраняя при этом его первоначальный  
    смысл: {prompt}"""


    url = "https://mts-aidocprocessing-case.olymp.innopolis.university/generate"
    data = {
        "prompt": content,
        "apply_chat_template": True,
        "system_prompt": system_prompt,
        "max_tokens": 512,
        "n": 1,
        "temperature": 1,
    }

    headers = {"Content-Type": "application/json"}

    response = requests.post(url, data=json.dumps(data), headers=headers)

    if response.status_code == 200:
        logger.info(response.json())
        questions = [question for question in str(response.json()).split("\n")]
        return questions
    else:
        return f"Error: {response.status_code} - {response.text}"


if __name__ == "__main__":
    prompt = "Кто подписал финансовый отчёт в 2022 году?"
    print(get_qwestions(prompt))