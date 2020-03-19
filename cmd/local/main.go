package main

import (
	"fmt"
	"magic-eightball/package/eightball"
	"os"
	"strings"
)

func main() {
	question := "No question? Well here's an answer anyway!"
	if len(os.Args) > 1 {
		question = strings.Join(os.Args[1:], " ")
		fmt.Printf("You asked: %s\n", question)
	} else {
		fmt.Println(question)
	}
	legitanswer := eightball.DetermineMyFortune(question)
	fmt.Println(legitanswer)
}
