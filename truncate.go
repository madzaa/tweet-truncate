package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	f, err := os.ReadFile("cases")
	check(err)

	// init the first uint of bytes to keep
	var bytesToKeep = f[0]
	strBuild := strings.Builder{}
	var endString string

	// start from 1 to account for the initial byte
	for i := 1; i < len(f)-1; i++ {
		if i > len(f) {
			break
		}

		// truncate the the string once we find a newline
		if f[i] == 0x0a {
			endString += truncateRunes(strBuild.String(), bytesToKeep)
			bytesToKeep = f[i+1]
			i += 2
			strBuild.Reset()
		}

		strBuild.WriteByte(f[i])
	}

	// Process the last line
	endString += truncateRunes(strBuild.String(), bytesToKeep)
	fmt.Printf("%s", endString)
}

func truncateRunes(truncString string, maxLength uint8) string {
	if maxLength >= uint8(len(truncString)) {
		return truncString + "\n"
	}

	if maxLength <= 0 {
		return "\n"
	}

	for i := maxLength; i > 0; i-- {
		// utf8.RuneStart checks for b&0xC0 != 0x80, a multibyte character
		if utf8.RuneStart(truncString[i]) {
			return truncString[:i] + "\n"
		}
	}

	return truncString[1:maxLength] + "\n"
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
