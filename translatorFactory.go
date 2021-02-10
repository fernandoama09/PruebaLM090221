package main

import "fmt"

type iValidator interface {
	validate(string) bool
}

type iTranslator interface {
	translate(string) string
}

type iProcess interface {
	Process(text string) string
}

type Process struct {
	Validator iValidator
	Translator iTranslator
}

func(p Process) Process(text string) string{
	if(p.Validator.validate(text)){
		return p.Translator.translate(text)
	}
	return ""
}

func getProcessFactory(inTextFormat string, outTextFormat string) (iProcess, error){
	var process Process
	var err error
	process.Translator, err= getTranslatorFactory(outTextFormat)
	if err!=nil{
		return nil, err
	}
	process.Validator,err = getValidatorFactory(inTextFormat)
	if err!=nil{
		return nil, err
	}
	return &process, nil
}
func getValidatorFactory(brand string) (iValidator, error) {

	if brand == "binario" {
		return &Binary{}, nil
	}
	if brand == "texto" {
		return &Text{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}

func getTranslatorFactory(brand string) (iTranslator, error) {

	if brand == "binario" {
		return &Binary{}, nil
	}
	if brand == "morse" {
		return &Binary{}, nil
	}
	if brand == "texto" {
		return &Text{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}
