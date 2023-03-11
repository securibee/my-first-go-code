package main

import (
  "fmt"
  "os"
)

func Hello(name string) string {
  if name == "" {
    name = "World"
  }
  return "Hello " + name
}

//str, _ := os.Hostname()

func main() {
  argument := ""
  if (len(os.Args) > 1) {
    argument = os.Args[1]
  }
  fmt.Println(Hello(argument))
}
