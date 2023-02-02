package commands

import (
	"fmt"
	"os"
)

func Rm(filename string) {
	err := os.Remove(filename)

	if err != nil {
		fmt.Println("Error removing file:", err)
		os.Exit(1)
	}

	fmt.Println("Removed file:", filename)
}