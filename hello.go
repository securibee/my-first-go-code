package main

import (
  "fmt"
  "os"
)

func Hello(name string) string {
  if name == "" {
    name = "World"
  } else if name == "os" {
    name, _ = os.Hostname()
  }
  return "Hello " + name
}

func main() {
  argument := ""
  if (len(os.Args) > 1) {
    argument = os.Args[1]
  }
  fmt.Println(Hello(argument))
}
