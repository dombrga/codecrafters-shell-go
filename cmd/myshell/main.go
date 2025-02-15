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
var builtinCmds = map[string]string{
	echoCommand: echoCommand,
	exitCommand: exitCommand,
	typeCommand: typeCommand,
}

func main() {
	// MAIN_LOOP:
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
				typeArg := split[1]
				var biCmd, ok = builtinCmds[typeArg] // biCmd is a builtin command that is the argument of type command

				if ok {
					fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", biCmd)
				} else {
					isInCmd, p := isCmdInPath(typeArg)
					if isInCmd {
						fmt.Fprintf(os.Stdout, "%s is %s/%s\n", typeArg, p, typeArg)
					} else {
						fmt.Fprintf(os.Stdout, "%s: not found\n", typeArg)
					}
				}

			}
		default:
			// external programs in PATH
			// args := os.Args
			// // cmd := args[0]
			// fmt.Printf("Program was passed %d args including program name.\n", len(args))
			// for i, arg := range args {
			// 	fmt.Printf("Arg #%d: %s\n", i+1, arg)
			// }

			// cmd not found
			fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
		}
	}
}

// returns true if in PATH and its absolute path
func isCmdInPath(cmd string) (bool, string) {
	var _PATH = os.Getenv("PATH")
	var paths = strings.Split(_PATH, string(os.PathListSeparator))

	// loop all paths
	for _, p := range paths {
		stat, err := os.Stat(p + "/" + cmd)
		if err != nil || stat == nil {
			continue
		}

		return true, p
	}

	return false, ""
}
