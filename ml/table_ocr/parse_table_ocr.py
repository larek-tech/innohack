
import tempfile
from pdf2image import convert_from_path
import numpy as np
import os
from PIL import Image as PILImage

from parse_ocr_pipeline import ParseTablePipeline
from report_saver import ReportSaver

def extract_tables(pdf_paths: list[str]):
    result = []
    table_ocr_pipeline = ParseTablePipeline()
    report_saver = ReportSaver()
    for pdf in pdf_paths:
        pdf_tables = []
        reports = []
        with tempfile.TemporaryDirectory() as td:
            images = convert_from_path(pdf)

            for idx, image in enumerate(images):
                pix = np.array(image)
                f_name = os.path.join(td, f'page_{idx}.jpg')
                im = PILImage.fromarray(pix)
                im.save(f"{f_name}")

                extr_tables = table_ocr_pipeline(f_name)
                pdf_tables.append(extr_tables)
                reports.extend(table_ocr_pipeline.postprocess_tables(extr_tables))

                if len(reports) == 4:
                    break
        
        print(reports)
        report_saver.save_reports_to_excel(reports)

        result.append(reports)
    return result
  
def main():
    local_excel_dir = 'pdf_data/'
    pdf_paths = [local_excel_dir + path for path in os.listdir(local_excel_dir)]
    res = extract_tables(pdf_paths)
    print(res)

main()