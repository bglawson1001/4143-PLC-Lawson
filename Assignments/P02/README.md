## P02 - Baby Steps
### Brayden Lawson
### Description:

This program will draw a black rectangle on an existing image. 

### Files

|   #   | File     | Description                      |
| :---: | -------- | -------------------------------- |
|   1   | [main.go](https://github.com/bglawson1001/4143-PLC-Lawson/blob/main/Assignments/P02/imagemod/main.go) | Main driver of this Golang program. It imports an imageManipulator package, and modifies the original image by drawing a black rectangle. |
|   2   | [mascot](https://github.com/bglawson1001/4143-PLC-Lawson/blob/main/Assignments/P02/imagemod/imageManipulator/imageManipulator.go) | This file contains the necessary functions needed to draw the rectangle on the image..
|   3   | [go.mod](https://github.com/bglawson1001/4143-PLC-Lawson/blob/main/Assignments/P02/imagemod/go.mod) | This file is necessary because it declares the program's dependencies and their version number. 
|   4   | [go.sum](https://github.com/bglawson1001/4143-PLC-Lawson/blob/main/Assignments/P02/imagemod/go.sum) | This file maintains the checksum so when the program runs again it will not reinstall the packages, and instead use the cache. 
|   5   | [mustangs.jpg](https://github.com/bglawson1001/4143-PLC-Lawson/blob/main/Assignments/P02/imagemod/mustangs.jpg) | Original image.
|   7  | [mustangs.png](https://github.com/bglawson1001/4143-PLC-Lawson/blob/main/Assignments/P02/imagemod/mustangs.png) | Modified image with the black rectangle that was created as a result of adding the new constructor. 
|   6   | [mustangs_blank.png](https://github.com/bglawson1001/4143-PLC-Lawson/blob/main/Assignments/P02/imagemod/mustangs_blank.png) | Blank image that's created from the first iteration of the program before a new constructor was added. 


### Instructions

https://github.com/rugbyprof/4143-PLC/tree/main/Assignments/P02

### Example Command

An example command for the terminal would be go run main.go, this command runs the program.




