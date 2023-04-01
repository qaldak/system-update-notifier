package main

import "testing"

func TestFoo(t *testing.T) {
	got := getFoo()
	want := "Foo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
