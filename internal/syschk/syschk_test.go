package syschk

import (
	"reflect"
	"testing"
)

func Test_getUpdateFiles(t *testing.T) {
	tests := []struct {
		name string
		want []UpdateFiles
	}{
		{
			name: "Update files for Dietpi",
			want: []UpdateFiles{
				{
					os:   "dietpi",
					file: "testdata/.update_available",
				},
				{
					os:   "dietpi",
					file: "testdata/.apt_updates",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUpdateFiles(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUpdateFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
