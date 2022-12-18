package main

import (
    "os"
    "bufio"
    "bytes"
    "io"
	"fmt"
	"strings"
	"strconv"
)

func readLines(path string) (lines []string, err error) {
    var (
        file *os.File
        part []byte
        prefix bool
    )
    if file, err = os.Open(path); err != nil {
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    buffer := bytes.NewBuffer(make([]byte, 0))
    for {
        if part, prefix, err = reader.ReadLine(); err != nil {
            break
        }
        buffer.Write(part)
        if !prefix {
            lines = append(lines, buffer.String())
            buffer.Reset()
        }
    }
    if err == io.EOF {
        err = nil
    }
    return
}

func reverseInts(input []string) []string {
    if len(input) == 0 {
        return input
    }
    return append(reverseInts(input[1:]), input[0]) 
}

func main() {
	temp := []string{}
	ship1 := []string{"S", "T","H","F","W","R"}
	ship2 := []string{"S","G","D","Q","W"}
	ship3 := []string{"B","T","W"}
	ship4 := []string{"D","R","W","T","N","Q","Z","J"}
	ship5 := []string{"F","B","H","G","L",",V","T","Z"}
	ship6 := []string{"L","P","T","C","V","B","S","G"}
	ship7 := []string{"Z","B","R","T","W","G","P"}
	ship8 := []string{"N","G","M","T","C","J","R"}
	ship9 := []string{"L","G","B","W"}

    lines, err := readLines("input.txt")

    if err != nil {
        fmt.Println("Error: %s\n", err)
        return
	}
	
    for i, line := range lines {
		if(i > 9) {
			a := strings.Split(line,"move ")
			a1 := strings.Split(a[1]," from")
			f := strings.Split(a1[1]," to")
			f1 := strings.Split(f[0]," ")
			t := strings.Split(f[1]," ")
			amount,_ := strconv.Atoi(a1[0])
			from,_ := strconv.Atoi(f1[1])
			to,_ := strconv.Atoi(t[1])

			for i := 0; i < amount; i++ {
				switch from {
				case 1:
					value := ship1[len(ship1)-1]
					temp = append(temp,value)
					ship1 = ship1[:len(ship1)-1]
					
				case 2:
					value := ship2[len(ship2)-1]
					temp = append(temp,value)
					ship2 = ship2[:len(ship2)-1]
				case 3:
					value := ship3[len(ship3)-1]
					temp = append(temp,value)
					ship3 = ship3[:len(ship3)-1]
				case 4:
					value := ship4[len(ship4)-1]
					temp = append(temp,value)
					ship4 = ship4[:len(ship4)-1]
				case 5:
					value := ship5[len(ship5)-1]
					temp = append(temp,value)
					ship5 = ship5[:len(ship5)-1]
				case 6:
					value := ship6[len(ship6)-1]
					temp = append(temp,value)
					ship6 = ship6[:len(ship6)-1]
				case 7:
					value := ship7[len(ship7)-1]
					temp = append(temp,value)
					ship7 = ship7[:len(ship7)-1]
				case 8:
					value := ship8[len(ship8)-1]
					temp = append(temp,value)
					ship8 = ship8[:len(ship8)-1]
				case 9:
					value := ship9[len(ship9)-1]
					temp = append(temp,value)
					ship9 = ship9[:len(ship9)-1]
				}

			}

			temp = reverseInts(temp)
			
			for _, tmp := range temp {

				switch to {
				case 1:
					tmp_string := string(tmp)
					ship1 = append(ship1,tmp_string)
				case 2:
					tmp_string := string(tmp)
					ship2 = append(ship2,tmp_string)
				case 3:
					tmp_string := string(tmp)
					ship3 = append(ship3,tmp_string)
				case 4:
					tmp_string := string(tmp)
					ship4 = append(ship4,tmp_string)
				case 5:
					tmp_string := string(tmp)
					ship5 = append(ship5,tmp_string)
				case 6:
					tmp_string := string(tmp)
					ship6 = append(ship6,tmp_string)
				case 7:
					tmp_string := string(tmp)
					ship7 = append(ship7,tmp_string)
				case 8:
					tmp_string := string(tmp)
					ship8 = append(ship8,tmp_string)
				case 9:
					tmp_string := string(tmp)
					ship9 = append(ship9,tmp_string)
				}

			}

			temp = nil
		}

	}


	output := ship1[len(ship1)-1]+ship2[len(ship2)-1]+ship3[len(ship3)-1]+ship4[len(ship4)-1]+ship5[len(ship5)-1]+ship6[len(ship6)-1]+ship7[len(ship7)-1]+ship8[len(ship8)-1]+ship9[len(ship9)-1]
	fmt.Println(output)
}