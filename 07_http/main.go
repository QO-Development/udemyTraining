package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// //Need to initialize byte slice so its big enough to fit data. Otherwise read function will stop because it thinks byte slice is full
	// bs := make([]byte, 99999)

	// i, er := resp.Body.Read(bs)

	// fmt.Println(i)
	// fmt.Println(er)
	// fmt.Println(string(bs))

	//Does the same as the above commented out code
	io.Copy(os.Stdout, resp.Body)
	fmt.Println("")
}
