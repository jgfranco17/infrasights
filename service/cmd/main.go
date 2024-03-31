package main

import (
	"flag"
	"fmt"
)

func main() {
	command := flag.String("command", "help", "Command to run")

	if flag.NArg() == 0 {
		fmt.Printf("Hello, user!\n")
	} else {
		switch *command {
		case "help":
			fmt.Printf("Welcome to infrasights!\n")
		case "run":
			fmt.Printf("Running a command!\n")
		}
	}
}
