sumPower=0
with open("input.txt") as data:
    for line in data:
        line = line.rstrip("\n")
        possible = True
        parts = line.split(":")
        game = int(parts[0].replace("Game ",""))
        shows = parts[1].split(";")
        minRed=1
        minGreen=1
        minBlue=1
        for show in shows:
            dices = show.split(",")
            for dice in dices:
                amount = int(dice.split(" ")[1])
                color = dice.split(" ")[2]
                if color == "red" and amount > minRed:
                    minRed = amount
                elif color == "blue" and amount > minBlue:
                    minBlue = amount
                elif color == "green" and amount > minGreen:
                    minGreen = amount
    
        power = minRed * minBlue * minGreen
        sumPower += power
    
print(sumPower)
        