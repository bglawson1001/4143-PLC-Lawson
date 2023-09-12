/************************************************************************
 *
 *  Author:           Brayden Lawson
 *  Title:            Run a Go program
 *  Course:           4143-101
 *  Semester:         Fall 2023
 *
 *  Description:
 * Basic Golang program that checks to see what mascot is returned in a mascot.go file.
 * Main.go prints out the name of the mascot along with a quote if the if statement in the
 *  mascot_test.go statement is not false.
 *
 *
 *  Usage:
 *        Shows how to import packages, write functions, and write a test file in Golang.
 *
 *
 *  Files: mascot_test.go, mascot.go, main.go
 ************************************************************************/

package main

import (
	"fmt" // this package formats basic strings, values, inputs, and outputs. 
	      // It can also be used to print and write from the terminal.

	"example.com/p01/mascot" 
	"rsc.io/quote" // package from the internet that allows us to print a quote
	               // to the screen. 
)

func main() {
	fmt.Println(mascot.BestMascot()) // prints the name of the mascot to the screen.
	fmt.Println(quote.Go()) // prints a quote to the screen.
}
