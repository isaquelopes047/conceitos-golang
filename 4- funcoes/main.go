package main

import "fmt"

func somar(n1 int8, n2 int8) int8{
	return n1 + n2
}

func main() {
	var soma int8 = somar(10, 20)
	fmt.Println(soma)
}