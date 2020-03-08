package main

import "fmt"

func errorPointer() {
	fmt.Println("-------errorPointer--------")
	var out []*int
	for i := 0; i < 3; i++ {
		out = append(out, &i)
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
	fmt.Println("Addresses:", out[0], out[1], out[2])
}

func truePointer() {
	fmt.Println("-------truePointer--------")
	var out []*int
	for i := 0; i < 3; i++ {
		// Make a copy of i because it will be reassigned with each loop.
		iCopy := i
		out = append(out, &iCopy)
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
	fmt.Println("Addresses:", out[0], out[1], out[2])
}

func main() {
	errorPointer()
	truePointer()
}
