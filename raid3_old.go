package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Get data from st input of raid1
	reader := bufio.NewReader(os.Stdin)
	output := []rune{}
	for {
		text, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		output = append(output, text)
	}

	//Check input
	if !inputCorrect(output) {
		notRaid1()
	}

	//Count rows and colomns and check for correct input
	rows := 0
	cols := 0
	colsCounted := false
	for index, char := range output {
		if char == 10 {
			rows++
			colsCounted = true
			if (index+1)%(cols+1) != 0 {
				notRaid1()
			}
		}
		if !colsCounted {
			cols++
		}
	}
	if cols < 1 || rows < 1 || output[len(output)-1] != 10 {
		notRaid1()
	}

	//Remove \n from output
	arr := []rune{}
	for _, char := range output {
		if char != 10 {
			arr = append(arr, char)
		}
	}

	//Check output is raid1 from a to e
	isRaid1 := [5]bool{}
	if checkRaids(getRaid1a, cols, rows, output) {
		isRaid1[0] = true
	}
	if checkRaids(getRaid1b, cols, rows, output) {
		isRaid1[1] = true
	}
	if checkRaids(getRaid1c, cols, rows, output) {
		isRaid1[2] = true
	}
	if checkRaids(getRaid1d, cols, rows, output) {
		isRaid1[3] = true
	}
	if checkRaids(getRaid1e, cols, rows, output) {
		isRaid1[4] = true
	}

	//Print result
	fits := 0
	for _, check := range isRaid1 {
		if check {
			fits++
		}
	}

	raidNames := [5]string{"raid1a", "raid1b", "raid1c", "raid1d", "raid1e"}
	solutions := 0
	for i := 0; i < 5; i++ {
		if isRaid1[i] {
			solutions++
			fits--
			fmt.Printf("[%v] [%v] [%v]", raidNames[i], cols, rows)
			if fits != 0 {
				fmt.Print(" || ")
			}
		}
	}
	if solutions == 0 {
		notRaid1()
	}
	fmt.Println()

	//Additional functional
	args := os.Args[1:]
	flags := [5]bool{}
	for _, param := range args {
		if param == "--printGraphic" || param == "-pg" {
			flags[0] = true
		} else if param == "--printRunes" || param == "-pr" {
			flags[1] = true
		} else if param == "--printAllRunes" || param == "-par" {
			flags[2] = true
		}
	}
	if flags[0] {
		fmt.Printf("\n%c\n", output)
	}
	if flags[1] {
		fmt.Println("\n", arr)
	}
	if flags[2] {
		fmt.Println("\n", output)
	}
}

func checkRaids(getRaid1 func(rows, cols int) []rune, rows, cols int, arr []rune) bool {
	etalon := getRaid1(rows, cols)
	for index, char := range etalon {
		if char != arr[index] {
			return false
		}
	}
	return true
}

func getRaid1a(x, y int) []rune {
	result := []rune{}
	for row := 0; row < y; row++ {
		for colomn := 0; colomn < x; colomn++ {
			if colomn == 0 && row == 0 || colomn == x-1 && row == y-1 || colomn == 0 && row == y-1 || colomn == x-1 && row == 0 {
				result = append(result, 'o')
			} else if colomn == 0 || colomn == x-1 {
				result = append(result, '|')
			} else if row == 0 || row == y-1 {
				result = append(result, '-')
			} else {
				result = append(result, ' ')
			}
			if colomn == x-1 {
				result = append(result, 10)
			}
		}
	}
	return result
}

func getRaid1b(x, y int) []rune {
	result := []rune{}
	for row := 0; row < y; row++ {
		for colomn := 0; colomn < x; colomn++ {
			if colomn == 0 && row == 0 || colomn == x-1 && row == y-1 && x != 1 && y != 1 {
				result = append(result, '/')
			} else if colomn == x-1 && row == 0 || colomn == 0 && row == y-1 {
				result = append(result, '\\')
			} else if colomn == 0 || colomn == x-1 || row == 0 || row == y-1 {
				result = append(result, '*')
			} else {
				result = append(result, ' ')
			}
			if colomn == x-1 {
				result = append(result, 10)
			}
		}
	}
	return result
}

func getRaid1c(x, y int) []rune {
	result := []rune{}
	for row := 0; row < y; row++ {
		for colomn := 0; colomn < x; colomn++ {
			if row == 0 && (colomn == x-1 || colomn == 0) {
				result = append(result, 'A')
			} else if row == y-1 && (colomn == x-1 || colomn == 0) {
				result = append(result, 'C')
			} else if colomn == 0 || colomn == x-1 || row == 0 || row == y-1 {
				result = append(result, 'B')
			} else {
				result = append(result, ' ')
			}
			if colomn == x-1 {
				result = append(result, 10)
			}
		}
	}
	return result
}

func getRaid1d(x, y int) []rune {
	result := []rune{}
	for row := 0; row < y; row++ {
		for colomn := 0; colomn < x; colomn++ {
			if colomn == 0 && row == 0 || colomn == 0 && row == y-1 {
				result = append(result, 'A')
			} else if colomn == x-1 && row == 0 || colomn == x-1 && row == y-1 {
				result = append(result, 'C')
			} else if colomn == 0 || row == 0 || colomn == x-1 || row == y-1 {
				result = append(result, 'B')
			} else {
				result = append(result, ' ')
			}
		}
		result = append(result, 10)
	}
	return result
}

func getRaid1e(x, y int) []rune {
	result := []rune{}
	for row := 0; row < y; row++ {
		for colomn := 0; colomn < x; colomn++ {
			if colomn == 0 && row == 0 || colomn == x-1 && row == y-1 && x != 1 && y > 1 {
				result = append(result, 'A')
			} else if colomn == x-1 && row == 0 || colomn == 0 && row == y-1 {
				result = append(result, 'C')
			} else if colomn == 0 || row == 0 || colomn == x-1 || row == y-1 {
				result = append(result, 'B')
			} else {
				result = append(result, ' ')
			}
		}
		result = append(result, 10)
	}
	return result
}

func inputCorrect(arr []rune) bool {
	for _, char := range arr {
		switch char {
		case 'o':
		case '-':
		case '|':
		case ' ':
		case '*':
		case '/':
		case '\\':
		case 'A':
		case 'B':
		case 'C':
		case 10:
		default:
			return false
		}
	}
	return true
}

func notRaid1() {
	fmt.Println("Not a Raid function")
	os.Exit(0)
}
