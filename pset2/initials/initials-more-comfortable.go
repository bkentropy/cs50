package main

import (
	"fmt"
	"regexp"
)

func main() {
	var re = regexp.MustCompile(`([a-zA-Z]+)`)
	var str = `___robert___thomas bowden`

	for i, match := range re.FindAllString(str, -1) {
		substr := string(match)
		fmt.Println(substr[0])
		//substr[0].toUpper()
		fmt.Println(substr, "found at index", i)
	}
}
