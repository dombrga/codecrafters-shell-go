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
		var command = strings.SplitN(input, " ", 2)[0]

		switch command {
		case "":
			runEmpty()
		case exitCommand:
			runExitCmd()
		case echoCommand:
			runEchoCmd(input)
		case pwdCommand:
			runPwdCmd()
		case cdCommand:
			runCdCmd(input)
		case typeCommand:
			runTypeCmd(input)
		default:
			runExtraCmd(input)
		}
	}
}

func runExtraCmd(input string) {
	// fmt.Println("extraa")
	var _input = strings.SplitN(input, " ", 2)
	var command = _input[0]
	args := extractSingleQuoted(input)

	if s, ok := args["quoteds"]; ok {
		// fmt.Println("extraa ok")
		// external programs that are in PATH
		runCmd := exec.Command(command, s...)
		runCmd.Stdout = os.Stdout
		runCmd.Stderr = os.Stderr

		err := runCmd.Run()
		if err != nil {
			if strings.Contains(err.Error(), exec.ErrNotFound.Error()) {
				fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
			}
		}
	} else {
		// external programs that are in PATH
		runCmd := exec.Command(command, _input[1])
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

func runTypeCmd(input string) {
	split := strings.SplitN(input, " ", 2)
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
}

func runPwdCmd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stdout, "error printing working directory:", err)
	} else {
		fmt.Fprintln(os.Stdout, dir)
	}
}

func runEmpty() {
	fmt.Print("")
}

func runExitCmd() {
	os.Exit(0)
}

// it first checks if the echo argument is enclosed in single quotes
func runEchoCmd(input string) {
	// var command = strings.SplitN(input, " ", 2)[0]
	if isSingleQuoted(input) {
		args := extractSingleQuoted(input)
		if s, ok := args["unquoteds"]; ok {
			fmt.Fprintln(os.Stdout, strings.Join(s, ""))
		}
	} else {
		fmt.Fprintln(os.Stdout, extractNonQuoted(input))
	}
}

func isSingleQuoted(input string) bool {
	split := strings.Split(input, " ")
	if len(split) > 1 {
		// split command and arguments
		s := strings.SplitN(input, " ", 2)

		argSingleQuoted := strings.Split(s[1], "")
		if argSingleQuoted[0] == "'" && argSingleQuoted[len(argSingleQuoted)-1] == "'" {
			return true
		}
	}
	return false
}

func extractSingleQuoted(input string) map[string][]string {
	split := strings.Split(input, " ")
	if len(split) > 1 {
		// split command and arguments
		s := strings.SplitN(input, " ", 2)

		argSingleQuoted := strings.Split(s[1], "")
		var quoteds []string
		var unquoteds []string
		if argSingleQuoted[0] == "'" && argSingleQuoted[len(argSingleQuoted)-1] == "'" {
			unquoted := "" // the string inside a pair of single quotes
			// quoted := ""
			sQuote := 0
			for _, _s := range argSingleQuoted {
				if _s == "'" {
					sQuote++
					if sQuote == 2 {
						unquoteds = append(unquoteds, unquoted)
						// quoteds = append(quoteds, "'"+unquoted+"'")
						quoteds = append(quoteds, unquoted)
						sQuote = 0
						unquoted = ""
					}
				} else {
					// only append to unquoted when there is a starting single quote already
					if sQuote == 1 {
						unquoted = unquoted + _s
					} else {
						// append space
						unquoteds = append(unquoteds, _s)
					}
				}
			}

			return map[string][]string{
				"quoteds":   quoteds,
				"unquoteds": unquoteds,
			}
		}
	}

	// return ""
	return map[string][]string{}
}

func extractNonQuoted(input string) string {
	echo := strings.Join(strings.Fields(input[len("echo")+1:]), " ")
	return echo
}

func runCdCmd(input string) {
	split := strings.Split(input, " ")
	dir := split[1]

	if len(split) > 1 {
		if dir == "~" {
			dir = os.Getenv("HOME")
		}

		err := os.Chdir(dir)
		if err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", dir)
		}
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
