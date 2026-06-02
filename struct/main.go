package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}
