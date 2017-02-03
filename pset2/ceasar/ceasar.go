package main

import "os"
import "fmt"
import "strconv"
import "bufio"

func main() {
	vargs := os.Args

	// implement
	// c_i = (p_i + k)mod26
	// where
	//	c_i is the i'th ciphertext
	//	p_i is the p'th plaintext
	//  and k is the rotation integer entered

	// A ASCII integer is 65 so lets say A-65 = 0
	//plaintext:  be sure to drink your Ovaltine
	//  98 101|115 117 114 101|116 111|
	//ciphertext: or fher gb qevax lbhe Binygvar
	// 111 114|102 104 101 114|103  98

	// O -> B 79  -> 66
	// v -> i 118 -> 105
	// a -> n 97  -> 110
	// l -> y 121 ->
	// t -> g 103 ->
	// i -> v 118 ->
	// n -> a 97  ->
	// e -> r 114 ->

	if len(vargs) > 1 {
		i := vargs[1]
		k, err := strconv.Atoi(i)
		// if j is 0 there should be no switch
		if err != nil {
			panic(err)
		}
		fmt.Print("plaintext: ")
		scanner := bufio.NewReader(os.Stdin)
		line, err := scanner.ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Print("ciphertext: ")

		for _, p_i := range line {
			p := int(p_i)
			if p == 32 {
				fmt.Printf(string(p))
			}
			if p > 64 && p < 91 {
				c_i := p + k%26
				if c_i > 90 {
					c_i -= 26
				}
				fmt.Print(string(c_i))
			} else if p > 96 && p < 123 {
				c_i := p + k%26
				if c_i > 122 {
					c_i -= 26
				}
				fmt.Print(string(c_i))
			} else {
				fmt.Print(string(p))
			}
		}
	} else {
		panic("You must provide more than one argument.\nTry a number between 1-26")
	}
	fmt.Println()

}
