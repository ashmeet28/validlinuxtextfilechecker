package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var hasNonASCIIPrintableCharacters bool = false
	var hasLineFeed bool = false
	var hasHorizontalTab bool = false

	for _, v := range data {
		if v == 0xa {
			hasLineFeed = true
		} else if v == 0x9 {
			hasHorizontalTab = true
		} else if (v < 0x20) || (v > 0x7e) {
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
		fmt.Println("Non ASCII Printable Characters: YES")
	} else {
		fmt.Println("Non ASCII Printable Characters: NO")
	}
}
