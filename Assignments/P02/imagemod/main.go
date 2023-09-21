/************************************************************************
 *
 *  Author:           Brayden Lawson
 *  Title:            Baby Steps
 *  Course:           4143-101
 *  Semester:         Fall 2023
 *
 *  Description:
 * Golang program that will manipulate an image by drawing a black rectangle on it. 
 *  It has two directories imagemod and imagemanipulator. imageManipulator contains the 
 *   imageManipulator.cpp file while imagemod contains all of the other files as well as
 * imageManipulator since it's a sub directory of imagemod.
 *
 *
 *  Usage:
 *        Used to manipulate an existing image.
 *
 *
 *  Files: imageManipulator.go, go.mod, go.sum, main.go, mustangs_blank.png, mustangs.jpg, mustangs.png
 ************************************************************************/

package main

import (
	"fmt"
	"myimageapp/imageManipulator"
)

func main() {

	// Create an ImageManipulator instance. 
	// This is the orginal line of code that was used before the new constructor was made.
	//im := imageManipulator.NewImageManipulator(800, 600)
	// Create an ImageManipulator instance with an existing image
	im, err := imageManipulator.NewImageManipulatorWithImage("mustangs.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Draw a rectangle
	im.DrawRectangle(150, 50, 560, 411)

	// Save the image to a file
	im.SaveToFile("mustangs.png")

}
