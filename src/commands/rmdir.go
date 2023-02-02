package commands

import (
	"fmt"
	"os"
)

<<<<<<< HEAD
func Rmdir(dirNames []string) {
	for _, dirName := range dirNames {
		err := os.RemoveAll(dirName)
	
		if err != nil {
			fmt.Printf("Error removing directory '%s': %s", dirName, err)
			os.Exit(1)
		}
	
		fmt.Println("Removed directory:", dirName)
	}
=======
func Rmdir(dirname string) {
	err := os.RemoveAll(dirname)

	if err != nil {
		fmt.Println("Error removing directory:", err)
		os.Exit(1)
	}

	fmt.Println("Removed directory:", dirname)
>>>>>>> f797011b5fa32a1fbd65fdbb79aa2754767c4680
}