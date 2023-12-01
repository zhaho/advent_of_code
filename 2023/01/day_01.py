first_number = None
second_number = None
sum = 0
with open("input.txt", "r") as input_file:
    for row in input_file:
        row = row.rstrip('\n')
        for i in range(0,len(row)):
            current_char = row[i]
            if current_char.isdigit():
                if first_number is None:
                    first_number = current_char
                second_number = current_char
        joined_numbers = first_number + second_number
        sum = sum + int(joined_numbers)
        first_number = None
print(sum)