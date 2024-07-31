package main

import (
  "bufio"
  "errors"
  "fmt"
  "os"
  "strconv"
  "strings"
  "time"
  // set "github.com/duke-git/lancet/v2/datastructure/set"
)

// advent of code 2023 Q2 part 1,2

func main()  {
    now := time.Now()
    lines, err := ReadFile("data.txt")
    if err != nil {
        fmt.Println("Error reading file: ", err)
        return
    }

    // var gameIdSum int
    var gamePoweredSum int

    for _, line := range lines {  
        game := strings.Split(line, ":")
        // gameId, err := strconv.Atoi(strings.Split(game[0], " ")[1])

        // fmt.Println(game)
        // fmt.Println(gameId)

        // if err != nil {
        //     fmt.Printf("string conversion issue: the string is not an int %v", err)
        // }

        subsets := strings.Split(game[1], ";")
        //  subset = 9 green, 10 red, 2 blue
        // var notPossible bool
        var red int
        var blue int
        var green int
        
        for _, subset := range subsets {
            cubes := strings.Split(strings.TrimSpace(subset), ",")
            for _, cube := range cubes {
                cubeDetails := strings.Split(strings.TrimSpace(cube), " ")

                // fmt.Println(cubeDetails)
                
                noOfCubes, err := strconv.Atoi(cubeDetails[0])
                if err != nil {
                    fmt.Printf("string conversion issue: the string is not an int %v", err)
                }
                colorOfCubes := cubeDetails[1]

                switch(colorOfCubes){
                    case "green":
                        if noOfCubes > green {
                            green = noOfCubes
                        }
                    case "blue":
                        if noOfCubes > blue {
                            blue = noOfCubes
                        }
                    case "red":
                        if noOfCubes > red {
                            red = noOfCubes
                        }
                }
                // if (colorOfCubes == "blue" && noOfCubes > 14) || (colorOfCubes == "red" && noOfCubes > 12) || (colorOfCubes == "green" && noOfCubes > 13) {
                //     notPossible = true
                // }
            }
        }

        gamePoweredSum += red * blue * green

        // if !notPossible {
        //     gameIdSum += gameId
        // }
    }

    // fmt.Printf("The game sum of possible games is %v\n", gameIdSum)
    fmt.Printf("The power of a set of cubes is equal to %v\n", gamePoweredSum)
    fmt.Println(time.Since(now))
}

func ReadFile(filePath string) ([]string, error) {
    var lines []string

    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening file: ", err)
        return nil, errors.New("error opening file")
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, nil
}
