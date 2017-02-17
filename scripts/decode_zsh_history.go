package main

import (
  "bufio"
  "strings"
  "fmt"
  "os"
)

func main() {
  var fp *os.File
  var err error
  //var char byte

  if len(os.Args) < 2 {
    fp  = os.Stdin
  } else {
    fp, err = os.Open(os.Args[1])
    if err != nil {
      panic(err)
    }
    defer fp.Close()
  }

  scanner := bufio.NewScanner(fp)

  for scanner.Scan() {
    chars := []byte{}
    meta := false
    line := scanner.Text()
    x83 := strings.IndexByte(line, 131)

    if x83 > -1 {
      chars = []byte(line[:x83])
    } else {
      fmt.Println(string(line))
      continue
    }

    lineBytes := []byte(line)
    lineSize := len(lineBytes)

    for i := x83; i < lineSize; i++ {
      char := lineBytes[i]
      if char == 131 { // 0x83
        meta = true
      } else if (meta) {
        chars = append(chars, char ^ 32) // 0x20
        meta = false
      } else {
        chars = append(chars, char)
      }
    }

    fmt.Println(string(chars))
  }

  if err := scanner.Err(); err != nil {
    panic(err)
  }
}
