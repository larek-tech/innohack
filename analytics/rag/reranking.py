from loguru import logger

from FlagEmbedding import FlagReranker

reranker = FlagReranker('BAAI/bge-reranker-v2-m3', use_fp16=True)

def rerank_documents(query, documents):
    ranked_results = []
    for doc in documents:
        score = reranker.compute_score([query, doc])
        logger.info(score)
        ranked_results.append((doc, score))

    ranked_results.sort(key=lambda x: x[1], reverse=True)


    return [docs[0] for docs in ranked_results[:7]]