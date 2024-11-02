import typing as t
from img2table.document import Image, PDF
import numpy as np
import cv2
from dataclasses import dataclass

from img2table.ocr import EasyOCR
from img2table.tables.objects.extraction import ExtractedTable

import pandas as pd


@dataclass
class Report:
    report_table: pd.DataFrame
    report_code: str


CODE_KEYS = ["Форма по ОКУД", "Форма"]
CODE_VALUES = ["0710001", "0710002", "0710003", "0710004", "0710005"]


class ParseTablePipeline:
    def __init__(self, doc_path: str, mode: str = "image", use_gpu: bool = False):
        self.ocr = EasyOCR(lang=["ru"], kw={"gpu": use_gpu})
        self.mode = mode
        if mode == "image":
            self.document = Image(src=doc_path, detect_rotation=True)
        else:
            self.document = PDF(
                src=doc_path,
                pages=None,
                detect_rotation=False,
                pdf_text_extraction=True,
            )

        self.preprocess_images()

    def preprocess_images(self):
        images: t.List[np.ndarray] = []
        for image_array in enumerate(self.document.images):
            fixed_image = self.whiten_stamp(image_array)
            images.append(fixed_image)

        self.document.images = images

    def whiten_stamp(image_array: t.List[np.ndarray]):
        img_cv2 = cv2.cvtColor(image_array, cv2.COLOR_RGB2BGR)

        # Преобразуем изображение в цветовое пространство HSV
        hsv = cv2.cvtColor(img_cv2, cv2.COLOR_BGR2HSV)

        # Определяем диапазон синего цвета в HSV
        lower_blue = np.array([90, 50, 50])
        upper_blue = np.array([130, 255, 255])

        # Создаем маску для синих оттенков
        mask = cv2.inRange(hsv, lower_blue, upper_blue)

        # Заменяем синий цвет на белый
        img_cv2[mask != 0] = [255, 255, 255]

        return img_cv2

    def postprocess_tables(tables: t.List[ExtractedTable]) -> t.List[Report]:
        found_code = None
        for table in tables:
            df = table.df
            if found_code is None:
                for value in CODE_VALUES:
                    if (df == value).any().any():
                        found_code = value
                        break
                    else:
                        for key in CODE_KEYS:
                            for i, row in df.iterrows():
                                if key in row.values:
                                    key_index = row[row == key].index[0]
                                    next_col_index = df.columns.get_loc(key_index) + 1
                                    if next_col_index < len(df.columns):
                                        next_value = df.iloc[i, next_col_index]
                                        found_code = next_value
                                        break
                            if found_code is not None:
                                break
                if found_code is not None:
                    continue
            else:
                pass

        return

    def __call__(self):
        extracted_tables = self.document.extract_tables(
            ocr=self.ocr,
            implicit_rows=True,
            implicit_columns=True,
            borderless_tables=True,
            min_confidence=50,
        )

        post_processed_reports: t.List[Report] = []
        if self.mode == "images":
            post_processed_reports.append(self.postprocess_tables(extracted_tables))
        else:
            for page in extracted_tables:
                post_processed_reports.append(self.postprocess_tables(page))

        return post_processed_reports


if __name__ == "__main__":

    image_path = "image_balance_sheet.jpg"

    table_parse_pipeline = ParseTablePipeline(doc_path=image_path, mode="image")

    parse_result = table_parse_pipeline()
