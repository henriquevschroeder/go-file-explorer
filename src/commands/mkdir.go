package commands

import (
	"fmt"
	"os"
)

func Mkdir(dirname string) {
	err := os.Mkdir(dirname, 0755)

	if err != nil {
		fmt.Println("Error creating directory:", err)
	} else {
		fmt.Println("Created directory:", dirname)
	}
}