package main

import (
	"bufio"
	"fmt"
	"os"
	// "rsc.io/quote"
    "strconv"
    "strings"
)

func main() {
    // fmt.Println(quote.Go())
    // fmt.Println(quote.Hello())
    // fmt.Println(quote.Glass())
    // fmt.Println(quote.Opt())

    lines, err := readFile("data.txt")

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(lines)

    var calibrationValues []int

    for _, char := range lines {
        numbers := findNumbers(char)

        fmt.Println(numbers)

        var str_slice []string
        for _, number := range numbers {
            str_slice = append(str_slice, strconv.Itoa(number))
        }

        resultString := strings.Join(str_slice, "")

        fmt.Println(resultString)

        if result, err := strconv.Atoi(resultString); err == nil {
            calibrationValues = append(calibrationValues, result)
        }

    }

    // sum the calibration values
    var sum int
    for _, value := range calibrationValues {
        sum += value
    }

    fmt.Println(sum)
}

func findNumbers(input string) []int {
    var numbers []int

    chars := []rune(input)
    
    for _, char:= range chars {
        if int(char) >= 48 && int(char) <= 57 {

            number, err := strconv.Atoi(string(char))
            if err != nil {
                fmt.Println(err)
            }
            numbers = append(numbers, number)
            break
        }
    }

    for i := len(chars) - 1; i >= 0; i-- {
        if int(chars[i]) >= 48 && int(chars[i]) <= 57 {

            number, err := strconv.Atoi(string(chars[i]))
            if err != nil {
                fmt.Println(err)
            }
            numbers = append(numbers, number)
            break
        }
    
    }

    return numbers
}

func readFile(filepath string) ([]string, error) {
    file, err := os.Open(filepath)
    
    if err != nil {
        return nil, err
    }

    // defer runs after the function has finished
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return lines, nil
}
