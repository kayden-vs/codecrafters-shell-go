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
		// firstWord := strings.Fields(command)[0]

		if err != nil {
			fmt.Println("Problem reading input, ", err)
		}

		// exit
		if command == "exit" {
			break
		}

		//echo
		if strings.Contains(command[:4], "echo") {
			// if firstWord == "echo" {
			fmt.Println(command[5:])
			continue
		}

		fmt.Println(command + ": command not found")
	}
}
