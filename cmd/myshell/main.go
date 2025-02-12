package main

import (
	"bufio"
	"fmt"
	"os"
)

var exitCommand = "exit 0"

func main() {

	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		var c, err = bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Println("error reading from standard input: ", err)
			os.Exit(1)
		}
		// fmt.Println("inout c", c)

		// without the \n
		// var commands = strings.Split(c, " ")
		var command = c[:len(c)-1]
		if command == exitCommand {
			// var firstArg = os.Args[0]
			// fmt.Println("first arg", firstArg)
			os.Exit(0)
		}
		fmt.Printf("%s: command not found\n", command)
	}
}
