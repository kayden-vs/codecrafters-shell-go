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
		if len(command) < 2 {
			continue
		}
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
			isBuiltin := false
			builtin := []string{"echo", "exit", "type", "pwd"}
			if len(words) == 2 {
				for _, s := range builtin {
					if s == words[1] {
						fmt.Println(s + " is a shell builtin")
						isBuiltin = true
						break
					}
				}
				if isBuiltin {
					continue
				}
			}

			// check path
			path, err = exec.LookPath(words[1])
			if err == nil {
				fmt.Println(words[1] + " is " + path)
				continue
			}

			fmt.Println(command[5:] + ": not found")
			continue
		}

		// pwd
		if firstWord == "pwd" {
			wd, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			fmt.Println(wd)
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

		// cd (absolute path)
		if firstWord == "cd" {
			path := words[1]

			err = os.Chdir(path)
			if err != nil {
				fmt.Print("cd: " + path + ": No such file or directory \n")
			}
			continue
		}

		fmt.Println(command + ": command not found")
	}
}
