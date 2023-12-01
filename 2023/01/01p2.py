sum = 0
temp_char = ""
num_dict = {
  "one": "1",
  "two": "2",
  "three": "3",
  "four": "4",
  "five": "5",
  "six": "6",
  "seven": "7",
  "eight": "8",
  "nine": "9",
}
with open("input.txt", "r") as input_file:
    for r in input_file:
        num_first = None
        r = r.rstrip('\n')
        for i in range(0,len(r)):
            char = r[i]
            if char.isdigit():
                if num_first is None:
                    num_first = char
                num_second = char
            else:
                for j in range(i,len(r)):
                    temp_char += r[j]
                    if temp_char in num_dict:
                        if num_first is None:
                            num_first = num_dict[temp_char]
                        else:
                            num_second = num_dict[temp_char]
                temp_char = ""
        score = int(num_first+num_second)
        sum += int(score)
print(sum)