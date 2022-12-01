import requests, browsercookie, pathlib

def get_input_lines(year, day):
    return requests.get(f"https://adventofcode.com/{year}/day/{day}/input", cookies=browsercookie.firefox()).text
