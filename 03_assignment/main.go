package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range s {
		if (value % 2) != 0 {
			//This number is odd
			fmt.Printf("%d is Odd\n", value)
		} else {
			//This number is even
			fmt.Printf("%d is Even\n", value)
		}

	}
}
