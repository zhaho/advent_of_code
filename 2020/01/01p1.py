import sys
lines = open('input.txt').read().split('\n')
for line in lines:
    for i in range(len(lines)): 
        if (int(line)+int(lines[i])) == 2020:
            print(int(line)*int(lines[i]))
            sys.exit()