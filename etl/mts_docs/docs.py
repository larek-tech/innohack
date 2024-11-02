import httpx
import asyncio
from bs4 import BeautifulSoup
import re
import pathlib
import os
from tqdm import tqdm

BASE_URL = "https://moskva.mts.ru/about/investoram-i-akcioneram/korporativnoe-upravlenie/raskritie-informacii/godovaya-otchetnost"
# FILE_NAME = "page.html"
# FILE_NAME = "vipusk-cennih-bumag.html"
# FILE_NAME = "vipush-cfa.html"
# FILE_NAME = "soobcheniya.html"
# FILE_NAME = "sushhestvennie-fakti.html"
# FILE_NAME = "ezhekvartalnie-otcheti.html"
# FILE_NAME = "otchety-emitenta-emissionnyh-cennyh-bumag.html"
# FILE_NAME = "spiski-affilirovannih-lic.html"
FILE_NAME = "insajderskaya-informacii-pao-mts.html"
DATA_PATH = pathlib.Path(__file__).parents[2] / "data" / "mts" / "docs"


async def get_page(client: httpx.AsyncClient, url: str) -> str:
    response = await client.get(url)
    if response.status_code != 200:
        return ""
    print(response)
    return str(response.content)


def extract_date_from_string(date_string):
    """
    Extracts the day, month, and year from a date string in the format "dd.mm.yyyy".

    Args:
        date_string (str): The input date string.

    Returns:
        dict: A dictionary containing the extracted day, month, and year.
    """
    # Define the regex pattern to match the date
    pattern = r"(\d{2})\.(\d{2})\.(\d{4})"
    # Use the re.search() function to find the match
    match = re.search(pattern, date_string)

    if match:
        # Extract the date components from the match
        day = match.group(1)
        month = match.group(2)
        year = match.group(3)

        return f"{day}.{month}.{year}"
    else:
        return None


def parse_links(html_doc: str):
    soup = BeautifulSoup(html_doc, "html.parser")

    # Find all sections with the class "section-box"
    sections = soup.find("section", class_="section-box")
    links = sections.find_all("a")

    elems = []
    for link in links:
        href = link.get("href")
        try:
            text = link.find("span", class_="tabs__item-text-decor").text
            print(text, href)
            elems.append({"name": text, "url": href})
        except Exception as e:
            print("Not found")

    return elems


async def main():
    client = httpx.AsyncClient()
    # downloads documents from links

    # html_page = await get_page(client, BASE_URL)
    with open(FILE_NAME, "r", encoding="utf-8") as file:
        html_page = file.read()

    soup = BeautifulSoup(html_page, "html.parser")

    links = soup.find_all(
        "a", text=lambda text: text and text.startswith("Бухгалтерская")
    )

    async def download(client: httpx.AsyncClient, url: str, name: str) -> dict:
        return {
            "data": (await client.get(url)).content,
            "name": "бухалтерская" + extract_date_from_string(name),
        }

    tasks = [download(client, l.get("href"), l.text) for l in links]
    results = await asyncio.gather(*tasks)
    print(len(results))
    for f in results:
        with open(f'{f["name"]}.pdf', "wb") as file:
            file.write(f["data"])


async def download_docs():
    client = httpx.AsyncClient()
    with open(FILE_NAME, "r", encoding="utf-8") as file:
        html_page = file.read()
    soup = BeautifulSoup(html_page, "html.parser")
    urls = soup.find("section", class_="section-box").find_all("a")
    tasks = []

    def get_filename(url: str) -> str:
        return url.split("/")[-1]

    async def download(url: str, filename: str):
        try:
            resp = await client.get(url)
            if resp.status_code != 200:
                print(resp.status_code)
        except Exception as e:
            print(f"retry for {filename} {url}")
            await asyncio.sleep(0.3)
            await download(url, filename)
            return
        print(f"saved {filename} {url}")
        with open(DATA_PATH / filename, "wb") as file:
            file.write(resp.content)

    skipped = 0
    for doc in tqdm(urls):
        url = doc.get("href")
        filename = get_filename(url)
        if os.path.exists(DATA_PATH / filename):
            skipped += 1
            continue
        if url.endswith(".pdf"):
            tasks.append(
                download(
                    url,
                    filename,
                )
            )
    print(f"skipped: {skipped} documents")

    await asyncio.gather(*tasks)


if __name__ == "__main__":
    asyncio.run(download_docs())
