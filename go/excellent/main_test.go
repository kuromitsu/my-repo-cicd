package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	if EvenOrOdd(2) != "even" {
		t.Error("Expected 2 to be even")
	}
	if EvenOrOdd(3) != "odd" {
		t.Error("Expected 3 to be odd")
	}
}
