from pathlib import Path

from db import files_to_vecdb
from utils.bi_encode import get_bi_encoder

from loguru import  logger

def main():
    dir_path = Path(__file__).parent / "data"

    file_paths = list(dir_path.glob('*'))
    logger.info(dir_path)
    bi_encoder, vec_size = get_bi_encoder("cointegrated/LaBSE-en-ru")
    logger.info(bi_encoder)
    files_to_vecdb(
        files=file_paths,
        bi_encoder=bi_encoder,
        vec_size=768,
        sep="\n",
        chunk_size=2000,
        chunk_overlap=60,
    )


if __name__ == "__main__":
    main()