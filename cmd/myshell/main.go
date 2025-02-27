package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
		var split = strings.Split(input, " ")
		var command = split[0]
		// var args = input[1:]

		// fmt.Println("asdzxc", builtinCmds[command], len(builtinCmds[command]))
		switch command {
		case "":
			fmt.Print("")
		case exitCommand:
			os.Exit(0)
		case echoCommand:
			// fmt.Fprintln(os.Stdout, strings.Join(split[1:], " "))
			runEchoCmd(input)
		case pwdCommand:
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stdout, "error printing working directory:", err)
			} else {
				fmt.Fprintln(os.Stdout, dir)
			}
		case cdCommand:
			if len(split) > 1 {
				runCdCmd(split[1])
			}
		case typeCommand:
			if len(split) > 1 {
				typeArg := split[1]
				_, ok := builtinCmds[typeArg]

				if ok {
					fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", typeArg)
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
			fmt.Println("defaulting", extractSingleQuoted(input))
			// external programs that are in PATH
			runCmd := exec.Command(command, extractSingleQuoted(input)...)
			runCmd.Stdout = os.Stdout
			runCmd.Stderr = os.Stderr
			err := runCmd.Run()
			if err != nil {
				if strings.Contains(err.Error(), exec.ErrNotFound.Error()) {
					fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
				}
			}
		}
	}
}

// it first checks if the echo argument is enclosed in single quotes
func runEchoCmd(input string) {
	if isSingleQuoted(input) {
		args := extractSingleQuoted(input)
		fmt.Fprintln(os.Stdout, strings.Join(args, ""))
	} else {
		fmt.Fprintln(os.Stdout, extractNonQuoted(input))
	}
}

func isSingleQuoted(input string) bool {
	split := strings.Split(input, " ")
	if len(split) > 1 {
		// split command and arguments
		s := strings.SplitAfterN(input, " ", 2)

		argSingleQuoted := strings.Split(s[1], "")
		if argSingleQuoted[0] == "'" && argSingleQuoted[len(argSingleQuoted)-1] == "'" {
			return true
		}
	}
	return false
}

func extractSingleQuoted(input string) []string {
	split := strings.Split(input, " ")
	if len(split) > 1 {
		// split command and arguments
		s := strings.SplitAfterN(input, " ", 2)

		argSingleQuoted := strings.Split(s[1], "")
		var args []string
		if argSingleQuoted[0] == "'" && argSingleQuoted[len(argSingleQuoted)-1] == "'" {
			arg := ""
			sQuote := 0
			for _, _s := range argSingleQuoted {
				// fmt.Println("single arg:", _s)
				if _s == "'" {
					// fmt.Println("single quote:", _s, sQuote)
					sQuote++
					if sQuote == 2 {
						args = append(args, arg)
						sQuote = 0
						arg = ""
					}
				} else {
					// fmt.Println("single not quote:", _s)
					arg = arg + _s
				}

				// fmt.Println("current arg:", arg)
			}

			// fmt.Println("args", len(args), args)
			// for _, s := range args {
			// 	fmt.Println("single arg", len(s), s)
			// }
			return args
			// return strings.Join(args, "")
		}
	}

	// return ""
	return []string{}
}

func extractNonQuoted(input string) string {
	echo := strings.Join(strings.Fields(input[len("echo")+1:]), " ")
	return echo
}

func runCdCmd(dir string) {
	var _dir = dir
	if dir == "~" {
		_dir = os.Getenv("HOME")
	}

	err := os.Chdir(_dir)
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", _dir)
	}
}

// returns true if in PATH and its absolute path
func isCmdInPath(cmd string) (bool, string) {
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
