package main

import (
	"bufio"
	"os"
	"fmt"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")
	text, _ := reader.ReadString('\n')
	fmt.Print(text)
}
