package filemgmt

import (
	"os"
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