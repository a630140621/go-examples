package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func foo() error {
	return fmt.Errorf("original error")
}

func bar() error {
	err := foo()
	if err != nil {
		return errors.Wrap(err, "bar")
	}

	return nil
}

func pong() error {
	err := bar()
	if err != nil {
		return errors.Wrap(err, "pong")
	}

	return nil
}

func main()  {
	if err := pong(); err != nil {
		// fmt.Println(err)
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	fmt.Printf("end\n")
}