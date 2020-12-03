package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    input, err := readInput("input.txt")
    if err != nil {
        panic(err)
    }

    for _, i := range input {
        for j := 1; j < len(input)-1; j++ {
            for k := 2; k < len(input)-2; k++ {
                if i+input[j]+input[k] == 2020 {
                    fmt.Println(i * input[j] * input[k])
                    return
                }
            }
        }
    }
}

func readInput(filename string) ([]int, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)

    var nums []int

    for scanner.Scan() {
        text, err := strconv.Atoi(scanner.Text())
        if err != nil {
            return nil, err
        }
        nums = append(nums, text)
    }

    return nums, nil
}
