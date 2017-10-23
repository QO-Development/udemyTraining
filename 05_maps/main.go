package main

import (
	"fmt"
)

func main() {

	//colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#4bf745",
		"blue":  "#4ff123",
	}

	//Add a value
	colors["white"] = "#ffffff"

	//Delete a value
	delete(colors, "white")

	//fmt.Println(colors)
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
