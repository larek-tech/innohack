from pymongo import MongoClient

mongo = MongoClient("mongodb://46.138.243.191:27017/data", timeoutMS=30000**2)
records_col = mongo.get_database("data").get_collection("records")
multipliers_col = mongo.get_database("data").get_collection("multipliers")
summary_col = mongo.get_database("data").get_collection("report_summary")

def load_json():
    records = records_col.find_one({}, {"_id": 0})
    multipliers = multipliers_col.find_one({}, {"_id": 0})
    report_summary = summary_col.find_one({}, {"_id": 0})
    
    return records, multipliers, report_summary