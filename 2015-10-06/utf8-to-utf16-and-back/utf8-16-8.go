package main

import (
	"bytes"
	"fmt"
	"unicode/utf16"
	"unicode/utf8"
)

func main() {
	testString := "Hello world!"

	fmt.Printf("Original string: %s\n", testString)

	var runeArray []rune

	// string -> []rune
	for idx := 0; idx < len(testString); {
		rune, size := utf8.DecodeRuneInString(string(testString[idx]))

		// Handle the variable size of a UTF-8 code point...
		idx += size

		runeArray = append(runeArray, rune)
	}

	// []rune -> []uint16 (UTF-16 encoding)
	myWinString := utf16.Encode(runeArray)

	// Now the reverse: []uint16 -> []rune
	reverseRuneArray := utf16.Decode(myWinString)

	// And now []rune -> Go string
	outBuf := new(bytes.Buffer)

	for idx, _ := range reverseRuneArray {
		outBuf.WriteRune(reverseRuneArray[idx])
	}

	fmt.Printf("Back to the original: %s\n", string(outBuf.Bytes()))
}
