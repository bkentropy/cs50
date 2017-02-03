package main

import "fmt"
import "unicode/utf8"

func main() {
	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")
	fmt.Printf("with the + %+q\n", placeOfInterest) // => \u2318
	fmt.Printf("without the %q\n", placeOfInterest) // => ⌘
	fmt.Printf("\u2318\n")                          // => ⌘

	const nihongo = "日本語"
	//const nihongo = "􏿿" // loop returns: U+10FFFF starts at byte position 0
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	/*
		Notice the outputs position, it takes 3 bytes to make up a unicode (sometimes?)
		U+65E5 '日' starts at byte position 0
		U+672C '本' starts at byte position 3
		U+8A9E '語' starts at byte position 6

		[Exercise: Put an invalid UTF-8 byte sequence into the string. (How?) What happens to the iterations of the loop?]
		the result is the nihongo duplicate. It was a square looking shape and
		instead that ? shape shows.
	*/
	// Second nihongo loop
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}

}

/*
To summarize, here are the salient points:
	-Go source code is always UTF-8.
	-A string holds arbitrary bytes.
	-A string literal, absent byte-level escapes, always holds valid UTF-8 sequences.
	-Those sequences represent Unicode code points, called runes.
	-No guarantee is made in Go that characters in strings are normalized.
*/
