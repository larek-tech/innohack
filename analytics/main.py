import logging
from concurrent import futures

import ollama
import grpc

from analytics import analytics_pb2, analytics_pb2_grpc
from process.process import preprocess_xlsx
from process.form_graphs import get_analitics_report
from process.get_report_summary import form_report_description

from rag.rag_model import RagClient


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
        for chunk in self.rag.llm_client.get_response(request.prompt):
            res = analytics_pb2.DescriptionReport()
            res.description = chunk["message"]["content"]
            print(res.description)
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
    # serve()
    rag = RagClient()

    # prompt = "Какой был резервный капитал в 2012 и 2013 годах?"
    # prompt = "Напиши функцию на Python, которая складывает два числа"
    # promt = "Какие были активы компании в 2023 году?"

    # bad_prompts = [
    #     "Какие продукты есть в экосистеме МТС?",
    #     "Сколько стоит отправка СМС в сети МТС?",
    #     "Какие есть тарифы в МТС?",
    #     "Расскажи о самом выгодном предложении в комании",
    #     "Посоветуй, какой тариф выгоднее всего приобрести",
    #     "Напиши код на Rust",
    #     "Реши пример: 2 + 2 = ",
    #     "Продолжи стих: У Лукоморья дуб зеленый",
    #     "Что взять с собой в похож?",
    #     "Где лучше купить сумку?",
    # ]

    good_prompts = [
        "Какой был резервный капитал в 2012 и 2013 годах?",
        "Какие были активы компании в 2023 году?",
    ]

    with open("result_rag.txt", "w") as file:
        for prompt in good_prompts:
            file.write(prompt + "\n")
            file.write(rag.llm_client.get_response(prompt) + "\n")
