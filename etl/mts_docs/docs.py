import httpx
import asyncio
from bs4 import BeautifulSoup
import re

BASE_URL = "https://moskva.mts.ru/about/investoram-i-akcioneram/korporativnoe-upravlenie/raskritie-informacii/godovaya-otchetnost"
FILE_NAME = "page.html"


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


if __name__ == "__main__":
    asyncio.run(main())
