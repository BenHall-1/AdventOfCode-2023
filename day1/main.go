package main

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var fEm embed.FS

func main() {
	part1()
	// part2()
}

func part1() {
	part1, part2 := 0, 0
	file, _ := fEm.ReadFile("input.txt")
	for _, line := range strings.Split(string(file), "\n") {
		part1Numbers := []int(nil)
		part2Numbers := []int(nil)
		for i := 0; i < len(line); i++ {
			if isDigit(line[i]) {
				part1Numbers = append(part1Numbers, int(line[i]-'0'))
				part2Numbers = append(part2Numbers, int(line[i]-'0'))
			} else if ok, val := isCharDigit(line, i); ok {
				part2Numbers = append(part2Numbers, val)
			}
		}
		if len(part1Numbers) > 0 {
			part1 += part1Numbers[0]*10 + part1Numbers[len(part1Numbers)-1]
		}
		if len(part2Numbers) > 0 {
			part2 += part2Numbers[0]*10 + part2Numbers[len(part2Numbers)-1]
		}

		// firstDigit, lastDigit := getFirstAndLastDigit(line)
		// add, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))
		// final += add
	}
	fmt.Printf("Part 1 answer: %d\n", part1)
	fmt.Printf("Part 2 answer: %d\n", part2)
}

func isDigit(u uint8) bool {
	return u >= '0' && u <= '9'
}

func isCharDigit(line string, i int) (bool, int) {
	for word, value := range charDigits {
		if i+len(word) > len(line) {
			continue
		}
		if line[i:i+len(word)] == word {
			return true, value
		}
	}

	return false, 0

}

var charDigits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// func convertString(str string) string {
// 	// Replace all keys with their values in the string based on conversionMap
// 	for k, v := range conversionMap {
// 		str = strings.ReplaceAll(str, k, v)
// 		// fmt.Println(str)
// 	}
// 	return str
// }

// var conversionMap = map[string]string{
// 	"one":   "1",
// 	"two":   "2",
// 	"three": "3",
// 	"four":  "4",
// 	"five":  "5",
// 	"six":   "6",
// 	"seven": "7",
// 	"eight": "8",
// 	"nine":  "9",
// }

// func getFirstAndLastDigit(s string) (int64, int64) {
// 	var firstDigit, lastDigit int64
// 	all := strings.Split(s, "")
// 	for i := range all {
// 		if digit, isNumeric := isNumeric(all[i]); isNumeric {
// 			firstDigit = digit
// 			break
// 		}
// 	}
// 	for i := len(all) - 1; i >= 0; i-- {
// 		if digit, isNumeric := isNumeric(all[i]); isNumeric {
// 			lastDigit = digit
// 			break
// 		}
// 	}
// 	return firstDigit, lastDigit
// }

// func isNumeric(s string) (int64, bool) {
// 	digit, err := strconv.ParseInt(s, 10, 64)
// 	return digit, err == nil
// }
