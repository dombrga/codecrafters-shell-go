package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var exitCommand = "exit"
var exit0Command = "exit 0"
var echoCommand = "echo"
var typeCommand = "type"
var cmds = map[string]string{
	echoCommand: echoCommand,
	exitCommand: exitCommand,
	typeCommand: typeCommand,
}

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

		// without the \n
		var input = strings.TrimSpace(c)
		var split = strings.Split(input, " ")
		var command = split[0]

		switch command {
		case "":
			fmt.Print()
		case exitCommand:
			os.Exit(0)
			// if len(split) > 1 && split[1] == "0" {
			// 	os.Exit(0)
			// }
		case echoCommand:
			fmt.Fprintln(os.Stdout, strings.Join(split[1:], " "))
		case typeCommand:
			if len(split) > 1 {
				var arg, ok = cmds[split[1]]
				if ok {
					fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", arg)
				} else {
					fmt.Fprintf(os.Stdout, "%s: not found\n", split[1])
				}
			}
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
		}
	}
}
