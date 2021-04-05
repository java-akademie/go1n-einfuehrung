package main

import (
	"fmt"
	"reflect"
)

func show(i interface{}) {
	k := reflect.ValueOf(i).Kind()
	t := reflect.TypeOf(i)

	if k.String() == t.String() {
		format := "Type/Kind: %-16T Value: %v \n"
		fmt.Printf(format, i, i)
	} else {
		format := "Type: %-16v Kind: %-8v Value: %v \n"
		fmt.Printf(format, t, k, i)
	}

}
