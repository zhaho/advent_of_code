# Read file
first_number = None
second_number = None
sum = 0
temp_char = ""
text_numbers = {
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
    for row in input_file:
        row = row.rstrip('\n')
        for i in range(0,len(row)):
            current_char = row[i]
            if current_char.isdigit():
                if first_number is None:
                    first_number = current_char
                second_number = current_char
            else:
                for char_i in range(i,len(row)):
                    temp_char = temp_char + row[char_i]
                    if temp_char in text_numbers:
                        if first_number is None:
                            first_number = text_numbers[temp_char]
                        else:
                            second_number = text_numbers[temp_char]
                temp_char = ""
        joined_numbers = first_number + second_number
        sum += int(joined_numbers)
        first_number = None

print(sum)