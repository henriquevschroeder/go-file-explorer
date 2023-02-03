package commands

import (
	"fmt"
	"os"
)

func Rmdir(dirNames []string) {
	for _, dirName := range dirNames {
		err := os.RemoveAll(dirName)

		if err != nil {
			fmt.Printf("Error removing directory '%s': %s", dirName, err)
			os.Exit(1)
		}

		fmt.Println("Removed directory:", dirName)
	}
}
