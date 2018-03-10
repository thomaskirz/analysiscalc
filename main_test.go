package main

import "testing"

func TestStoreAndLoadWithKeyword(t *testing.T) {
	evaluate("STORE f(x) = 3x^2 +4")
	str := evaluate("LOAD f(x)")
	if str != "f(x) = 3x^2 + 4" {
		t.Errorf("Expected f(x) = 3x^2 + 4, found %v", str)
	}
}

func TestStoreAndLoadWithoutKeyword(t *testing.T) {
	evaluate("f(x) = 3x^2 +4")
	str := evaluate("f(x)")
	if str != "f(x) = 3x^2 + 4" {
		t.Errorf("Expected f(x) = 3x^2 + 4, found %v", str)
	}
}

func TestDerive(t *testing.T) {
	evaluate("test(x) = 552x^55 + 38x^4 + 6x^2 + 33")
	str := evaluate("derive test(x)")
	if str != "test'(x) = 30360x^54 + 152x^3 + 12x" {
		t.Errorf("Expected test'(x) = 30360x^54 + 152x^3 + 12x, found %v", str)
	}
}

func TestSimpleZeroes(t *testing.T) {
	evaluate("g(x) = 4x^4")
	str := evaluate("zeroes g(x)")
	if str != "0.000\t" && str != "-0.000\t" {
		t.Errorf("Expected 0.000, found %v", str)
	}
}