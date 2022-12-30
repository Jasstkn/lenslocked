package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func main() {
	err := B()
	if errors.Is(err, ErrNotFound) {
		fmt.Print("error is ErrNotFound")
	}
	fmt.Print("unknown error")
}

func A() error {
	return ErrNotFound
}

func B() error {
	err := A()
	if err != nil {
		return fmt.Errorf("b: %w", err)
	}
	return nil
}
