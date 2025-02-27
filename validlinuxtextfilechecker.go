package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var hasNonASCIIPrintableCharacters bool = false
	var hasLineFeed bool = false
	var hasHorizontalTab bool = false

	var lineFeedCount int = 0
	var firstNonASCIIPrintableCharacterLine int = 0

	for _, v := range data {
		if v == 0xa {
			hasLineFeed = true
			lineFeedCount++
		} else if v == 0x9 {
			hasHorizontalTab = true
		} else if (v < 0x20) || (v > 0x7e) {
			if !hasNonASCIIPrintableCharacters {
				firstNonASCIIPrintableCharacterLine = lineFeedCount
			}
			hasNonASCIIPrintableCharacters = true
		}
	}

	if hasLineFeed {
		fmt.Println("Line Feed: YES")
	} else {
		fmt.Println("Line Feed: NO")
	}

	if hasHorizontalTab {
		fmt.Println("Horizontal Tab: YES")
	} else {
		fmt.Println("Horizontal Tab: NO")
	}

	if hasNonASCIIPrintableCharacters {
		fmt.Println("Non ASCII Printable Characters: YES" + " " + "(" + "Line" + " " +
			strconv.FormatInt(int64(firstNonASCIIPrintableCharacterLine+1), 10) + ")")
	} else {
		fmt.Println("Non ASCII Printable Characters: NO")
	}
}
