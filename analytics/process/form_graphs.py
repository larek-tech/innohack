import json
from analytics import analytics_pb2, analytics_pb2_grpc


def load_json():
    records = {}
    multipliers = {}

    with open('records.json', 'r') as file:
        records = json.load(file)
    with open('multipliers.json', 'r') as file:
        multipliers = json.load(file)
    return records, multipliers

def form_graph_info(records: dict, multipliers: dict, request: analytics_pb2.Params):

    pass