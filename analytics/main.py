import logging
from concurrent import futures

import grpc

from analytics import analytics_pb2, analytics_pb2_grpc

from rag.rag_model import RagClient


rag = RagClient()

class Analytics(analytics_pb2_grpc.AnalyticsServicer):
    def GetCharts(self, request: analytics_pb2.Params, context: grpc.ServicerContext):
        return analytics_pb2.Report()

    def GetDescriptionStream(
        self, request: analytics_pb2.Params, context: grpc.ServicerContext
    ):
        res = analytics_pb2.Report()

        response = rag.generate(res) 

        for token in response:
            res.description += f"{token} "
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
