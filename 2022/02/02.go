package main

import (
    "os"
    "bufio"
    "bytes"
    "io"
    "fmt"
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

func main() {
	m := make(map[string]int)
	var sum int = 0
    lines, err := readLines("input.txt")

    if err != nil {
        fmt.Println("Error: %s\n", err)
        return
	}
	
	m["Rock"] = 1
	m["Paper"] = 2
	m["Scissor"] = 3

	m["Loss"] = 0
	m["Draw"] = 3
	m["Win"] = 6

    for _, line := range lines {
		var amount int = 0
		opponentsChoice := line[:1]
		myChoice := line[2:]

		if (opponentsChoice == "A") { // Rock
			if (myChoice == "X") {
				amount = m["Rock"] + m["Draw"]
			}else if (myChoice == "Y") {
				amount = m["Paper"] + m["Win"]
			}else if (myChoice == "Z") {
				amount = m["Scissor"] + m["Loss"]
			}	
		}

		if (opponentsChoice == "B") { // Paper
			if (myChoice == "X") {
				amount = m["Rock"] + m["Loss"]
			}else if (myChoice == "Y") {
				amount = m["Paper"] + m["Draw"]
			}else if (myChoice == "Z") {
				amount = m["Scissor"] + m["Win"]
			}	
		}

		if (opponentsChoice == "C") { // Scissor
			if (myChoice == "X") {
				amount = m["Rock"] + m["Win"]
			}else if (myChoice == "Y") {
				amount = m["Paper"] + m["Loss"]
			}else if (myChoice == "Z") {
				amount = m["Scissor"] + m["Draw"]
			}	
		}

		sum = sum + amount
		amount = 0
    
	}
	
	fmt.Println(sum)

}