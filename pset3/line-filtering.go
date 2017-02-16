package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ^D or EOF is a great way to end Stdin type programs
func main() {
	arg, _ := strconv.Atoi(os.Args[1])
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// FUCK YEAH!
		line, _ := strconv.Atoi(scanner.Text())
		fmt.Println("randomly generated number:", line)
		if line == arg {
			fmt.Println("FOUND IT!")
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
