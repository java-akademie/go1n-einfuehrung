package main

import (
	"fmt"

	tools "github.com/java-akademie/myutils"
)

func testRunes() {
	tools.H2("testRunes")

	for i := 0; i <= 65535; i++ {
		if i <= 255 {
			fmt.Printf("%-3v %-5v %v", i, string(rune(i)), nl(i, 10))
		} else {
			if i == 256 {
				tools.Wait()
				fmt.Println()
			}
			fmt.Printf("%-2v %v", string(rune(i)), nl(i, 20))
		}
	}
	tools.Wait()
}

func nl(i int, max int) string {
	if i%max == 0 {
		return "  \n"
	}
	return "  "
}
