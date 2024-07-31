package main

import (
  "fmt"
  "strconv"
  "math"
  "strings"
)

// import readFile function from q1.go
// import readFile function from q1.go

func q4() {
  var filepath string = "data.txt"
  fmt.Printf("Reading file %s\n", filepath)
  
  lines, err := ReadFile(filepath)

  if(err != nil) {
    fmt.Printf("The error is %v\n", err)
  }

  var sum int = 0
  for _ , line := range lines {
    tokens := strings.Split(line, "|")

    fmt.Printf("\nTokens: %s\n", tokens)
    winners := strings.Fields(strings.TrimSpace(strings.Split(tokens[0], ":")[1]))
    myNumbers := strings.Fields(strings.TrimSpace(tokens[1]))

    fmt.Printf("Winners: %s\n", winners)
    fmt.Printf("My Numbers: %s\n", myNumbers)

    fmt.Printf("len(winners): %d, len(myNumbers): %d\n", len(winners), len(myNumbers))
    sum += int(math.Exp2(float64(checkWinners(winners, myNumbers) - 1)))

    fmt.Printf("Winners: %s\n", winners)
    fmt.Printf("My Numbers: %s\n", myNumbers)
  }

  fmt.Printf("The sum of all the powered games is %d\n", sum)

}

func checkWinners(winners []string, myNumbers []string) int {
  var count int = 0
  
  for _, winner := range winners {
    for _, myNumber := range myNumbers {
      // fmt.Printf("Winner: %s, My Number: %s\n", winner, myNumber)

      winnerInt, err := strconv.Atoi(winner)
      if err != nil {
          fmt.Printf("string conversion issue: the string is not an int %v\n", err)
      }
      
      myNumberInt, err := strconv.Atoi(myNumber)
      if err != nil {
          fmt.Printf("string conversion issue: the string is not an int %v\n", err)
      }
      
      if winnerInt == myNumberInt {
          count++
      }
    }
  }

  fmt.Printf("The count is %d\n", count)

  return count
}
