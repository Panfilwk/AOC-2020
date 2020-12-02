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

func CheckValidity(entry string) bool {
  vals := strings.Split(entry, " ")
  quant_range, req_string, pass := vals[0], vals[1], vals[2]

  quants := strings.Split(quant_range, "-")
  min, err := strconv.Atoi(quants[0])
  check(err)
  max, err := strconv.Atoi(quants[1])
  check(err)

  var req_chara rune
  for _, rune := range req_string {
    req_chara = rune
    break
  }

  count := 0
  for _, chara := range pass {
    if chara == req_chara {
      count +=1
    }
    if count > max {
      return false
    }
  }
  return count >= min
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

  valid_count := 0
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    if CheckValidity(scanner.Text()) {
      valid_count++
    }
  }
  check(scanner.Err())
  fmt.Println(valid_count, "passwords are valid!")
}
