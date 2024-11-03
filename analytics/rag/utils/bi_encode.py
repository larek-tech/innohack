import torch

from sentence_transformers import SentenceTransformer
from sentence_transformers.models import Pooling, Transformer


def get_bi_encoder(bi_encoder_name):
    raw_model = Transformer(model_name_or_path=f'{bi_encoder_name}')

    bi_encoder_dim = raw_model.get_word_embedding_dimension()
    
    pooling_model = Pooling(
        bi_encoder_dim,
        pooling_mode_cls_token = False,
        pooling_mode_mean_tokens = True
    )
    bi_encoder = SentenceTransformer(
        modules = [raw_model, pooling_model],
        device = "cuda" if torch.cuda.is_available() else "cpu"
    )
    
    return bi_encoder, bi_encoder_dim

def str_to_vec(bi_encoder, text):
    embeddings = bi_encoder.encode(
        text,
        convert_to_tensor = True,
        show_progress_bar = False
    )
    return embeddings