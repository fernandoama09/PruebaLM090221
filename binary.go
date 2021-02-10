package main

import "fmt"

type Binary struct {
}

func (b Binary) validate(s string) bool {
	return true
}

func (b Binary) translate(s string) string {
	var binString string
	for i, c := range s {
		if i!=0{
			binString+=" "
		}
		binString = fmt.Sprintf("%s%b",binString, c)
	}
	return binString
}

