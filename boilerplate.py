import requests, browsercookie, pathlib

def get_input_lines(day):
    text = requests.get(f"https://adventofcode.com/2020/day/{day}/input", cookies=browsercookie.firefox()).text
    pathlib.Path("i").write_text(text)
    return text
