package main

import "fmt"

func add(x, y int) int {//можем не писать кроме последнего если хотим сделать у всех один тип
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
