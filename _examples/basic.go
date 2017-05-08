package main

import "fmt"
import "github.com/bhargav175/noop"

func returnNoopOrNot(val int) func() {
	var a func()
	if val%2 == 0 {
		a = noop.Noop
	} else {
		a = func() {
			fmt.Printf("%d results in not a noop", val)
		}
	}
	return a
}

func main() {
	returnNoopOrNot(14)()
	returnNoopOrNot(15)()
}
