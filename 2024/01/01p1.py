# Ser variables
left = []
right = []
distance = 0

# Separate the input and sort it
with open("input.txt", "r") as file:
    for r in file:
        left.append(int(r.split("   ")[0]))
        right.append(int(r.split("   ")[1].split("\n")[0]))

s_left = sorted(left)        
s_right = sorted(right)

# Count distance
for i, num in enumerate(s_left):
    if s_left[i] > s_right[i]:
        distance += s_left[i] - s_right[i]
    elif s_left[i] < s_right[i]:
        distance += s_right[i] - s_left[i]
    else:
        pass
    
print(distance)
