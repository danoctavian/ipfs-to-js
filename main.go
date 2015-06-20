package main

import (
  "fmt"
  "github.com/gopherjs/gopherjs/js"
)

func add(x, y int) int {
  return x + y
}

func main() {
  fmt.Println("running a main")

  js.Module.Get("exports").Set("ops", map[string]interface{}{
        "add": add,
  })
}

