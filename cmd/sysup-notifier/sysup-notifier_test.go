package main

import (
	"testing"
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
			if got := checkUpdates(); got != tt.want {
				t.Errorf("determineUpdates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_determineUpdateFiles(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := determineUpdateFiles(tt.args.file); got != tt.want {
				t.Errorf("determineUpdateFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
