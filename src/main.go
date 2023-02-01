package main

import ( 
    "fmt"
    "os"
    
    "go-file-explorer/src/filemgmt"
)

func main() {
	
	root := "."
	
	if len(os.Args) > 1 {
		root = os.Args[1]
	}
	
	entries, err := filemgmt.ListEntries(root)
	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}

}
