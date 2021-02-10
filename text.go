package main

import "fmt"

type Text struct {
}

func (b Text) validate(s string) bool {
	return true
}

func (b Text) translate(s string) string {
	var binString string
	for i, c := range s {
		if i!=0{
			binString+=" "
		}
		binString = fmt.Sprintf("%s%b",binString, c)
	}
	return binString
}
