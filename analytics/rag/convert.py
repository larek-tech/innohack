import json
import random
from loguru import logger

from rag.const_data import CODE_NAME




def format_data(data):
    result = ""
    
    # Список случайных вводных фраз
    intro_phrases = [
        "Давайте рассмотрим показатель",
        "Обратите внимание на показатель",
        "Интересно отметить, что показатель",
        "Следующий показатель",
        "Показатель"
    ]
    
    # Список вариативных фраз о значимости показателя
    significance_phrases = [
        "является ключевым элементом анализа данных.",
        "играет важную роль в понимании тенденций.",
        "представляет собой значимый аспект исследования.",
        "является важным индикатором для принятия решений.",
        "служит основой для дальнейшего анализа."
    ]
    
    # Список случайных заключительных фраз
    conclusion_phrases = [
        "Эти данные могут быть полезны для дальнейшего анализа.",
        "Эти значения помогут в принятии обоснованных решений.",
        "Эти цифры дают представление о тенденциях.",
        "Эти данные могут служить основой для будущих исследований."
    ]
    
    # Извлекаем ключ и значения
    keys = list(data.keys())
    for key in keys:
        if key == "_id":
            continue
        values = data[key]
        logger.info(key)
        
        # Формируем строку с более развернутым текстом и случайными фразами
        intro = random.choice(intro_phrases)
        significance = random.choice(significance_phrases)
        conclusion = random.choice(conclusion_phrases)
        
        result += f"{intro} '{CODE_NAME[key]}' {significance} "
        result += f"В течение следующих лет его значения изменялись следующим образом: "
        
        for year, value in values.items():
            result += f"в {year} году он составил {value}. "
        
        result += f"{conclusion} "
        result += "\n"  # Добавляем новую строку для разделения показателей
    
    return result

if __name__ == "__main__":
    js_ph = "records.json"

    output_file = "formatted_output.txt"

    with open(js_ph, 'r', encoding='utf-8') as file:
        json_data = json.load(file)
    formatted_data = format_data(json_data)

    with open(output_file, 'w', encoding='utf-8') as file:
        file.write(formatted_data)
    
    print(f"Данные успешно сохранены в файл: {output_file}")