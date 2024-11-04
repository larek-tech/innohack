from pymongo import MongoClient
import json

mongo = MongoClient("mongodb://10.0.1.80:27017/data", timeoutMS=30000**2)

summary_col = mongo.get_database("data").get_collection("report_summary")

def send_dict(summary_dict: dict):
    has_summary_dict = len([s for s in summary_col.find({})]) > 0
    if not has_summary_dict:
        summary_col.insert_one(summary_dict)
        print("Sent")

def main():
    path = "/home/hope/Hope/Innohack/summary_json.json"
    f = open(path)
    data = json.load(f)

    send_dict(data)

main()