package syschk

import (
	"reflect"
	"testing"

	"sysup-notifier/internal/utils/logger"
)

func init() {
	logger.InitLogger("none", false)
}

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
		{
			name:   "Exact one OSRelease file found",
			distro: All,
			usage:  OSRelease,
			want: []DistroFile{
				{
					distro: All,
					file:   "/etc/os-release",
					usage:  OSRelease,
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
				"PRETTY_NAME":        "Ubuntu 22.04.2 LTS",
				"VERSION_ID":         "22.04",
				"VERSION":            "22.04.2 LTS (Jammy Jellyfish)",
				"SUPPORT_URL":        "https://help.ubuntu.com/",
				"BUG_REPORT_URL":     "https://bugs.launchpad.net/ubuntu/",
				"NAME":               "Ubuntu",
				"VERSION_CODENAME":   "jammy",
				"ID":                 "ubuntu",
				"ID_LIKE":            "debian",
				"HOME_URL":           "https://www.ubuntu.com/",
				"PRIVACY_POLICY_URL": "https://www.ubuntu.com/legal/terms-and-policies/privacy-policy",
				"UBUNTU_CODENAME":    "jammy",
			},
			want1: true,
		},
		{
			name: "Returns no map and boolean 'false'",
			args: args{
				file: "../../testdata/os-release/foo",
			},
			want:  map[string]string{},
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

func Test_determineDietPi(t *testing.T) {
	type args struct {
		df []DistroFile
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isDietPi true",
			args: args{
				[]DistroFile{
					{
						distro: DietPi,
						file:   "../../testdata/dietpi.txt",
						usage:  Identifier,
					},
				},
			},
			want: true,
		},
		{
			name: "isDietPi false",
			args: args{
				[]DistroFile{
					{
						distro: DietPi,
						file:   "../../testdata/foo.txt",
						usage:  Identifier,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := determineDietPi(tt.args.df); got != tt.want {
				t.Errorf("determineDietPi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_determineDistroByOSRelease(t *testing.T) {
	type args struct {
		df []DistroFile
	}
	tests := []struct {
		name       string
		args       args
		wantDistro Distro
	}{
		{
			name: "determine distro by ID successfully",
			args: args{
				[]DistroFile{
					{
						distro: All,
						file:   "../../testdata/os-release/raspbian",
						usage:  OSRelease,
					},
				},
			},
			wantDistro: Raspbian,
		},
		{
			name: "determine distro by ID_LIKE successfully",
			args: args{
				[]DistroFile{
					{
						distro: All,
						file:   "../../testdata/os-release/ubuntu",
						usage:  OSRelease,
					},
				},
			},
			wantDistro: Debian,
		},
		{
			name: "determine failed, return Other",
			args: args{
				[]DistroFile{
					{
						distro: All,
						file:   "../../testdata/os-release/opensuse",
						usage:  OSRelease,
					},
				},
			},
			wantDistro: Other,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistro := determineDistroByOSRelease(tt.args.df); !reflect.DeepEqual(gotDistro, tt.wantDistro) {
				t.Errorf("determineDistroByOSRelease() = %v, want %v", gotDistro, tt.wantDistro)
			}
		})
	}
}

func Test_searchForUpdatesOnDietPi(t *testing.T) {
	type args struct {
		df []DistroFile
	}
	tests := []struct {
		name            string
		args            args
		wantUpdateAvbl  bool
		wantNewVersion  string
		wantCntAptPacks string
	}{
		{
			name: "New DietPi verison and Apt packages available",
			args: args{
				[]DistroFile{
					{
						distro: DietPi,
						file:   "../../testdata/.update_available",
						usage:  Updates,
					},
					{
						distro: DietPi,
						file:   "../../testdata/.apt_updates",
						usage:  Updates,
					},
				},
			},
			wantUpdateAvbl:  true,
			wantNewVersion:  "vFoo.bar",
			wantCntAptPacks: "4",
		},
		{
			name: "New DietPi verison available",
			args: args{
				[]DistroFile{
					{
						distro: DietPi,
						file:   "../../testdata/.update_available",
						usage:  Updates,
					},
					{
						distro: DietPi,
						file:   "../../testdata/.foo",
						usage:  Updates,
					},
				},
			},
			wantUpdateAvbl:  true,
			wantNewVersion:  "vFoo.bar",
			wantCntAptPacks: "",
		},
		{
			name: "Update for Apt packages available",
			args: args{
				[]DistroFile{
					{
						distro: DietPi,
						file:   "../../testdata/.foo",
						usage:  Updates,
					},
					{
						distro: DietPi,
						file:   "../../testdata/.apt_updates",
						usage:  Updates,
					},
				},
			},
			wantUpdateAvbl:  true,
			wantNewVersion:  "",
			wantCntAptPacks: "4",
		},
		{
			name: "No update files found",
			args: args{
				[]DistroFile{
					{
						distro: DietPi,
						file:   "../../testdata/.foo",
						usage:  Updates,
					},
					{
						distro: DietPi,
						file:   "../../testdata/.bar",
						usage:  Updates,
					},
				},
			},
			wantUpdateAvbl:  false,
			wantNewVersion:  "",
			wantCntAptPacks: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUpdateAvbl, gotNewVersion, gotCntAptPacks := searchForUpdatesOnDietPi(tt.args.df)

			if gotUpdateAvbl != tt.wantUpdateAvbl {
				t.Errorf("searchForUpdatesOnDietPi(): update availabe = %v, want %v", gotUpdateAvbl, tt.wantUpdateAvbl)
			}
			if gotNewVersion != tt.wantNewVersion {
				t.Errorf("searchForUpdatesOnDietPi(): new DietPi version = %v, want %v", gotNewVersion, tt.wantNewVersion)
			}
			if gotCntAptPacks != tt.wantCntAptPacks {
				t.Errorf("searchForUpdatesOnDietPi(): counter Apt updates = %v, want %v", gotCntAptPacks, tt.wantCntAptPacks)
			}
		})
	}
}
