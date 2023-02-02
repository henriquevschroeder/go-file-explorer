package commands

import (
	"fmt"
	"os"
)

func Cd(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: cd <directory>")
		return
	}

	path := args[0]

	err := os.Chdir(path)

	if err != nil {
		fmt.Printf("Error changing directory: %s\n", err)
	}
}