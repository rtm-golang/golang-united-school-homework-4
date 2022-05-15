// package string_sum
package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	// Use when the operands expected
	errorOperandExpected = errors.New("Operands Error: operand expected")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	si := -1
	o := make([]int, 0)
	m := 1
	b := false
	for i, r := range input + " " {
		if unicode.IsDigit(r) {
			if si == -1 {
				si = i
				if m == 0 {
					err = errors.New("Unexpected Error: operator (+/-) expected")
				}
			}
			b = false
		} else if r == '+' || r == '-' || r == ' ' {
			if si != -1 {
				n, e := strconv.Atoi(input[si:i])
				if e == nil {
					o = append(o, m*n)
				} else {
					err = fmt.Errorf("Unexpected Error: %w", e)
				}
				si = -1
			}
			if r == '-' {
				m = -1
				if b {
					err = errorOperandExpected
				}
				b = true
			} else {
				if r == '+' {
					m = 1
					if b {
						err = errorOperandExpected
					}
					b = true
				} else if !b {
					m = 0
				}
			}
		} else {
			err = errors.New("Input Error: not valid, only operator (+/-), spaces and digits allowed")
			si = -1
		}
		if len(o) > 2 {
			err = fmt.Errorf("Operands Error: %w", errorNotTwoOperands)
		}
	}
	if b {
		err = errorOperandExpected
	}
	if err != nil {
		return output, err
	}
	if len(o) == 1 {
		err = fmt.Errorf("Operands Error: %w", errorNotTwoOperands)
	} else if len(o) == 0 {
		err = fmt.Errorf("Input Error: %w", errorEmptyInput)
	}
	if err == nil {
		output = strconv.Itoa(o[0] + o[1])
	}
	return output, err
}

func main() {
	s := "5f+5"
	fmt.Println(StringSum(s))
}
