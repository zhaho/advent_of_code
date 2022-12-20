
def valid(min,max,c,pw):
    if pw.count(c) >= min and pw.count(c) <= max:
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