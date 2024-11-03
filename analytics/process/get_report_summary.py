import json
import requests
from analytics import analytics_pb2
import math

from process.load_json import load_json

from process.const import CODE_NAME, MULTYPLIER_NAME

REPORT_CODES = [1100, 1200, 1300, 1400]
   

def get_multy_analysis(graph_data: str) -> str:
    url = "https://mts-aidocprocessing-case.olymp.innopolis.university/generate"
    data = {
        "prompt": graph_data,
        "apply_chat_template": True,
        "system_prompt": "Проанализируй финансовые показатели, кратко опиши тенденции и сделай выводы по показателям.",
        "max_tokens": 2048,
        "n": 1,
        "temperature": 0.8,
    }

    headers = {"Content-Type": "application/json"}

    response = requests.post(url, data=json.dumps(data), headers=headers)

    if response.status_code == 200:
        return response.json()
    else:
        return f"Error: {response.status_code} - {response.text}"

def generate_str_from_json(filtered_records):
    report = []
    
    for indicator, year_data in filtered_records.items():
        for year, value in year_data.items():
            sentence = f"{indicator} в {year} году - {value}."
            report.append(sentence)
    
    return " ".join(report)

def generate_str_from_json_multy(filtered_multipliers):
    report = []
    
    for indicator, year_data in filtered_multipliers.items():
        for year, value in year_data.items():
            sentence = f"{indicator} в {year} году составлял {value}."
            report.append(sentence)
    
    return " ".join(report)

def form_report_description(records: dict, multipliers: dict, start_date: str, end_date: str):
    start_date = int(start_date)
    end_date = int(end_date)

    filtered_records = {}
    for code, years_data in records.items():
        if int(code) in REPORT_CODES:
            filtered_records[CODE_NAME[int(code)]] = {
                year: value for year, value in years_data.items() 
                if start_date <= int(year) <= end_date and not math.isnan(value)
            }
        
    filtered_multipliers = {}
    for name, years_data in multipliers.items():
        filtered_multipliers[MULTYPLIER_NAME[name]] = {year: value for year, value in years_data.items() 
                                                  if start_date <= int(year) <= end_date and not math.isnan(value)}
    
    # data_report = get_multy_analysis(generate_str_from_json(filtered_records))
    multy_report = get_multy_analysis(generate_str_from_json_multy(filtered_multipliers))
    return multy_report

