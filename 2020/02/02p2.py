
def valid(pos1,pos2,c,pw):
    count, found = 1,0
    for ch in pw:
        if (count == pos1 and c == ch) or (count == pos2 and c == ch):
            found += 1
        count += 1
    if found == 1:
        return True
    else:
        return False
                    
validPass = 0

lines = open('input.txt').read().split('\n')
for line in lines:
    min = int(line.split(' ')[0].split('-')[0])
    max = int(line.split(' ')[0].split('-')[1])
    c = line.split(' ')[1].split(':')[0]
    pw = line.split(' ')[2]
    
    if valid(min,max,c,pw):
        validPass += 1

print(validPass)
    
        
