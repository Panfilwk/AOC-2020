package main

import (
  "bufio"
  "log"
  "fmt"
  "os"
  "sort"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
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

func get_seat_ids() []int {
  file, err := os.Open("input.txt")
  check(err)
  defer file.Close()

  var seats []int

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    seats = append(seats, get_seat_id(scanner.Text()))
  }
  check(scanner.Err())

  return seats
}

func main() {
  defer func() {
    if r := recover(); r != nil {
      log.Fatal(recover())
    }
  }()

  seat_ids := get_seat_ids()
  sort.Ints(seat_ids)

  fmt.Println("Max seat ID is", seat_ids[len(seat_ids) - 1])

  first_id := -1
  for idx, seat_id := range seat_ids {
    if first_id == -1 {
      first_id = seat_id
    }
    if seat_id != first_id + idx {
      fmt.Println("Your seat ID is", first_id + idx)
      break
    }
  }
}
