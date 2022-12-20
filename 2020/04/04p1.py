with open("input.txt") as f:
    s = f.read().strip().split("\n\n")

def valid(arr):
    requiredTypes = ['ecl','iyr','eyr','hgt','hcl','ecl','pid','cid']
    temp = arr.split(' ')
    types = 0
    for f in temp:
        
        currentType = f.split(':')[0]        
        if currentType in requiredTypes:
            types += 1
        
    print(types)
validPassports = 0

for r in s:
    if (valid(r)):
        validPassports += 1

print(validPassports)