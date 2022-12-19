pos, depth = 0,0
for f in open('input.txt','r'):
    action = f.split(" ")[0]
    amount = int(f.split(" ")[1])
    if action == 'forward':
        pos += amount
    elif action == "down":
        depth += amount
    else:
        depth = depth - amount
        
print(pos*depth)