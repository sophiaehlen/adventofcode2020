package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    lines, err := readInput("input.txt")
    if err != nil {
        panic(err)
    }
    var validCount int
    for _, line := range lines {
        p := process(line)
        if p.valid() {
            validCount = validCount + 1
        }
    }
    fmt.Println(validCount)

}

type policy struct {
    min      int
    max      int
    letter   string
    password string
}

func (p *policy) valid() bool {
    count := strings.Count(p.password, p.letter)
    if count >= p.min && count <= p.max {
        return true
    }
    return false
}

func process(line string) *policy {
    c := strings.Split(line, " ")

    letterRange := strings.Split(c[0], "-")

    min, _ := strconv.Atoi(letterRange[0])
    max, _ := strconv.Atoi(letterRange[1])

    letter := string(c[1])
    letter = string(letter[0])
    password := c[2]
    return &policy{
        min:      min,
        max:      max,
        letter:   letter,
        password: password,
    }

}
