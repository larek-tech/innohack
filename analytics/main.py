import logging
from concurrent import futures

import grpc

from analytics import analytics_pb2, analytics_pb2_grpc
from process.process import preprocess_xlsx
from process.form_graphs import get_analitics_report


class Analytics(analytics_pb2_grpc.AnalyticsServicer):
    def __init__(self):
        super().__init__()
        preprocess_xlsx()


    def GetCharts(self, request: analytics_pb2.Params, context: grpc.ServicerContext):
        return get_analitics_report(request)

    def GetDescriptionStream(
        self, request: analytics_pb2.Params, context: grpc.ServicerContext
    ):
        res = analytics_pb2.Report()

        for i in range(10):
            res.description += f"{i} "
            yield res


logging.basicConfig(level=logging.DEBUG)


def serve():
    s = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    analytics_pb2_grpc.add_AnalyticsServicer_to_server(Analytics(), s)
    s.add_insecure_port("[::]:9990")
    print("starting server")
    s.start()
    s.wait_for_termination()


if __name__ == "__main__":
    serve()
