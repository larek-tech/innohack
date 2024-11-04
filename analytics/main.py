import logging
from concurrent import futures

import grpc

from analytics import analytics_pb2, analytics_pb2_grpc
from process.process import preprocess_xlsx
from process.form_graphs import get_analitics_report
from process.get_report_summary import form_report_description

from rag.rag_model import RagClient

import random


def random_chunking(text: str, min_chunk_size: int = 1, max_chunk_size: int = 5):
    tokens = text.split()

    start_index = 0

    while start_index < len(tokens):
        chunk_size = random.randint(min_chunk_size, max_chunk_size)

        end_index = min(start_index + chunk_size, len(tokens))

        yield " ".join(tokens[start_index:end_index])

        start_index = end_index


class Analytics(analytics_pb2_grpc.AnalyticsServicer):
    def __init__(self):
        super().__init__()
        # Data preporation - USE ONLY ONCE
        # preprocess_xlsx()

        self.rag = RagClient()

    def GetCharts(self, request: analytics_pb2.Filter, context: grpc.ServicerContext):
        return get_analitics_report(request)

    def GetDescriptionStream(
        self, request: analytics_pb2.Params, context: grpc.ServicerContext
    ):
        res = analytics_pb2.DescriptionReport()
        res.description = self.rag.llm_client.get_response(request.prompt)
        yield res


logging.basicConfig(level=logging.INFO)


def serve():
    s = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    analytics_pb2_grpc.add_AnalyticsServicer_to_server(Analytics(), s)
    s.add_insecure_port("[::]:9990")
    s.start()
    print("starting server")
    s.wait_for_termination()


if __name__ == "__main__":
    serve()
