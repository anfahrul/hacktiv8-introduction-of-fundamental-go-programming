package main

import "fmt"

func main() {
	var sentence = "selamat malam"

	mapOfChar := map[string]int{}

	for _, char := range sentence {
		fmt.Println(string(char))
		mapOfChar[string(char)] += 1
	}

	fmt.Println(mapOfChar)
}
