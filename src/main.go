package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go-file-explorer/src/commands"
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
			commands.Cd(args[1:])
		case "ls":
			commands.Ls()
		case "mkdir":
			commands.Mkdir(args[1:])
		case "rmdir":
			commands.Rmdir(args[1:])
		case "mkfile":
			commands.Mkfile(args[1:])
		case "rm":
			commands.Rm(args[1:])
		default:
			fmt.Printf("Unknown command: %s\n", args[0])
		}
	}
}