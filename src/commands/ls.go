package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func Ls() {
	root := "."

	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	entries, err := listEntries(root)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
}

func listEntries(root string) ([]os.FileInfo, error) {
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