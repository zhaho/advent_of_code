prevNum, incNum, decNum = 0,0,0
for f in open('input.txt','r'):
    line = int(f.strip())
    if(line > prevNum and prevNum != 0):
        incNum += 1
    elif (line < prevNum):
        decNum += 1 
    prevNum = line

print("Increased: "+str(incNum)+"Decreased: "+str(decNum))
