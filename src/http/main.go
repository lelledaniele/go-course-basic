package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type CustomWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")
	cw := CustomWriter{}

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// r := make([]byte, 99999)

	// resp.Body.Read(r)

	// fmt.Println(string(r))

	io.Copy(cw, resp.Body)
}

func (CustomWriter) Write(p []byte) (int, error) {
	return len(p), ioutil.WriteFile("ciao.txt", p, 0666)
}
