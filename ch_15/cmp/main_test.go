package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreatePerson(t *testing.T) {

	expected := Person{
		Name:  "Russel",
		Email: "russel@gmail.com",
	}
	result := NewPerson("Russel", "russel@gmail.com")
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}
