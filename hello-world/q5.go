package main

import (
  "fmt"
  "regexp"
  "strconv"
  "strings"
)

func main() {
  fmt.Printf("Hello, World!\n")
  const SIZE_LIMIT = 100000
  // store the file system paths in a map
  // filesys := make(map[string][]string)

  // var fileRegex string = "^\\d+ [a-zA-Z0-9]+\\.[a-z]+$"
  var dirRegex string = "^dir [a-zA-Z0-9]+$"

  var currentDir string

  var file_map = map[string][]string{}

  // reFile, err := regexp.Compile(fileRegex)
  // if err != nil {
  //   fmt.Printf("Error compiling regex: %v\n", err)
  // }

  reDir, err := regexp.Compile(dirRegex)
  if err != nil {
    fmt.Printf("Error compiling regex: %v\n", err)
  }
  
  lines, err := ReadFile("data.txt")
  if err != nil {
    fmt.Printf("The error is %v\n", err)
  }

  var dirSums = map[string]int{}

  var totalDirSum int
  for _, line := range lines {
    fmt.Printf("Line: %s\n", line)

    if strings.HasPrefix(line, "$") {
      var cmd string
      var args string
      
      tokens := strings.Split(line, " ")
      
      if validCmd(tokens[1]) {
        cmd = tokens[1]
        fmt.Printf("Command: %s\n", cmd)
      }

      if len(tokens) > 2 {
        args = tokens[2]
        fmt.Printf("Args: %s\n", args)
      }

      if cmd == "cd" && args != ".." {
        currentDir = args
        fmt.Printf("Current dir: %s\n", currentDir)
      }else if cmd == "cd" && args == ".." {
        // Go up one directory
        // assing currentDir to the parent directory

        fmt.Printf("Current dir: %s\n", currentDir)
        
      } else if cmd == "ls" {
        continue
      }
    } else if reDir.MatchString(line) {
      directoryName := strings.Split(line, " ")[1]
      if currentDir != "" {
        file_map[currentDir] = append(file_map[currentDir], directoryName)
      }
    } else {
      tokens := strings.Split(line, " ")
      size := tokens[0]
      name := tokens[1]

      sizeInt, err := strconv.Atoi(size)
      if err != nil {
        fmt.Printf("Error converting size to int: %v\n", err)
      }
      
      if currentDir != "" {
        file_map[currentDir] = append(file_map[currentDir], name)
        dirSums[currentDir] += sizeInt
      }
    }

    for _ , size := range dirSums {
      if size <= SIZE_LIMIT {
        totalDirSum += size
      }
    }

    fmt.Printf("File map: %v\n", file_map)
    fmt.Printf("Total dir sum: %d\n", totalDirSum)
  }
}

// Issues:
// 1. file map is only adding files and not the directories. Should add both.
// 2. Need to implement the logic for going up one directory

func validCmd(cmd string) bool {
  return cmd == "cd" || cmd == "ls"
}
// func cd(args ...string) {
//   fmt.Printf("cd\n")
// }

// func ls(args ...string) {
//   fmt.Printf("ls\n")
// }
