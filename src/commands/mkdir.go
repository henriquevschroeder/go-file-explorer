package commands

import (
	"fmt"
	"os"
)

<<<<<<< HEAD
func Mkdir(dirNames []string) {
	for _, dirName := range dirNames {
		err := os.Mkdir(dirName, 0755)

		if err != nil {
			fmt.Printf("Error creating directory '%s': %s", dirName, err)
		} else {
			fmt.Println("Created directory:", dirName)
		}
=======
func Mkdir(dirname string) {
	err := os.Mkdir(dirname, 0755)

	if err != nil {
		fmt.Println("Error creating directory:", err)
	} else {
		fmt.Println("Created directory:", dirname)
>>>>>>> f797011b5fa32a1fbd65fdbb79aa2754767c4680
	}
}