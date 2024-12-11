# Variables
safe = 0

def is_safe(row):
    row = row.split(" ")
    decrease = 0
    increase = 0
    for i, num in enumerate(row):
        num = int(num)
        # Check if not the last row
        if i != len(row)-1:
            if num > int(row[i+1]):
                diff = num - int(row[i+1])
                if diff > 0 and diff < 4:
                    decrease += 1
            elif num < int(row[i+1]):
                diff = int(row[i+1]) - num
                if diff > 0 and diff < 4:
                    increase += 1
                    
    if decrease == len(row)-1 or increase == len(row)-1:
        decrease = 0
        increase = 0
        return True
    else:
        decrease = 0
        increase = 0
        return False
        
        

# Separate the input and sort it
with open("input.txt", "r") as file:
    for r in file:
        safe += 1 if is_safe(r) else 0

print(safe)

