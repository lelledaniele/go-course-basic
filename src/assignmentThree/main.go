package main

import (
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		b := strings.Builder{}
		b.WriteString("Require 2 arguments, got ")
		b.WriteString(strconv.Itoa(len(os.Args))) // string(int) does not work

		log.Fatal(errors.New(b.String()))
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
