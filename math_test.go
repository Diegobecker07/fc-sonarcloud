package main

import "testing"

func TestSoma(t *testing.T) {
	if Soma(2, 2) != 4 {
		t.Error("Expected 4")
	}
}
