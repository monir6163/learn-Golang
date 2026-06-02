package main

import "fmt"

func main() {
	a := 500
	p := &a
	fmt.Println(a) // 500
	fmt.Println(p) // 0xc000016098

	b := &a
	fmt.Println(b)  // 0xc000016098
	fmt.Println(*b) // 500

	*p = 1000
	fmt.Println(a) // 1000
}
