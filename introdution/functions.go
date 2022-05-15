package main

import "fmt"

var num = 1

func main() {

	test(num)
}
func test(value int) {
	value += 1

	fmt.Println(value)
}
