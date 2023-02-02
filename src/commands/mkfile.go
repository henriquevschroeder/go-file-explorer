package commands

import (
	"fmt"
	"os"
)

func Mkfile(filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}

	defer file.Close()
	fmt.Println("Created file:", filename)
}