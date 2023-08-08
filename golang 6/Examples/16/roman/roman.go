package roman

import (
	"errors"
	"strings"
)

var ErrInvalidFormat = errors.New("Invalid Format")

// Converts Roman number into Decimal
func RomanToDec(roman string) (result int, err error) {
	row := []rune(roman + ".")
	for i, r := range row {
		switch r {
		case 'M':
			result += 1000
		case 'D':
			result += 500
		case 'C':
			if row[i+1] == 'M' || row[i+1] == 'D' {
				result -= 100
			} else {
				result += 100
			}
		case 'L':
			result += 50
		case 'X':
			if row[i+1] == 'C' || row[i+1] == 'L' {
				result -= 10
			} else {
				result += 10
			}
		case 'V':
			result += 5
		case 'I':
			if row[i+1] == 'X' || row[i+1] == 'V' {
				result -= 1
			} else {
				result += 1
			}
		default:
			if r != '.' || i != len(row)-1 {
				return 0, ErrInvalidFormat
			}
		}
	}
	//Check the result by converting it back to the Roman system
	if s, err := DecToRoman(result); s != roman || err != nil {
		return 0, ErrInvalidFormat
	}
	return result, nil
}

var numeral = []int{1000, 500, 100, 50, 10, 5, 1}

// Converts Decimal into Roman
func DecToRoman(n int) (roman string, err error) {
	if n <= 0 || n >= 4000 {
		return "", errors.New("Out of range")
	}
	var letter string
	for _, v := range numeral {
		switch v {
		case 1000:
			letter = "M"
		case 500:
			letter = "D"
		case 100:
			letter = "C"
		case 50:
			letter = "L"
		case 10:
			letter = "X"
		case 5:
			letter = "V"
		case 1:
			letter = "I"
		}
		for n >= v {
			n -= v
			roman += letter
		}
	}
	roman = strings.Replace(roman, "DCCCC", "CM", 1)
	roman = strings.Replace(roman, "CCCC", "CD", 1)
	roman = strings.Replace(roman, "LXXXX", "XC", 1)
	roman = strings.Replace(roman, "XXXX", "XL", 1)
	roman = strings.Replace(roman, "VIIII", "IX", 1)
	roman = strings.Replace(roman, "IIII", "IV", 1)
	return roman, nil
}
