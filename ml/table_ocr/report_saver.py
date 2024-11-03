import os
from typing import List
from parse_ocr_pipeline import Report
import pandas as pd

class ReportSaver:
    @staticmethod
    def save_reports_to_excel(reports: List[Report], output_dir: str = "."):
        valid_report = next((report for report in reports if report.current_year.startswith("20")), None)
        if valid_report is None:
            raise ValueError("No valid report found with a year starting with '20'.")

        year = valid_report.current_year
        file_name = f"MТС_{year}.xlsx"
        file_path = os.path.join(output_dir, file_name)

        with pd.ExcelWriter(file_path, engine="openpyxl") as writer:
            for idx, report in enumerate(reports):
                sheet_name = f"Форма {idx + 1}"
                report.report_table.to_excel(writer, sheet_name=sheet_name, index=False)

        print(f"Reports saved to {file_path}")