package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world") //может принимать 2 стринга и менять их
	fmt.Println(a, b)
}