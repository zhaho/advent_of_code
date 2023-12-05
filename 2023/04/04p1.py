import re

def scores_nums(line):
    matches = re.match(r"Card\s+(\d+):\s+(.*?)\s\|\s+(.*?)$", line)
    wins = set(int(n) for n in matches.group(2).strip().split())
    numbers = set(int(n) for n in matches.group(3).strip().split())
    return wins.intersection(numbers)
    
lines = open("input.txt").read().strip().splitlines()
sum = 0
for line in lines:
    matches = scores_nums(line)
    sum += int(2 ** (len(matches) - 1))

print(sum)