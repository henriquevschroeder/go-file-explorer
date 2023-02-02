package commands

import (
	"fmt"
	"os"
)

<<<<<<< HEAD
func Rm(fileNames []string) {
	for _, fileName := range fileNames {
		err := os.Remove(fileName)
	
		if err != nil {
			fmt.Printf("Error removing file '%s': %s", fileName, err)
			os.Exit(1)
		}
	
		fmt.Println("Removed file:", fileName)
	}

=======
func Rm(filename string) {
	err := os.Remove(filename)

	if err != nil {
		fmt.Println("Error removing file:", err)
		os.Exit(1)
	}

	fmt.Println("Removed file:", filename)
>>>>>>> f797011b5fa32a1fbd65fdbb79aa2754767c4680
}