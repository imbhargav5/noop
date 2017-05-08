package main

import "fmt"

var num = 14

func generateA() func() {
	var a func()
	if num%2 == 0 {
		a = Noop
	} else {
		a = func() {
			fmt.Printf("This is soemething")
		}
	}
	return a
}

func main() {
	generateA()()
}
