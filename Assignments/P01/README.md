## P01 - Run a Go Program
### Brayden Lawson
### Description:

This program shows how to import packages, write functions, and use a test file 
in Golang.

### Files

|   #   | File     | Description                      |
| :---: | -------- | -------------------------------- |
|   1   | [main.go](https://github.com/bglawson1001/4143-PLC-Lawson/blob/main/Assignments/P01/main.go) | Main driver of this Golang program. It imports some very important packages not included by default, and returns the name of the mascot in the mascot.go file. |
|   2   | [mascot](https://github.com/bglawson1001/4143-PLC-Lawson/tree/main/Assignments/P01/mascot) | Folder that contains the mascot.go file and the mascot test file.
|   3   | [go.mod](https://github.com/bglawson1001/4143-PLC-Lawson/blob/main/Assignments/P01/go.mod) | This file is necessary because it declares the program's dependencies and their version number. 
|   4   | [go.sum](https://github.com/bglawson1001/4143-PLC-Lawson/blob/main/Assignments/P01/go.sum) | This file maintains the checksum so when the program runs again it will not reinstall the packages, and instead use the cache. 


### Instructions

This video provides all of the instructions. (https://www.youtube.com/watch?v=1MXIGYrMk80)

### Example Command

An example command for the terminal would be go run main.go, this command runs the program.

