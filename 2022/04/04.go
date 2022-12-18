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

func main() {
	count := 0
	pair_1_arr := []int{}
	pair_2_arr := []int{}
	foundItems := 0

    lines, err := readLines("input.txt")

    if err != nil {
        fmt.Println("Error: %s\n", err)
        return
	}
	
    for _, line := range lines {
		// Split assignments and pairs
		assignment := strings.Split(line,",")
		pair_1 := strings.Split(assignment[0],"-")
		pair_2 := strings.Split(assignment[1],"-")

		// Fill the first array
		pair_1_start,_ := strconv.Atoi(pair_1[0])
		pair_1_end,_ := strconv.Atoi(pair_1[1])
		for i := pair_1_start; i < pair_1_end+1; i++ {
			pair_1_arr = append(pair_1_arr,i)
		}

		// Fill the second array
		pair_2_start,_ := strconv.Atoi(pair_2[0])
		pair_2_end,_ := strconv.Atoi(pair_2[1])
		for i := pair_2_start; i < pair_2_end+1; i++ {
			pair_2_arr = append(pair_2_arr,i)
		}
		
		// Loop through the first array, and check within the second array if there is an identical value
		for _, item_1 := range pair_1_arr {
			for _,item_2 := range pair_2_arr {
				if item_1 == item_2 {
					foundItems = foundItems + 1
				}
			}
		}

		// Check if we fully covered these numbers in one of the arrays
		if (foundItems == len(pair_1_arr)) || (foundItems == len(pair_2_arr)) {
			count = count + 1
		}
		
		// Reset variables
		foundItems = 0
		pair_1_arr = nil
		pair_2_arr = nil

	}

	// Output the amount of amount of events
	fmt.Println(count)
}