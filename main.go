package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hey")
}

func Hello(name string) (string, error) {
	return name, nil
}
