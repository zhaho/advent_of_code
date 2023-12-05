import re

def scores_nums(line):
    matches = re.match(r"Card\s+(\d+):\s+(.*?)\s\|\s+(.*?)$", line)
    wins = set(int(n) for n in matches.group(2).strip().split())
    numbers = set(int(n) for n in matches.group(3).strip().split())
    return wins.intersection(numbers)
    
lines = open("input.txt").read().strip().splitlines()
cards  = [1] * len(lines)
for number, line in enumerate(lines):
    matches = scores_nums(line)
    for i in range(len(matches)):
        cards [number + i + 1] += cards [number]

print(sum(cards))