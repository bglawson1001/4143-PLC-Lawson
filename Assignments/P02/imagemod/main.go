package main

import (
	"fmt"
	"myimageapp/imageManipulator"
)

func main() {

	// Create an ImageManipulator instance
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
