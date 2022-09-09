package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := returnError(true)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func returnError(returnError bool) (string, error) {
	if returnError {
		return "", errors.New("error here")
	}
	return "Random result", nil
}
