row_nest_dict = {}
with open("input.txt","r") as data:
    for count, line in enumerate(data,0):
        temp_nr = ""
        line = line.rstrip("\n")
        
        for idx, ch in enumerate(line,0):
            row_nest_dict[(count,idx)] = ch
            
            
        for idx, ch in enumerate(line,0):
            if ch.isdigit():
                if count != 0 or idx != 0:
                    print("is legit:",count,idx,ch)
                    
                else:
                    print(count,idx,ch) 
            else:
                print(count,idx,ch) 