//Что выведется?

package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic caught:", r)
		}
	}()

	fmt.Println("Before panic")
	panic("Something went wrong!")
	fmt.Println("After panic")
}

// Before panic
// Panic caught: Something went wrong!
