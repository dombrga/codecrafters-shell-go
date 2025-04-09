package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var exitCmd = "exit"
var exit0Command = "exit 0"
var echoCommand = "echo"
var typeCommand = "type"
var pwdCommand = "pwd"
var cdCommand = "cd"
var builtinCmds = map[string]string{
	echoCommand: echoCommand,
	exitCmd:     exitCmd,
	typeCommand: typeCommand,
	pwdCommand:  pwdCommand,
	cdCommand:   cdCommand,
}

var _PATH = os.Getenv("PATH")
var paths = strings.Split(_PATH, string(os.PathListSeparator))

func main() {
	startRepl()
}

func startRepl() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		var c, err = bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Println("error reading from standard input: ", err)
			os.Exit(1)
		}

		// without the \n
		input := strings.TrimSpace(c)
		command := getCommand(input)
		// command := _split[0]
		// _args := _split[1]

		switch command {
		case "":
			fmt.Fprintf(os.Stdout, "\n")
		case exitCmd:
			runExitCmd(input)
		case echoCommand:
			runEchoCmd(input)
		default:
			runInvalidCmd(input)
		}
	}
}

func getCommand(input string) string {
	// return strings.Fields(input)[0]
	return strings.SplitN(input, " ", 2)[0]
}

func runEchoCmd(input string) {
	fmt.Fprintf(os.Stdout, "%s\n", GetEchoPrint(input))
}

func GetEchoPrint(input string) string {
	args := getEchoArg(input)
	return fmt.Sprintf(args)
}

func getEchoArg(input string) string {
	return strings.TrimSpace(strings.SplitN(input, " ", 2)[1])
}

func runInvalidCmd(input string) {
	fmt.Fprintf(os.Stdout, "%s\n", GetInvalidPrint(input))
}

func GetInvalidPrint(input string) string {
	return fmt.Sprintf("%s: command not found", input)
}

func runExitCmd(input string) {
	os.Exit(0)
}
