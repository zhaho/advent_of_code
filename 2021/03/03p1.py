from collections import Counter

lines = open('input.txt').read().split('\n')
gamma = ''
epsilon = ''
for i in range(len(lines[0])):
    c = Counter(l[i] for l in lines)
    gamma += c.most_common()[0][0]
    epsilon += c.most_common()[1][0]

gamma = int(gamma, 2)
epsilon = int(epsilon, 2)

print(gamma * epsilon)
        