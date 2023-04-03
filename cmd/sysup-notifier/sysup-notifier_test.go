package main

import (
	"testing"

	"sysup-notifier/internal/sysos"
)

func Test_determineUpdates(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sysos.CheckForUpdates(); got != tt.want {
				t.Errorf("determineUpdates() = %v, want %v", got, tt.want)
			}
		})
	}
}
