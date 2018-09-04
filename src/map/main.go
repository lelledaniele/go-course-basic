package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "FF0000",
		"green": "00FF00",
	}

	colors["white"] = "0000FF"

	delete(colors, "green")

	for i, c := range colors {
		fmt.Println(i, c)
	}
}
