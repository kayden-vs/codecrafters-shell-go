package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {

	for {
		fmt.Print("$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		command = command[:len(command)-1]
		words := strings.Fields(command)
		firstWord := words[0]

		if err != nil {
			fmt.Println("Problem reading input, ", err)
		}

		// exit
		if command == "exit" {
			break
		}

		//echo
		if firstWord == "echo" {
			fmt.Println(command[5:])
			continue
		}

		// type
		if firstWord == "type" {
			if len(words) == 2 {
				if words[1] == "echo" || words[1] == "exit" || words[1] == "type" {
					fmt.Println(words[1] + " is a shell builtin")
					continue
				}
			}
			fmt.Println(command[5:] + ": not found")
			continue
		}

		fmt.Println(command + ": command not found")
	}
}
