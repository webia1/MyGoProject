package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker(fullpath string) error {
	f, err := os.Open(fullpath)
	if err != nil {
		return fmt.Errorf("that's from fileChecker and this is from os.Open: %w", err)
	}
	f.Close()
	return nil
}

func main() {

	err := fileChecker("WrongNameWith.wrongExtension")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Custom Message instead of err: ", err)
		}
	}

	fmt.Println("Debugger")
}

/*

 */
