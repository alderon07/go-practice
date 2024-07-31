package main

import (
  "fmt"
  "regexp"
)

func main() {
  fmt.Printf("Hello, World!\n")

  // store the file system paths in a map
  // filesys := make(map[string][]string)

  var fileRegex string = "^\\d+ [a-zA-Z0-9]+\\.[a-z]+$"
  var dirRegex string = "^dir [a-zA-Z0-9]+$"

  reFile, err := regexp.Compile(fileRegex)
  if err != nil {
    fmt.Printf("Error compiling regex: %v\n", err)
  }

  reDir, err := regexp.Compile(dirRegex)
  if err != nil {
    fmt.Printf("Error compiling regex: %v\n", err)
  }

  fmt.Printf("Matched: %v\n", reFile.MatchString("345 1.txt"))

  fmt.Printf("Matched: %v\n", reDir.MatchString("dir test546hgf"))
}

func cd(args ...string) {
  fmt.Printf("cd\n")
}

func ls(args ...string) {
  fmt.Printf("ls\n")
}
