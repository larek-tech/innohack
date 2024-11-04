import typing as t
from img2table.document import Image, PDF
import numpy as np
import cv2
from io import BytesIO
from dataclasses import dataclass
from pdf2image import convert_from_path

from img2table.ocr import EasyOCR
from img2table.tables.objects.extraction import ExtractedTable

import pandas as pd

@dataclass
class Report:
    report_table: pd.DataFrame
    report_code: str
    current_year: str
    curernt_title: str

CODE_KEYS = ["Форма по ОКУД", "Форма"]
CODE_VALUES = ["0710001", "0710002", "0710004", "0710005"]
HEADER_KEY = ["Наименование показателя"]

class ParseTablePipeline:
    def __init__(self, use_gpu: bool = False):
        self.ocr = EasyOCR(lang=["ru"], kw={"gpu": use_gpu})

    def init_doc(self, doc_path: str = None):
        self.document = Image(
            src=doc_path,
            detect_rotation=True
        )

        self.preprocess_images()
    

    def preprocess_images(self):
        images: t.List[np.ndarray] = []
        for image_array in self.document.images:
            fixed_image = self.whiten_stamp(image_array)
            images.append(fixed_image)
        
        self.document.images = images
    
    def whiten_stamp(self, image_array: t.List[np.ndarray]):
        img_cv2 = cv2.cvtColor(image_array, cv2.COLOR_RGB2BGR)
        hsv = cv2.cvtColor(img_cv2, cv2.COLOR_BGR2HSV)
        lower_blue = np.array([90, 50, 50])
        upper_blue = np.array([130, 255, 255])

        mask = cv2.inRange(hsv, lower_blue, upper_blue)
        img_cv2[mask != 0] = [255, 255, 255]

        return img_cv2
    
    def postprocess_tables(self, tables: t.List[ExtractedTable]) -> t.List[Report]:
        report_titles = {
            "БУХГАЛТЕРСКИЙ БАЛАНС": None,
            "ОТЧЕТ О ФИНАНСОВЫХ РЕЗУЛЬТАТАХ": None,
            "ОТЧЕТ ОБ ИЗМЕНЕНИЯХ КАПИТАЛА": None,
            "ОТЧЕТ О ДВИЖЕНИИ ДЕНЕЖНЫХ СРЕДСТВ": None,
        }

        required_columns = {
            "БУХГАЛТЕРСКИЙ БАЛАНС": 4,
            "ОТЧЕТ О ФИНАНСОВЫХ РЕЗУЛЬТАТАХ": 3,
            "ОТЧЕТ ОБ ИЗМЕНЕНИЯХ КАПИТАЛА": 8,
            "ОТЧЕТ О ДВИЖЕНИИ ДЕНЕЖНЫХ СРЕДСТВ": 3,
        }

        combined_df = pd.DataFrame()
        current_year = None
        current_title = None
        
        for table in tables:
            title = table.title
            # Извлекаем год из заголовка
            if title:
                year = self.extract_year_from_title(title)
                if year:
                    current_year = year
                # Определяем тип таблицы с учетом опечаток
                for key in report_titles.keys():
                    if self.is_similar_title(key, title):
                        current_title = key
                        report_titles[key] = title
                        break

            # Проверяем, можем ли склеить текущий DataFrame
            if current_title and self.is_data_table(table.df, required_columns[current_title]):
                combined_df = pd.concat([combined_df, table.df], ignore_index=True)

        # Создаем отчеты на основе собранной информации
        reports = []
        for key in report_titles.keys():
            if report_titles[key]:
                reports.append(Report(report_table=combined_df, report_code=key, 
                                      curernt_title=current_title, current_year=current_year))

        return reports

    def extract_year_from_title(self, title: str) -> str:
        # Находим год в заголовке, используя регулярные выражения
        import re
        match = re.search(r'\b(\d{4})\b', title)
        return match.group(1) if match else None

    def is_similar_title(self, key: str, title: str) -> bool:
        # Проверяем, совпадает ли часть заголовка с ключом
        if title in key: return True
        key_parts = key.split()
        title_parts = title.split()
        return any(part in title_parts for part in key_parts)

    def is_data_table(self, df: pd.DataFrame, expected_columns: int) -> bool:
        # Проверяем, соответствует ли DataFrame ожидаемому количеству столбцов
        return not df.empty and df.shape[1] >= expected_columns
    
    def __call__(self, doc_path: str):
        self.init_doc(doc_path)

        extracted_tables = self.document.extract_tables(
            ocr=self.ocr,
            implicit_rows=True,
            implicit_columns=True,
            borderless_tables=True,
            min_confidence=50
        )
        print('Tables extracted:', extracted_tables)

        # post_processed_reports: t.List[Report] = self.postprocess_tables(extracted_tables)

        # table_dfs = [table.df for table in extracted_tables]
        
        return extracted_tables
