package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {

	for {
		fmt.Print("$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		command = command[:len(command)-1]

		if err != nil {
			fmt.Println("Problem reading input, ", err)
		}

		if command == "exit" {
			break
		}

		fmt.Println(command + ": command not found")
	}
}
