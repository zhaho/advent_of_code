maxRed=12
maxGreen=13
maxBlue=14
sumPossible=0
with open("input.txt") as data:
    for line in data:
        line = line.rstrip("\n")
        possible = True
        parts = line.split(":")
        game = int(parts[0].replace("Game ",""))
        shows = parts[1].split(";")
        for show in shows:
            dices = show.split(",")
            for dice in dices:
                amount = int(dice.split(" ")[1])
                color = dice.split(" ")[2]
                if color == "red" and amount > maxRed:
                    possible = False
                elif color == "blue" and amount > maxBlue:
                    possible = False
                elif color == "green" and amount > maxGreen:
                    possible = False
                
        if possible == True:
            sumPossible += game
    possible = False
    
print(sumPossible)
        