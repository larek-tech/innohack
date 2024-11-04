import logging
from concurrent import futures

import grpc

from analytics import analytics_pb2, analytics_pb2_grpc
from process.process import preprocess_xlsx
from process.form_graphs import get_analitics_report

from rag.rag_model import RagClient


class Analytics(analytics_pb2_grpc.AnalyticsServicer):
    def __init__(self):
        super().__init__()
        preprocess_xlsx()

        self.rag = RagClient()

    def GetCharts(self, request: analytics_pb2.Params, context: grpc.ServicerContext):
        return get_analitics_report(request)

    def GetDescriptionStream(
        self, request: analytics_pb2.Params, context: grpc.ServicerContext
    ):
        for token in self.rag.llm_client.get_response(request.prompt):
            res = analytics_pb2.Report()
            res.description = token
            yield res


logging.basicConfig(level=logging.INFO)


def serve():
    s = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    analytics_pb2_grpc.add_AnalyticsServicer_to_server(Analytics(), s)
    s.add_insecure_port("[::]:9990")
    print("starting server")
    s.start()
    s.wait_for_termination()


if __name__ == "__main__":
    # serve()
    rag = RagClient()

    # prompt = "Какой был резервный капитал в 2012 и 2013 годах?"
    prompt = "Напиши функцию на Python, которая складывает два числа"
    # promt = "Какие были активы компании в 2023 году?"

    print(rag.llm_client.get_response(prompt))
