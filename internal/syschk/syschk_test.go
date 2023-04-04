package syschk

import (
	"testing"
)

func Test_determineUpdateFile(t *testing.T) {
	tests := []struct {
		name string
		file string
		want bool
	}{
		{
			name: "Update file found",
			file: "../../testdata/.update_available",
			want: true,
		},
		{
			name: "Update file not found",
			file: "../../testdata/.foo",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := determineUpdateFile(tt.file); got != tt.want {
				t.Errorf("determineUpdateFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
