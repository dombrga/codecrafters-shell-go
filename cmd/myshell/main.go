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
var pwdCommand = "pwd"
var cdCommand = "cd"
var builtinCmds = map[string]string{
	echoCommand: echoCommand,
	exitCommand: exitCommand,
	typeCommand: typeCommand,
	pwdCommand:  pwdCommand,
	cdCommand:   cdCommand,
}

var _PATH = os.Getenv("PATH")
var paths = strings.Split(_PATH, string(os.PathListSeparator))

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		var c, err = bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Println("error reading from standard input: ", err)
			os.Exit(1)
		}

		// without the \n
		var input = strings.TrimSpace(c)
		var command = strings.SplitN(input, " ", 2)[0]

		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)

		// switch command {
		// case "":
		// runEmpty()
		// case exitCommand:
		// 	runExitCmd()
		// case echoCommand:
		// 	runEchoCmd(input)
		// case pwdCommand:
		// 	runPwdCmd()
		// case cdCommand:
		// 	runCdCmd(input)
		// case typeCommand:
		// 	runTypeCmd(input)
		// default:
		// 	runExtraCmd(input)
		// }
	}
}
