package main

import (
	"errors"
	"log"
	"strings"
	"unicode"
)

type SSN string

const lengthAllowed = 9

var (
	ErrInvalidSSNLength  = errors.New("ssn is not 9 characters long")
	ErrInvalidSSNNumbers = errors.New("ssn has non-numeric digits")
	ErrInvalidSSNPrefix  = errors.New("invalid ssn prefix")
)

func validSsnLength(ssn SSN) (ssnlen int, err error) {
	runelen := len([]rune(ssn))
	if runelen != lengthAllowed {
		return runelen, ErrInvalidSSNLength
	} else {
		return runelen, nil
	}
}

func isNumbers(ssn SSN) error {

	for _, runeVal := range ssn {
		if !unicode.IsDigit(runeVal) {
			return ErrInvalidSSNNumbers
		}
	}
	return nil
}

func validSsnPrefix(ssn SSN) error {
	if strings.HasPrefix(string(ssn), "000") {
		return ErrInvalidSSNPrefix
	} else {
		return nil
	}
}

func main() {
	validateSSN := []string{"123-45-6789", "012-8-678", "000-12-0962", "999-33-3333", "087-65-4321", "123-45-zzzz"}
	log.Println("Checking data: ", validateSSN)
	log.Println("====================================================================")
	for idx, str := range validateSSN {
		log.Printf(`Validate %v, %v of %v`, str, idx, len(validateSSN))
		ssn := SSN(strings.ReplaceAll(str, "-", ""))
		_, err := validSsnLength(ssn)
		if err != nil {
			log.Printf("the value of %s cause an error: %v", str, err)
		}

		err = isNumbers(ssn)
		if err != nil {
			log.Printf("the value of %s cause an error: %v", str, err)
		}

		err = validSsnPrefix(ssn)
		if err != nil {
			log.Printf("the value of %s cause an error: %v", str, err)
		}
	}
}
