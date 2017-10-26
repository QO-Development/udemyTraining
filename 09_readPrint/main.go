package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	args := os.Args
	index := len(args) - 1
	fileToOpen := args[index]

	if index == 0 || index > 1 {
		fmt.Println("Program only accepts one and only one argument. Check your command and try again.")
		os.Exit(1)
	}

	//Easy way to do things:
	/*
		bs, err := ioutil.ReadFile(fileToOpen)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println(string(bs))
	*/

	file, err := os.Open(fileToOpen) // For read access.
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)

}
