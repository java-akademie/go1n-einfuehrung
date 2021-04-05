package main

import "fmt"

type katze struct{}

func (t katze) bruellen() { fmt.Println("Miau") }
