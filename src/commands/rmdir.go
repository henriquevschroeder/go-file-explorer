package commands

import (
	"fmt"
	"os"
)

func Rmdir(dirname string) {
	err := os.RemoveAll(dirname)

	if err != nil {
		fmt.Println("Error removing directory:", err)
		os.Exit(1)
	}

	fmt.Println("Removed directory:", dirname)
}