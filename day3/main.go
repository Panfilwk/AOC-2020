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

func tree_count(right, down int) int {
  file, err := os.Open("input.txt")
  check(err)
  defer file.Close()

  width, idx, count := -1, 0, 0

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    if width == -1 {
      width = len(scanner.Text())
    }
    if idx % down == 0 && []rune(scanner.Text())[right*idx/down % width] == '#' {
      count++
    }
    idx++
  }
  check(scanner.Err())
  return count
}

func main() {
  defer func() {
    if r := recover(); r != nil {
      log.Fatal(recover())
    }
  }()

  count11 := tree_count(1, 1)
  count31 := tree_count(3, 1)
  count51 := tree_count(5, 1)
  count71 := tree_count(7, 1)
  count12 := tree_count(1, 2)
  fmt.Println("On the \"Right 3, Down 1\" path, you hit", count31, "trees")
  fmt.Println("Product of all paths", count11*count31*count51*count71*count12)
}
