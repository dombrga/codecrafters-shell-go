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
MAIN_LOOP:
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
					var _path = os.Getenv("PATH")
					var path = strings.Split(_path, ":")

					for _, p := range path {
						ls, err := os.ReadDir(p)
						if err != nil {
							// fmt.Fprintf(os.Stdout, "%s: not found\n", typeArg)
							continue
							// fmt.Fprintln(os.Stdout, "error ReadDir:", err, p)
							// continue MAIN_LOOP
						}

						for _, entry := range ls {
							entryName := entry.Name()
							if entryName == typeArg {
								fmt.Fprintf(os.Stdout, "%s is %s/%s\n", typeArg, p, typeArg)
								continue MAIN_LOOP // once found, don't run below by continuing MAIN_LOOP
							}
						}
					}

					fmt.Fprintf(os.Stdout, "%s: not found\n", typeArg)
				}

			}
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
		}
	}
}
