package main

import "fmt"

func main() {
  fmt.Print("Ingrese un texto: ")
  var input string
  fmt.Scanln("%s", &input)

  fmt.Println(input)
}