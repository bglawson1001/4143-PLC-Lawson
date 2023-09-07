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
 *
 *
 *
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
	"fmt"

	"example.com/p01/mascot"
	"rsc.io/quote"
)

func main() {
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())
}
