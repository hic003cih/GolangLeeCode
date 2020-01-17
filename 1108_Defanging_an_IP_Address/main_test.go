package main

import "testing"

func TestDefangIPaddr(t *testing.T) {
	got := defangIPaddr("1.1.1.1")

	want := "sasc"

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
