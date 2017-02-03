package main

import "fmt"
import "unicode"
import "bufio"
import "os"

func main() {

	initials := ""
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your full name with each first letter capitalized.")
	text, _ := reader.ReadString('\n')
	for _, v := range text {
		if unicode.IsUpper(v) {
			initials = initials + string(v)
		}
	}
	fmt.Println(initials)
}
