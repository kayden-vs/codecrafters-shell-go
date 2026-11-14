package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
		var path string
		if firstWord == "type" {
			buildin := []string{"echo", "exit", "type", "pwd"}
			if len(words) == 2 {
				for _, s := range buildin {
					if s == words[1] {
						fmt.Println(s + " is a shell builtin")
						break
					}
				}
				continue
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

		// exec ext cmd
		extcmd := exec.Command(firstWord, words[1:]...)
		if extcmd.Err == nil {
			output, err := extcmd.CombinedOutput()
			if err != nil {
				fmt.Println("Error executing the external command", err)
			}
			fmt.Print(string(output))
			continue
		}

		// pwd
		if firstWord == "pwd" {
			ex, err := os.Executable()
			if err != nil {
				panic(err)
			}
			exPath := filepath.Dir(ex)
			fmt.Println(exPath)
			continue
		}

		fmt.Println(command + ": command not found")
	}
}
