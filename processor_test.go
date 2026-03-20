package main

import "testing"

func TestBaseConversion(t *testing.T) {
	input := "simply add 1E (hex) files"
	output := "simply add 30 files"
	expected := baseConversion(input)
	if expected != output {
		t.Errorf("Expected %s, got %s", output, expected)
	}
	input1 := "simply add 10 (bin) files"
	output1 := "simply add 2 files"
	expected1 := baseConversion(input1)
	if expected1 != output1 {
		t.Errorf("Expected1 %s, got %s", output1, expected1)
	}
}

func TestApplyCaseTransform(t *testing.T) {
	input := "it (cap) was (up) FUN AND (low, 2)"
output := "It WAS fun and"
	expected := applyCaseTransform(input)
	if output != expected {
		t.Errorf("Expected %s, got %s", output, expected)
	}
}
