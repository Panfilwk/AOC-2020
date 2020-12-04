package main

import (
  "bufio"
  "log"
  "fmt"
  "os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  defer func() {
    if r := recover(); r != nil {
      log.Fatal(recover())
    }
  }()

  file, err := os.Open("input.txt")
  check(err)
  defer file.Close()

  width, idx, count := -1, 0, 0

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    if width == -1 {
      width = len(scanner.Text())
    }
    if []rune(scanner.Text())[3*idx % width] == '#' {
      fmt.Println("Hit tree on row", idx)
      count++
    }
    idx++
  }
  check(scanner.Err())
  fmt.Println(count)
}
