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
	var charString string = ""
	m := make(map[string]int)
	var sum int = 0

	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	m["d"] = 4
	m["e"] = 5
	m["f"] = 6
	m["g"] = 7
	m["h"] = 8
	m["i"] = 9
	m["j"] = 10
	m["k"] = 11
	m["l"] = 12
	m["m"] = 13
	m["n"] = 14
	m["o"] = 15
	m["p"] = 16
	m["q"] = 17
	m["r"] = 18
	m["s"] = 19
	m["t"] = 20
	m["u"] = 21
	m["v"] = 22
	m["w"] = 23
	m["x"] = 24
	m["y"] = 25
	m["z"] = 26

	m["A"] = 27
	m["B"] = 28
	m["C"] = 29
	m["D"] = 30
	m["E"] = 31
	m["F"] = 32
	m["G"] = 33
	m["H"] = 34
	m["I"] = 35
	m["J"] = 36
	m["K"] = 37
	m["L"] = 38
	m["M"] = 39
	m["N"] = 40
	m["O"] = 41
	m["P"] = 42
	m["Q"] = 43
	m["R"] = 44
	m["S"] = 45
	m["T"] = 46
	m["U"] = 47
	m["V"] = 48
	m["W"] = 49
	m["X"] = 50
	m["Y"] = 51
	m["Z"] = 52


    lines, err := readLines("input.txt")

    if err != nil {
        fmt.Println("Error: %s\n", err)
        return
	}
	
    for _, line := range lines {
		half := len(line) / 2
		firstHalf := line[:half]
		secondHalf := line[half:]
		
		for i := 0; i < len(firstHalf); i++ {
			for j := 0; j < len(secondHalf); j++ {
				if(firstHalf[i] == secondHalf[j]) {
					charString = string(firstHalf[i])
				}
				
			}
		 }
			
		sum = sum + m[charString]
    
	}

	fmt.Println(sum)

}