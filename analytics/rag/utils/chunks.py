import fitz

from langchain_community.document_loaders import TextLoader, Docx2txtLoader, PyPDFLoader
from langchain.text_splitter import RecursiveCharacterTextSplitter

from loguru import logger


def file_to_chunks(file_name, sep, chunk_size, chunk_overlap):
    file_ext = str(file_name).split(".")[-1]
    logger.info(file_name)

    if file_ext == "txt":
        loader = TextLoader(file_name, encoding="utf-8")
        file = loader.load()
        content = file[0].page_content

        text_splitter = RecursiveCharacterTextSplitter(
            separators=sep,
            chunk_size=200,
            chunk_overlap=20,
            length_function=len,
            is_separator_regex=False,
            add_start_index=False,
        )

    elif file_ext == "pdf":
        document = fitz.open(file_name)
        content = ""

        for page in document:
            content += page.get_text()

        logger.info(content)

        document.close()

        text_splitter = RecursiveCharacterTextSplitter(
            separators=sep,
            chunk_size=chunk_size,
            chunk_overlap=chunk_overlap,
            length_function=len,
            is_separator_regex=False,
            add_start_index=False,
        )

    else:
        return None

    chunks = text_splitter.split_text(content)

    logger.info(chunks)

    return chunks
