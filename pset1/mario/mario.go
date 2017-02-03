package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter number: ")
	num, _ := reader.ReadString('\n')
	for i, v := range num {
		fmt.Println(i, v)
	}

	// this will be user generated
	target, _ := strconv.Atoi(num)
	// this is how many #'s to print
	position := target - 1
	// this will be the row loop
	for j := 0; j < target; j++ {
		var buf bytes.Buffer
		// this will be he column loop
		for i := 0; i < target; i++ {
			if i >= position {
				buf.WriteString("#")
			} else {
				buf.WriteString(" ")
			}
		}
		fmt.Println(buf.String())
		position -= 1
	}
}
