package testingutil

import (
	"mygoproject/reusable"
)

func RunTests() {
	// Include your unit tests here.
	// Example:
	testAddition()
}

func testAddition() {
	result := reusable.Add(5, 10)
	if result != 15 {
		panic("Addition test failed")
	}
}
