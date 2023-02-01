package filemgmt

import (
	"os"
	"fmt"
	"path/filepath"
)

func ListEntries(root string) ([]os.FileInfo, error) {
	entries, err := filepath.Glob(filepath.Join(root, "*"))

	if err != nil {
		return nil, err
	}

	var infos []os.FileInfo

	for _, entry := range entries {
		info, err := os.Stat(entry)

		if err != nil {
			return nil, err
		}

		infos = append(infos, info)
	}
	
	return infos, nil
}

func CdCommand(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: cd <directory>")
		return
	}
	
	path := args[0]
	err := changeDirectory(path)
	
	if err != nil {
		fmt.Printf("Error changing directory: %s\n", err)
	}
}

func LsCommand() {
	root := "."
	
	if len(os.Args) > 1 {
		root = os.Args[1]
	}
	
	entries, err := ListEntries(root)
	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
}

func MkdirCommand(dirname string) {
	err := os.Mkdir(dirname, 0755)

	if err != nil {
		fmt.Println("Error creating directory:", err)
	} else {
		fmt.Println("Created directory:", dirname)
	}
}

func RmdirCommand(dirname string) {
	err := os.RemoveAll(dirname)

	if err != nil {
		fmt.Println("Error removing directory:", err)
		os.Exit(1)
	}

	fmt.Println("Removed directory:", dirname)
}

func MkfileCommand(filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}

	defer file.Close()
	fmt.Println("Created file:", filename)
}

func RmCommand(filename string) {
	err := os.Remove(filename)

	if err != nil {
		fmt.Println("Error removing file:", err)
		os.Exit(1)
	}

	fmt.Println("Removed file:", filename)
}

func changeDirectory(path string) error {
	return os.Chdir(path)
}