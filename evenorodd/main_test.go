package main

import "testing"

func Test_evenOrOdd(t *testing.T) {
	if evenOrOdd(1) != "odd" {
		t.Errorf("One has to be a odd number, but got %v", evenOrOdd(1))
	}

	if evenOrOdd(2) != "even" {
		t.Errorf("Two has to be a even number, but got %v", evenOrOdd(2))
	}
}
