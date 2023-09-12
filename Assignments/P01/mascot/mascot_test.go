package mascot_test

import (
	"testing"

	"example.com/p01/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {

		t.Fatal("Wrong mascot :(") // prints to the screen if the name of the mascot is wrong.
	}

}
