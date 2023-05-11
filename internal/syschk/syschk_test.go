package syschk

import (
	"reflect"
	"testing"
)

func Test_getOsFiles(t *testing.T) {
	tests := []struct {
		name  string
		os    OS
		usage FileUsage
		want  []OsFiles
	}{
		{
			name:  "Update files for Dietpi",
			os:    DietPi,
			usage: Updates,
			want: []OsFiles{
				{
					os:    DietPi,
					file:  "/run/dietpi/.update_available",
					usage: Updates,
				},
				{
					os:    DietPi,
					file:  "/run/dietpi/.apt_updates",
					usage: Updates,
				},
			},
		},
		{
			name:  "Identifier file for Dietpi",
			os:    DietPi,
			usage: Identifier,
			want: []OsFiles{
				{
					os:    DietPi,
					file:  "/boot/dietpi.txt",
					usage: Identifier,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOsFiles(tt.os, tt.usage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getOsFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_determineFile(t *testing.T) {
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
		{
			name: "Identifier file found",
			file: "../../testdata/dietpi.txt",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := determineFile(tt.file); got != tt.want {
				t.Errorf("determineFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
