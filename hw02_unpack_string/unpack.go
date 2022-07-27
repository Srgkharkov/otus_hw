package hw02unpackstring

//package hw02unpackstring

import (
	"errors"
	"fmt"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Concatandrepeatrune(strPtr *string, symbol rune, count int) {
	var str string
	if symbol == 0 {
		return
	}
	for i := 0; i < count; i++ {
		str += string(symbol)
	}
	*strPtr += str
}

func Unpack(str string) (string, error) {
	var ostr string
	var cursymbol rune
	var isControl bool

	for _, c := range str {
		fmt.Printf("For %q code %d:\n", c, int(c))

		if c == 92 {
			if isControl {
				cursymbol = c
				isControl = false
			} else {
				Concatandrepeatrune(&ostr, cursymbol, 1)
				cursymbol = 0
				isControl = true
			}
		}
		if unicode.IsLetter(c) {
			if isControl {
				return "", ErrInvalidString
			}

			Concatandrepeatrune(&ostr, cursymbol, 1)
			cursymbol = c
			isControl = false
		}
		if unicode.IsDigit(c) {
			if isControl {
				cursymbol = c
				isControl = false
			} else {
				if cursymbol == 0 {
					return "", ErrInvalidString
				}

				Concatandrepeatrune(&ostr, cursymbol, int(c-'0'))
				cursymbol = 0
			}
		}
	}
	Concatandrepeatrune(&ostr, cursymbol, 1)
	println(ostr)
	return ostr, nil
}
