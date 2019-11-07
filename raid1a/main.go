package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	xStr := os.Args[1]
	yStr := os.Args[2]

	x, _ := strconv.Atoi(xStr)
	y, _ := strconv.Atoi(yStr)

	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			if i == 1 || i == y {
				for j := 1; j <= x; j++ {
					if j == 1 || j == x {
						z01.PrintRune('o')
					} else {
						z01.PrintRune('-')
					}
				}
				z01.PrintRune('\n')
			} else {
				for j := 1; j <= x; j++ {
					if j == x || j == 1 {
						z01.PrintRune('|')
					} else {
						z01.PrintRune(' ')
					}
				}
				z01.PrintRune('\n')
			}
		}
	}
}
