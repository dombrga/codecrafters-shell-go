package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var exit0Command = "exit 0"
var echoCommand = "echo"

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
		var input = strings.TrimSpace(c)
		var split = strings.Split(input, " ")
		var command = split[0]

		switch command {
		case exit0Command:
			os.Exit(0)
		case echoCommand:
			fmt.Fprintln(os.Stdout, strings.Join(split[1:], " "))
		default:
			fmt.Printf("%s: command not found\n", input)
		}
	}
}
