package main

import "fmt"

func ErrorPointer() {
	fmt.Println("-------ErrorPointer--------")
	var out []*int
	for i := 0; i < 3; i++ {
		out = append(out, &i)
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
	fmt.Println("Addresses:", out[0], out[1], out[2])
}

func TruePointer() {
	fmt.Println("-------TruePointer--------")
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
	ErrorPointer()
	TruePointer()
}
