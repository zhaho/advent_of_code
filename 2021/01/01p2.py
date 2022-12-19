arr, incSum, prevSum = [],0,0
lines = open('input.txt').read().split('\n')
for i in range(len(lines)-2):
    arr.extend((int(lines[i]),int(lines[i+1]),int(lines[i+2])))
    if(len(arr) == 3):
        print(sum(arr))
        if sum(arr) > prevSum:
            incSum += 1
        prevSum = sum(arr)
        arr = []
        
print(incSum-1)