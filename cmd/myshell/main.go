package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	var input, err = bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Println("error reading from standard input: ", err)
		os.Exit(1)
	}

	fmt.Printf("%s: command not found", input[:len(input)-1])
}
