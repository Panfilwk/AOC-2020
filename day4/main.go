package main

import (
  "bufio"
  "log"
  "fmt"
  "os"
  "strings"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func read_data() []string {
  file, err := os.Open("input.txt")
  check(err)
  defer file.Close()

  entry := ""
  var data []string

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    if (line == "") {
      data = append(data, entry)
      entry = ""
    }
    entry = entry + line + " "
  }
  data = append(data, entry)
  return data
}

func verify_passport (entry string) bool{
validator := map[string]bool{
    "byr": false,
    "iyr": false,
    "eyr": false,
    "hgt": false,
    "hcl": false,
    "ecl": false,
    "pid": false,
  }

  fields := strings.Split(entry, " ")
  for _, field := range fields {
    validator[strings.Split(field, ":")[0]] = true
  }
  for _, valid := range validator {
    if !valid {
      return false
    }
  }
  return true
}

func main() {
  defer func() {
    if r := recover(); r != nil {
      log.Fatal(recover())
    }
  }()

  data := read_data()
  valid_passports := 0
  for _, entry := range data {
    if verify_passport(entry) {
      valid_passports++
    }
  }
  fmt.Println("There are", valid_passports, "valid_passports")
}
