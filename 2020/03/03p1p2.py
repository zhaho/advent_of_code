with open("input.txt") as f:
    s = f.read().strip().split("\n")

s = [list(x) for x in s]

def encounters(xd,yd):
    x,y = 0,0
    c = 0
    while y < len(s):
        if s[y][x%len(s[y])] == "#":
            c += 1
        x += xd
        y += yd
    return c

print(encounters(3,1))
print(encounters(1,1)*encounters(3,1)*encounters(5,1)*encounters(7,1)*encounters(1,2))
