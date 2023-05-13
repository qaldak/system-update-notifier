package syschk

import (
	"reflect"
	"testing"
)

func Test_getDistroFiles(t *testing.T) {
	tests := []struct {
		name   string
		distro Distro
		usage  FileUsage
		want   []DistroFile
	}{
		{
			name:   "Update files for Dietpi",
			distro: DietPi,
			usage:  Updates,
			want: []DistroFile{
				{
					distro: DietPi,
					file:   "/run/dietpi/.update_available",
					usage:  Updates,
				},
				{
					distro: DietPi,
					file:   "/run/dietpi/.apt_updates",
					usage:  Updates,
				},
			},
		},
		{
			name:   "Identifier file for Dietpi",
			distro: DietPi,
			usage:  Identifier,
			want: []DistroFile{
				{
					distro: DietPi,
					file:   "/boot/dietpi.txt",
					usage:  Identifier,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDistroFiles(tt.distro, tt.usage); !reflect.DeepEqual(got, tt.want) {
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

func Test_readOSRelease(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]string
		want1 bool
	}{
		{
			name: "Map file content and return map and boolean 'true'",
			args: args{
				file: "../../testdata/os-release/ubuntu",
			},
			want: map[string]string{
				"PRETTY_NAME":"Ubuntu 22.04.2 LTS",
				"VERSION_ID":"22.04",
				"VERSION":"22.04.2 LTS (Jammy Jellyfish)",
				"SUPPORT_URL":"https://help.ubuntu.com/",
				"BUG_REPORT_URL":"https://bugs.launchpad.net/ubuntu/",
				"NAME":"Ubuntu",
				"VERSION_CODENAME":"jammy",
				"ID":"ubuntu",
				"ID_LIKE":"debian",
				"HOME_URL":"https://www.ubuntu.com/",
				"PRIVACY_POLICY_URL":"https://www.ubuntu.com/legal/terms-and-policies/privacy-policy",
				"UBUNTU_CODENAME":"jammy",
			},
			want1: true,
		},
		{
			name: "Returns no map and boolean 'false'",
			args: args{
				file: "../../testdata/os-release/foo",
			},
			want: map[string]string{

			},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := readOSRelease(tt.args.file)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readOSRelease() got file content (map) = %v, want file content (map) = %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("readOSRelease() got = %v, want = %v", got1, tt.want1)
			}
		})
	}
}

func Test_determineDistro(t *testing.T) {
	tests := []struct {
		name string
		want Distro
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := determineDistro(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("determineDistro() = %v, want %v", got, tt.want)
			}
		})
	}
}
