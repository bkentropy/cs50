package main

import "os"
import "fmt"

import "bufio"

func alphabetKey() map[string]int {
	alphaKey := make(map[string]int)
	count := 0
	for i := 65; i < 91; i++ {
		alphaKey[string(i)] = count
		count++
	}
	count = 0
	for j := 97; j < 123; j++ {
		alphaKey[string(j)] = count
		count++
	}
	return alphaKey
}

func algo(k_i, kIdx int, line, keyword string) {
	aKey := alphabetKey()

	kLen := len(keyword)
	k_inter := keyword[kIdx]
	k_i = aKey[string(k_inter)]

	for _, p_i := range line {
		p := int(p_i)
		if p > 64 && p < 91 {
			if kIdx < kLen {
				k_inter = keyword[kIdx]
				k_i = aKey[string(k_inter)]
				kIdx += 1
				if kIdx == kLen {
					kIdx = 0
				}
			}
			c_i := p + int(k_i)%26
			if c_i > 90 {
				c_i -= 26
			}
			fmt.Print(string(c_i))
		} else if p > 96 && p < 123 {
			if kIdx < kLen {
				k_inter = keyword[kIdx]
				k_i = aKey[string(k_inter)]
				kIdx += 1
				if kIdx == kLen {
					kIdx = 0
				}
			}
			c_i := p + int(k_i)%26
			if c_i > 122 {
				c_i -= 26
			}
			fmt.Print(string(c_i))
		} else {
			fmt.Print(string(p))
		}
	}
}

func main() {
	vargs := os.Args
	if len(vargs) > 1 {
		keyword := vargs[1]
		fmt.Print("plaintext: ")
		scanner := bufio.NewReader(os.Stdin)
		line, err := scanner.ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Print("ciphertext: ")

		algo(0, 0, line, keyword)

	} else {
		panic("You must provide a single word string argument.\n")
	}
	fmt.Println()

}
