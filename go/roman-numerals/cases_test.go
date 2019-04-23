package romannumerals

import (
	"testing"
)

// Source: exercism/problem-specifications
// Commit: 3c78ac4 Add test case for input value 49 to test double-normalization
// Problem Specifications Version: 1.2.0

type romanNumeralTest struct {
	arabic   int
	roman    string
	hasError bool
}

var romanNumeralTests = []romanNumeralTest{
	{1, "I", false},
	{2, "II", false},
	{3, "III", false},
	{4, "IV", false},
	{5, "V", false},
	{6, "VI", false},
	{9, "IX", false},
	{27, "XXVII", false},
	{48, "XLVIII", false},
	{49, "XLIX", false},
	{59, "LIX", false},
	{93, "XCIII", false},
	{141, "CXLI", false},
	{163, "CLXIII", false},
	{402, "CDII", false},
	{575, "DLXXV", false},
	{911, "CMXI", false},
	{1024, "MXXIV", false},
	{3000, "MMM", false},
}

func TestCurDigit(t *testing.T) {
	//fmt.Println("num: 27, expect: 2, got: ", CurDigit(27, 2))
	//fmt.Println("num: 27, expect: 7: got: ", CurDigit(27, 1))
	//
	//fmt.Println("num: 123, expect: 1, got: ", CurDigit(123, 3))
	//fmt.Println("num: 123, expect: 2: got: ", CurDigit(123, 2))
	//fmt.Println("num: 123, expect: 3: got: ", CurDigit(123, 1))
	//
	//fmt.Println("num: 1234, expect: 1: got: ", CurDigit(1234, 4))
	//fmt.Println("num: 1234, expect: 2: got: ", CurDigit(1234, 3))
	//fmt.Println("num: 1234, expect: 3: got: ", CurDigit(1234, 2))
	//fmt.Println("num: 1234, expect: 4: got: ", CurDigit(1234, 1))
}
