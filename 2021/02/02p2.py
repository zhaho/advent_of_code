pos, aim, depth = 0,0,0
for f in open('input.txt','r'):
    action = f.split(" ")[0]
    amount = int(f.split(" ")[1])
    if action == 'forward':
        pos += amount
        depth += (aim * amount)
    elif action == "down":
        aim += amount
    else:
        aim = aim - amount
        
print(pos*depth)