package commands

import (
	"fmt"
	"os"
)

func Mkfile(fileNames []string) {
	for _, fileName := range fileNames {
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	
		if err != nil {
			fmt.Printf("Error creating file '%s': %s", fileName, err)
			os.Exit(1)
		}
	
		defer file.Close()

		fmt.Println("Created file:", fileName)
	}
}