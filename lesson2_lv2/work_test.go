package main

import "testing"

func TestCalculatorAdd(t *testing.T) {
	add := Calculator("add")
	if result := add(4, 5); result != 9 {
		t.Errorf("Add expected 9, got %d", result)
	}
}

func TestCalculatorSubtract(t *testing.T) {
	subtract := Calculator("subtract")
	if result := subtract(4, 5); result != -1 {
		t.Errorf("Subtract expected -1, got %d", result)
	}
}

func TestCalculatorMultiply(t *testing.T) {
	multiply := Calculator("multiply")
	if result := multiply(4, 5); result != 20 {
		t.Errorf("Multiply expected 20, got %d", result)
	}
}

func TestCalculatorDivide(t *testing.T) {
	divide := Calculator("divide")
	if result := divide(4, 5); result != 0 {
		t.Errorf("Divide expected 0, got %d", result)
	}
}
