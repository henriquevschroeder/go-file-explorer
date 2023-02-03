package commands

import (
	"fmt"
	"os"
)

func Mkdir(dirNames []string) {
	for _, dirName := range dirNames {
		err := os.Mkdir(dirName, 0755)

		if err != nil {
			fmt.Printf("Error creating directory '%s': %s", dirName, err)
		} else {
			fmt.Println("Created directory:", dirName)
		}
	}
}
