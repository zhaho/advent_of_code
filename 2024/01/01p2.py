left = []
right = []
similarity = 0

# Separate the input
with open("input.txt", "r") as file:
    for r in file:
        left.append(int(r.split("   ")[0]))
        right.append(int(r.split("   ")[1].split("\n")[0]))

# Count Similarity
for i, num in enumerate(left):
    similarity += left[i] * right.count(left[i])
    
print(similarity)
    
