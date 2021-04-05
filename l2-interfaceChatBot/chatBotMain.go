package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func main() {
	myutils.H1("chatBotMain with Interface")
	english := botEnglish{}
	spanish := botSpanish{}
	printGreeting(english)
	printGreeting(spanish)
}

type bot interface{ getGreeting() string }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
