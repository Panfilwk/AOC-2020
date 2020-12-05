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

func max(x, y int) int {
  if x >= y {
    return x
  }
  return y
}

func get_seat_id(pass string) int {
  min_row, max_row, min_col, max_col := 0, 127, 0, 7
  for _, dir := range pass {
    if dir == 'F' {
      max_row = (min_row + max_row) / 2
    } else if dir == 'B' {
      min_row = (min_row + max_row) / 2
    } else if dir == 'R' {
      min_col = (min_col + max_col) / 2
    } else if dir == 'L' {
      max_col = (min_col + max_col) / 2
    } else {
      panic("Malformed boarding pass!")
    }
  }
  return max_row * 8 + max_col
}

func get_max_seat_id() int {
  file, err := os.Open("input.txt")
  check(err)
  defer file.Close()

  max_seat_id := -1

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    max_seat_id = max(max_seat_id, get_seat_id(scanner.Text()))
  }
  check(scanner.Err())

  return max_seat_id
}

func main() {
  defer func() {
    if r := recover(); r != nil {
      log.Fatal(recover())
    }
  }()

  fmt.Println(get_max_seat_id())
}
