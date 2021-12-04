package messages

import (
	"testing"
)

func TestGreet(t *testing.T) {
	got := Greet("Hello, Gopher!")
	expected := "Hello, Gopher!\n"

	if got != expected {
		t.Errorf("Did NOT get the expected result, Got '%q' and wanted '%q'", got , expected)

	}
}

func TestDepart(t *testing.T) {
	got := depart("Gopher be good")
	expected := "Gopher be good\n"

	if got != expected {
		t.Errorf("Did NOT get the expected result, Got '%q' and wanted '%q'", got , expected)
		
	}
}

func TestGlobal(t *testing.T) {

}