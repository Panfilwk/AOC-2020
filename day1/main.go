package main

import (
  "bufio"
  "log"
  "fmt"
  "strconv"
  "sort"
  "os"
)

const TARGET int = 2020

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func ReadData() []int {
  file, err := os.Open("input.txt")
  check(err)
  defer file.Close()

  var data []int
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
      i, err := strconv.Atoi(scanner.Text())
      check(err)
      data = append(data, i)
  }
  check(scanner.Err())
  return data
}

func SearchPair(data []int) int {
  for lin := 0; lin < len(data); lin ++ {
    var min, max, oldbin int = 0, len(data) - 1, -1

    for bin := (min+max)/2; oldbin != bin; bin = (min+max)/2 {
      if data[lin] + data[bin] == TARGET {
        return data[lin] * data[bin]
      } else if data[lin] + data[bin] < TARGET {
        min = bin
      } else {
        max = bin
      }

      oldbin = bin
    }
  }
  panic("No mathches for target")
}

func SearchTriple(data []int) int {
  for lin1 := 0; lin1 < len(data); lin1 ++ {
    for lin2 := 0; lin2 < len(data); lin2 ++ {
      var min, max, oldbin int = 0, len(data) - 1, -1

      for bin := (min+max)/2; oldbin != bin; bin = (min+max)/2 {
        if data[lin1] + data[lin2] + data[bin] == TARGET {
          return data[lin1] * data[lin2] * data[bin]
        } else if data[lin1] + data[lin2] + data[bin] < TARGET {
          min = bin
        } else {
          max = bin
        }

        oldbin = bin
      }
    }
  }
  panic("No mathches for target")
}

func main() {
  defer func() {
    if r := recover(); r != nil {
      log.Fatal(recover())
    }
  }()

  data := ReadData()
  sort.Ints(data)
  fmt.Println(SearchPair(data))
  fmt.Println(SearchTriple(data))
}