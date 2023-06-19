package pub

import (
	"strings"
	"testing"

	"sysup-notifier/internal/utils/logger"
)

func init() {
	logger.InitLogger("none", false)
}

func Test_generateMsg(t *testing.T) {
	type args struct {
		newVersion  string
		cntAptPacks string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "New version and Apt packages available",
			args: args{
				newVersion: "vFoo.bar",
				cntAptPacks: "4",
			},
			want: "System updates available. New version: vFoo.bar, APT package updates: 4",
		},
		{
			name: "New version available",
			args: args{
				newVersion: "vFoo.bar",
				cntAptPacks: "",
			},
			want: "System updates available. New version: vFoo.bar,",
		},
		{
			name: "Apt packages available",
			args: args{
				newVersion: "",
				cntAptPacks: "4",
			},
			want: "System updates available. APT package updates: 4",
		},
		{
			name: "No update infos available",
			args: args{
				newVersion: "",
				cntAptPacks: "",
			},
			want: "System updates available.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateMsg(tt.args.newVersion, tt.args.cntAptPacks); 
			f := strings.Contains(got, tt.want)
			if !f {
				t.Errorf("generateMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
