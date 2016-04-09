package main

import "fmt"

func main() {
 fmt.Print("Ingrese un texto: ")
 var input string
 fmt.Scanf("%s", &input)
 var p string
 for i := 1; i < (len(input)+1); i++ {
 	p=input[0:i]
 	fmt.Println(p)
 }

 for i := len(input)-1; i > 0; i-- {
 	p=input[0:i]
 	fmt.Println(p)
 }
}