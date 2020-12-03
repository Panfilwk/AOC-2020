package main

import (
  "bufio"
  "log"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func ParseLine(entry string) (int, int, rune, string) {
  vals := strings.Split(entry, " ")
  quant_range, req_string, pass := vals[0], vals[1], vals[2]

  quants := strings.Split(quant_range, "-")
  first, err := strconv.Atoi(quants[0])
  check(err)
  second, err := strconv.Atoi(quants[1])
  check(err)

  var req_rune rune
  for _, rune := range req_string {
    req_rune = rune
    break
  }

  return first, second, req_rune, pass
}

func CheckValidity1(entry string) bool {
  min, max, req_rune, pass := ParseLine(entry)

  count := 0
  for _, cur_rune := range pass {
    if cur_rune == req_rune {
      count +=1
    }
    if count > max {
      return false
    }
  }
  return count >= min
}

func CheckValidity2(entry string) bool {
  first, second, req_rune, pass := ParseLine(entry)

  valid := false
  for i, cur_rune := range pass {
     valid = valid != ((i == first-1 || i == second-1) && cur_rune == req_rune)
  }
  return valid
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

  valid_count_1, valid_count_2 := 0, 0
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    if CheckValidity1(scanner.Text()) {
      valid_count_1++
    }
    if CheckValidity2(scanner.Text()) {
      valid_count_2++
    }
  }
  check(scanner.Err())
  fmt.Println(valid_count_1, "passwords are valid by the old metric!")
  fmt.Println(valid_count_2, "passwords are valid!")
}
