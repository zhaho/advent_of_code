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
			
			if(len(m) == 4) {
				fmt.Println(i+4)
				break
			}

			for k := range m {
				delete(m, k)
			}
		}
	}


	
}