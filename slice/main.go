package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5}

	fmt.Println("slice len is: ", len(slice))
	slice = append(slice, 6)
	fmt.Println("slice cap is: ", cap(slice))
	fmt.Println("slice is: ", slice)
}
