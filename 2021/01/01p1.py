prevNum, incNum, decNum = 0,0,0
for f in open('input.txt','r'):
    line = int(f.strip())
    if(line > prevNum and prevNum != 0):
        incNum += 1
    prevNum = line

print(str(incNum))