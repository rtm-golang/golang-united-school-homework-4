package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	// Use when the expression invalid
	errorInvalidInput = errors.New("Expression Error: input invalid")
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

func ParseExpression(input string) (output []string, err error) {
	spaced := strings.ReplaceAll(strings.ReplaceAll(input, "+", " + "), "-", " - ")
	o := make([]string, 0)
	sign := 1
	for _, s := range strings.Split(spaced, " ") {
		if s == "+" {
			sign = 1
		} else if s == "-" {
			sign = -1
		} else if s == "" {
			//
		} else {
			if sign == 0 {
				return nil, errorInvalidInput
			} else if sign == 1 {
				o = append(o, s)
			} else if sign == -1 {
				o = append(o, "-"+s)
			}
			sign = 0
		}
	}
	if len(o) == 0 {
		return nil, errorEmptyInput
	} else if sign != 0 {
		return nil, errorInvalidInput
	}
	return o, nil
}

func StringSum(input string) (output string, err error) {
	s, e := ParseExpression(input)
	if e != nil {
		return "", e
	} else if len(s) != 2 {
		return "", errorNotTwoOperands
	}
	i := make([]int, 0)
	for _, o := range s {
		n, e := strconv.Atoi(o)
		if e != nil {
			return "", fmt.Errorf("Operand Error: %w", e)
		}
		i = append(i, n)
	}
	return strconv.Itoa(i[0] + i[1]), nil
}
