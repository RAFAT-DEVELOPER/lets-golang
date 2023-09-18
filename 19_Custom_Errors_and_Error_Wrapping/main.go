package main

import (
	"errors"
	"fmt"

	"github.com/pkg/errors"
)

func doSomething() error {
	return errors.Wrap(errors.New("original error"), "doSomething failed")
}

func main() {
	if err := doSomething(); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("No errors")
}
