package roman

import (
	"bytes"
	"errors"
	"regexp"
)

type romanNumPair struct {
	Roman string
	Num   int
}

var (
	romanNumParis []romanNumPair
	romanRegex    *regexp.Regexp
)

var (
	ErrOutOfRange   = errors.New("out of range")
	ErrInvalidRoman = errors.New("invalid roman")
)

func init() {
	romanNumParis = []romanNumPair{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}

	romanRegex = regexp.MustCompile(`^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`)
}

func ToRoman(n int) (string, error) {
	if n <= 0 || n >= 4000 {
		return "", ErrOutOfRange
	}
	var buf bytes.Buffer
	for _, pair := range romanNumParis {
		for n >= pair.Num {
			buf.WriteString(pair.Roman)
			n -= pair.Num
		}
	}

	return buf.String(), nil
}

func FromRoman(roman string) (int, error) {
	if !romanRegex.MatchString(roman) {
		return 0, ErrInvalidRoman
	}

	var result int
	var index int
	for _, pair := range romanNumParis {
		for roman[index:index+len(pair.Roman)] == pair.Roman {
			result += pair.Num
			index += len(pair.Roman)
		}
	}

	return result, nil
}
