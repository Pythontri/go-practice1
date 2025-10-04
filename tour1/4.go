package main

import "fmt"

func add(x int, y int) int {//пишем тип после аргумента
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
