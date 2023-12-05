sum = 0
with open("input.txt", "r") as lines:
    for line in lines:
        pts = 0
        winning_numbers = line.rstrip("\n").split("|")[0].split(":")[1].split()
        lottery_numbers = line.rstrip("\n").split("|")[1].split()
        first_won = True
        for nr in lottery_numbers:
            if nr in winning_numbers:
                if not first_won:
                   pts = pts * 2
                else:
                    pts = 1
                    first_won = False
        sum += pts
print(sum)