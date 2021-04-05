package main

import (
	"fmt"
	str "strings" // Package Alias

	"github.com/java-akademie/go1n-einfuehrung/k0-packages/numbers"
	"github.com/java-akademie/go1n-einfuehrung/k0-packages/strings"
	"github.com/java-akademie/go1n-einfuehrung/k0-packages/strings/greeting"
)

func main() {
	fmt.Println(numbers.IsPrime(19))
	fmt.Println(greeting.WelcomeText)
	fmt.Println(strings.Reverse("java-akademie"))
	fmt.Println(str.Count("Go is Awesome. I love Go", "Go"))
}
