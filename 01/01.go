package main

import (
    "os"
    "bufio"
    "bytes"
    "io"
    "fmt"
    "strings"
    "strconv"
    "sort"
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

func writeLines(lines []string, path string) (err error) {
    var (
        file *os.File
    )

    if file, err = os.Create(path); err != nil {
        return
    }
    defer file.Close()

    for _,item := range lines {
        _, err := file.WriteString(strings.TrimSpace(item) + "\n"); 
        if err != nil {
            fmt.Println(err)
            break
        }
    }
    return
}

func main() {
    sum := 0
    lines, err := readLines("input.txt")
    arr := []int{}

    if err != nil {
        fmt.Println("Error: %s\n", err)
        return
    }

    for _, line := range lines {
        
        chars := len(line)
        if(chars > 0) {
            amount , err := strconv.Atoi(line)
            sum = amount + sum
            if err != nil{
                fmt.Println(sum)
            }
        }else{
            arr = append(arr, sum)
            sum = 0
        }
    
    }
    arrSlice := arr[:]
    sort.Sort(sort.Reverse(sort.IntSlice(arrSlice)))

    fmt.Println("Elf #1",arr[0], "calories!")
    fmt.Println("Elf #2",arr[1], "calories!")
    fmt.Println("Elf #3",arr[2], "calories!")
    fmt.Println("\nTotal",arr[0]+arr[1]+arr[2], "calories!")
    err = writeLines(lines, "output.txt")
    fmt.Println(err)
}