package main

import "fmt"

func main() {
	i := 21

	fmt.Printf("%v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%%\n")

	j := true

	fmt.Printf("%t\n", j)
	fmt.Printf("%c\n", 1071)
	fmt.Printf("%d\n", 21)
	fmt.Printf("%o\n", 21)
	fmt.Printf("%x\n", 15)
	fmt.Printf("%X\n", 15)
	fmt.Printf("%U\n", 1071)

	var k float64 = 123.456

	fmt.Printf("%f\n", k)
	fmt.Printf("%E\n", k)
}
