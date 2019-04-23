package phonenumber

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type NANP struct {
	number   string
	areaCode string
}

func newNANP(phone string) (*NANP, error) {
	replacer := strings.NewReplacer("(", "", ")", "", "-", "", "+", "", ".", "", " ", "")
	number := replacer.Replace(phone)
	if len(number) > 10 {
		countryCode := number[:len(number)-10]
		if countryCode != "1" {
			return nil, errors.New("invalid country code")
		}
		number = number[len(number)-10:]
	}

	for _, n := range number {
		if !unicode.IsNumber(n) {
			return nil, errors.New("invalid number")
		}
	}

	areaCode := number[:3]
	if areaCode[0] < '2' || number[3] < '2' {
		return nil, errors.New("invalid number")
	}
	return &NANP{number, areaCode}, nil
}

func (n NANP) String() string {
	return fmt.Sprintf("(%s) %s-%s", n.number[0:3], n.number[3:6], n.number[6:])
}

func Number(input string) (string, error) {
	n, err := newNANP(input)
	if err != nil {
		return "", err
	}

	return n.number, nil
}

func AreaCode(input string) (string, error) {
	n, err := newNANP(input)
	if err != nil {
		return "", err
	}

	return n.areaCode, nil
}

func Format(input string) (string, error) {
	n, err := newNANP(input)
	if err != nil {
		return "", err
	}

	return n.String(), nil
}
