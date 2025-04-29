package main

import (
	"fmt"
	"log"
	"os"
)

var FatalLogger *log.Logger

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}
	FatalLogger = log.New(file, "Fatal: ", log.LstdFlags|log.Lshortfile)
}

func handlePanic() {
	a := recover()

	if a != nil {
		FatalLogger.Println(a)
		fmt.Println("Do not repeat")

	}
}

func division(num1, num2 int) {
	defer handlePanic()

	if num2 == 0 {
		panic("Cannot divide a number by zero.")
	} else {
		result := num1 / num2
		fmt.Println("Result:", result)
	}
}

func main() {
	division(10, 5)
	division(10, 0)
	division(10, 3)
}
