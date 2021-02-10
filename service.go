package main

type Translator interface {
	Translate(text string, inFormat string, outFormat string)
}

