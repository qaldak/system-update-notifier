package main

import (
	"testing"
)

func TestFoo(t *testing.T) {
	got := getFoo()
	want := "Foo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func Test_determineUpdates(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := determineUpdates(); got != tt.want {
				t.Errorf("determineUpdates() = %v, want %v", got, tt.want)
			}
		})
	}
}
