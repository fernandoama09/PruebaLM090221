package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your text")
	fmt.Println("---------------------")
	text,_ := reader.ReadString('\n')

	fmt.Println("Enter your inTextFormat")
	fmt.Println("---------------------")
	textIn,_ := reader.ReadString('\n')
	textIn = strings.ToLower(textIn)
	fmt.Println("Enter your outTextFormat")
	fmt.Println("---------------------")
	textout,_ := reader.ReadString('\n')
	textout = strings.ToLower(textout)
	process,err	:=getProcessFactory(textIn,textout)
	if err!=nil{
		fmt.Print(err.Error())
	}else{
		resp := process.Process(text)
		fmt.Print(resp)
	}

}

func stringToBin(s string) (binString string) {
	for i, c := range s {
		if i!=0{
			binString+=" "
		}
		binString = fmt.Sprintf("%s%b",binString, c)
	}
	return
}

func textToMorse(s string) string{
	var binString string
	arr:= strings.Split(s," ")
	for i, c := range arr {
		if i!=0{
			binString+=" "
		}
		binString += stringToMorse[c]
	}
	return ""
}

var stringToMorse = map[string]string{
	"A":	".-",
	"B":	"-...",
	"C":	"-.-.",
	"D":	"-..",
	"E":	".",
	"F":	"..-.",
	"G":	"--.",
	"H":	"....",
	"I":	"..",
	"J":	".---",
	"K":	"-.-",
	"L":	".-..",
	"M":	"--",
	"N":	"-.",
	"O":	"---",
	"P":	".--.",
	"Q":	"--.-",
	"R":	".-.",
	"S":	"...",
	"T":	"-",
	"U":	"..-",
	"V":	"...-",
	"W":	".--",
	"X":	"-..-",
	"Y":	"-.--",
	"Z":	"--..",

}