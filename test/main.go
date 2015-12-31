package main

import (
	"../"
	"fmt"
)

func main() {
	errIn := "This is an error!"
	fmt.Printf("in: %v\n", errIn)
	errOut := yogi.Catch(func() {
		panic(errIn)
	})
	fmt.Printf("out: %v\n", errOut)
	fmt.Printf("eq: %v\n", (errOut.Error() == errIn))

	c := 0
	d := &c
	err := yogi.Catch(func() {
		*d = 5
	})
	fmt.Printf("ptr: %v\n", d)
	fmt.Printf("value: %v\n", *d)
	fmt.Printf("no error: %v\n", err == nil)

	inner := func(shouldPanic bool) {
		if shouldPanic {
			panic("Panicking!")
		}
	}

	fmt.Println("Should NOT panic")
	err = yogi.Catch(func() {
		inner(false)
	})
	fmt.Printf("Panicking? %v\n", err != nil)

	fmt.Println("Should panic")
	err = yogi.Catch(func() {
		inner(true)
	})
	fmt.Printf("Panicking?: %v\n", (err.Error() == "Panicking!"))

}
