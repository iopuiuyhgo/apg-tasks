package main

import "fmt"

//Исправить функцию, чтобы она работала.
//Сигнатуру менять нельзя

func printNumber(ptrToNumber interface{}) {

	if ptrToNumber == nil {
		fmt.Println("nil")
	}
	v, ok := ptrToNumber.(*int)
	if !ok || v == nil {
		fmt.Println("nil")
	} else {
		fmt.Println(*ptrToNumber.(*int))
	}
}

func main() {

	v := 10
	printNumber(&v)
	var pv *int

	printNumber(pv)
	pv = &v
	printNumber(pv)
}
