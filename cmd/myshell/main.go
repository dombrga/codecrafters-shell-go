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
var builtinCmds = map[string]string{
	echoCommand: echoCommand,
	exitCommand: exitCommand,
	typeCommand: typeCommand,
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
		// var args = split[1:]

		switch command {
		case "":
			fmt.Print()
		case exitCommand:
			os.Exit(0)
		case echoCommand:
			fmt.Fprintln(os.Stdout, strings.Join(split[1:], " "))
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
			// external programs that are in PATH
			runCmd := exec.Command(command, split[:1]...)
			runCmd.Stdout = os.Stdout
			runCmd.Stderr = os.Stderr
			err := runCmd.Run()
			if err != nil {
				fmt.Println("err runcmd:", err)
			} else {
				fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
			}

			// ok, _ := isCmdInPath(command)
			// if ok {
			// 	fmt.Printf("Program was passed %d args (including program name).\n", len(split))

			// 	for i, arg := range split {
			// 		if i == 0 {
			// 			fmt.Printf("Arg #%d (program name): %s\n", i, arg)
			// 		} else {
			// 			fmt.Printf("Arg #%d: %s\n", i, arg)
			// 		}
			// 	}

			// 	sig, err := rand.Int(rand.Reader, big.NewInt(9999999))
			// 	if err != nil {
			// 		fmt.Println("error in signature:", err)
			// 	}

			// 	fmt.Printf("Program Signature: %d\n", sig.Int64())
			// } else {
			// 	// cmd not found
			// 	fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
			// }
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
