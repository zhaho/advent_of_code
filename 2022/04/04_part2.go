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
	arr1 := []int{}
	arr2 := []int{}
	
	overlap := false
	overlapCounter := 0

    lines, err := readLines("input.txt")

    if err != nil {
        fmt.Println("Error: %s\n", err)
        return
	}
	
    for _, line := range lines {
		assignment := strings.Split(line,",")
		pair_1 := strings.Split(assignment[0],"-")
		pair_2 := strings.Split(assignment[1],"-")
		pair_1_start,_ := strconv.Atoi(pair_1[0])
		pair_1_end,_ := strconv.Atoi(pair_1[1])
		pair_2_start,_ := strconv.Atoi(pair_2[0])
		pair_2_end,_ := strconv.Atoi(pair_2[1])

		// Fill those arrays
		for i := pair_1_start; i < pair_1_end+1; i++ {
			arr1 = append(arr1,i)
		}
		for i := pair_2_start; i < pair_2_end+1; i++ {
			arr2 = append(arr2,i)
		}

		min1 := arr1[0]
		min2 := arr2[0]
		max1 := arr1[len(arr1)-1]
		max2 := arr2[len(arr2)-1]

		if !(max1 < min2 || max2 < min1) {
			overlap = true
		}
		
		if(overlap == true) {
			overlapCounter++ 
			overlap = false
		}
		
		arr1 = nil
		arr2 = nil

	}

	fmt.Println(overlapCounter,"overlaps")
}