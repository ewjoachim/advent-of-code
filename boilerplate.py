import requests, browsercookie, pathlib

def get_input_lines(year, day):
    text = requests.get(f"https://adventofcode.com/{year}/day/{day}/input", cookies=browsercookie.firefox()).text
    pathlib.Path("i").write_text(text)
    return text
