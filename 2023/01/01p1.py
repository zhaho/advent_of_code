sum = 0
with open("input.txt", "r") as file:
    for r in file:
        num_first = None
        r = r.rstrip('\n')
        for i in range(0,len(r)):
            char = r[i]
            if char.isdigit():
                if num_first is None:
                    num_first = char
                num_second = char
        score = int(num_first+num_second)
        sum += score
print(sum)