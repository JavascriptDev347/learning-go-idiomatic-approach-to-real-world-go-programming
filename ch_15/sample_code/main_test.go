package main

import "testing"

func Test_addNumber(t *testing.T) {
	result := addNumbers(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}

}
