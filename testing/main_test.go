package main_test

import "testing"

func TestAddition(t *testing.T) {
	got := 2 + 2
	expected := 4
	if got != expected {
		t.Errorf("Did not get expected Addition result. Got '%v', and wanted '%v' ", got, expected)

	}
}

func TestSubstraction(t *testing.T){
	got := 10 - 5
	expected := 3
	if got != expected {
		t.Errorf("Did NOT get the expected Subtract result Got '%v' and wanted '%v'", got, expected)

	}
}

func TestMultiplication(t *testing.T) {
	got := 100 * 10
	expected := 1000
	if got != expected {
		t.Errorf("Did not get the expected Multiplication result Got '%v' and wanted '%v'", got, expected)
	}
}