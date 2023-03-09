package main

import "fmt"

func main() {
	for i := 0; i <= 4; i++ {
		fmt.Printf("Nilai i = %v\n", i)
	}

	// var characters = []rune{'С', 'А', 'Ш', 'А', 'Р', 'В', 'О'}
	var characters = "САШАРВО"
	for j := 0; j <= 10; j++ {
		if j == 5 {
			for index, value := range characters {
				fmt.Printf("character %#U starts at byte position %d\n", value, index)
			}
			continue
		}

		fmt.Printf("Nilai j = %v\n", j)
	}
}
