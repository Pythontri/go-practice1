package main

import "fmt"

func split(sum int) (x, y int) {//если 2 аргумента то скобки нужны будут, они именные поэтому после ретерн можно не писать
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
