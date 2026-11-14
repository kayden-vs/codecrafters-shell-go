package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
		// validExec := false

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
		var path string
		if firstWord == "type" {
			if len(words) == 2 {
				if words[1] == "echo" || words[1] == "exit" || words[1] == "type" {
					fmt.Println(words[1] + " is a shell builtin")
					continue
				}
			}

			// check path
			path, err = exec.LookPath(words[1])
			if err == nil {
				// validExec = true
				fmt.Println(words[1] + " is " + path)
				continue
			}

			fmt.Println(command[5:] + ": not found")
			continue
		}

		// run the exec
		path, err = exec.LookPath(firstWord)
		if err == nil {
			extcmd := exec.Command(path, words[1:]...)
			output, err := extcmd.CombinedOutput()
			if err != nil {
				fmt.Println("Error executing the external command", err)
			}
			fmt.Println(string(output))
			continue
		}

		fmt.Println(command + ": command not found")
	}
}
