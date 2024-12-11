safe = 0

def is_safe(row):
    row = list(map(int, row.split()))
    decrease = increase = 0
    
    for i in range(len(row) - 1):
        diff = abs(row[i] - row[i+1])
        if 0 < diff < 4:
            if row[i] > row[i+1]:
                decrease += 1
            elif row[i] < row[i+1]:
                increase += 1

    return decrease == len(row) - 1 or increase == len(row) - 1

with open("input.txt", "r") as file:
    safe = sum(1 for r in file if is_safe(r))

print(safe)
