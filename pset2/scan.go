package main

import "os"
import "fmt"
import "bufio"

func scanIt() {
}

func main() {
	scanner := bufio.NewReader(os.Stdin)
	line, err := scanner.ReadString('\n')
	if err != nil {
		panic(err)
	}
	for _, s := range line {
		fmt.Println(s)
	}
}
