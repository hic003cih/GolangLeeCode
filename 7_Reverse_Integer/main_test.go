package main

import "testing"

func TestReverse(t *testing.T) {
	got := reverse(123)

	want := 321

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
