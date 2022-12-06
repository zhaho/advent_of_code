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
    lines, err := readLines("input.txt")

    if err != nil {
        fmt.Println("Error: %s\n", err)
        return
	}
	
    for _, line := range lines {
		for i := 0; i < len(line); i++ {
			m := make(map[string]bool)
			m[string(line[i])] = true
			m[string(line[i+1])] = true
			m[string(line[i+2])] = true
			m[string(line[i+3])] = true
			m[string(line[i+4])] = true
			m[string(line[i+5])] = true
			m[string(line[i+6])] = true
			m[string(line[i+7])] = true
			m[string(line[i+8])] = true
			m[string(line[i+9])] = true
			m[string(line[i+10])] = true
			m[string(line[i+11])] = true
			m[string(line[i+12])] = true
			m[string(line[i+13])] = true
			fmt.Println(m)
			if(len(m) == 14) {
				fmt.Println(i+14)
				break
			}
			
			for k := range m {
				delete(m, k)
			}
		}
	}


	
}