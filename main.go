package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println("Hey")
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("")
	}
	return name, nil

}
