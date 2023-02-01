package main

import ( 
    "fmt"
    "os"
	"bufio"
	"strings"
    
    "go-file-explorer/src/filemgmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("⚡️ ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		if input == "exit" {
			break
		}

		args := strings.Split(input, " ")

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "cd":
			filemgmt.CdCommand(args[1:])
		case "ls":
			filemgmt.LsCommand()
		default:
			fmt.Printf("Unknown command: %s\n", args[0])
		}
	}
}
