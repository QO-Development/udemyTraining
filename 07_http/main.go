package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

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

	lw := logWriter{}

	io.Copy(lw, resp.Body)
	fmt.Println("")
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
