package commands

import (
	"fmt"
	"os"
)

func Rm(fileNames []string) {
	for _, fileName := range fileNames {
		err := os.Remove(fileName)

		if err != nil {
			fmt.Printf("Error removing file '%s': %s", fileName, err)
			os.Exit(1)
		}

		fmt.Println("Removed file:", fileName)
	}

}
