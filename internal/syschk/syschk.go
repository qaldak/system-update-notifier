package syschk

import (
	"context"
	"os"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"sysup-notifier/internal/utils/logger"
)

type Distro int

const (
	All    Distro = iota // for all Linux distros
	Debian               // Debian bases OS by ID_Like
	DietPi
	Raspbian
	Other // Undefined distro/OS
)

func (d Distro) toString() string {
	return []string{"All", "DietPi", "Raspbian", "Debian", "Other"}[d]
}

type FileUsage int

const (
	Updates FileUsage = iota
	Identifier
	OSRelease
)

type DistroFile struct {
	distro Distro
	file   string
	usage  FileUsage
}

// Checks for available updates on OS.
// Returns "true" if updates available, otherwise "false".
func SearchForUpdates() bool {
	var d Distro

	isDietPi := determineDietPi(getDistroFiles(DietPi, Identifier))
	if isDietPi {
		d = DietPi
	} else {
		d = determineDistroByOSRelease(getDistroFiles(All, OSRelease))
	}

	logger.Debug("Distro %v", d.toString())

	updatesAvbl := false
	newVersion := ""
	cntAptPacks := ""

	switch d {
	case Other:
		logger.Warn("Distro not defined: %v", d.toString())
	case DietPi:
		updatesAvbl, newVersion, cntAptPacks = searchForUpdatesOnDietPi(getDistroFiles(DietPi, Updates))
		logger.Debug("DietPi updates available: %v, new DietPi version: %v, APT packages updates: %v", updatesAvbl, newVersion, cntAptPacks)
	default:
		updatesAvbl = SearchForUpdatesWithApt()
	}

	if updatesAvbl {
		return true
	} else {
		return false
	}
}

// Checks whether updates are available for Dietpi operating system.
// If updates available return boolean "true" incl optional additional update informations (newVersion, cntAptPacks)
// If no updates available return boolean "false"
func searchForUpdatesOnDietPi(df []DistroFile) (updatesAvbl bool, newVersion string, cntAptPacks string) {
	updatesAvbl = false
	newVersion = ""
	cntAptPacks = ""

	for _, f := range df {
		fileExists := determineFile(f.file)
		if fileExists {
			logger.Debug("Update file found: %v.", f.file)

			fc, err := os.ReadFile(f.file)
			if err != nil {
				logger.Warn("Error while read file: %v, Error: %v", f.file, err)
			}

			if strings.Contains(f.file, ".update_available") {
				newVersion = string(fc)
				logger.Debug("new DietPi version available: %v.", newVersion)
			}

			if strings.Contains(f.file, ".apt_updates") {
				cntAptPacks = string(fc)
				logger.Debug("Updates for APT packages available: %v.", cntAptPacks)
			}

			updatesAvbl = true
		}
	}

	if !updatesAvbl {
		logger.Debug("No update files found on DietPi.")
	}

	return
}

// Checks whether updates are available with APT package manager.
// Returns boolean value "true" if updates available, otherwise "false".
func SearchForUpdatesWithApt() bool {
	// Set timeout 90 sec for apt-get update
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	err := exec.CommandContext(ctx, "sudo", "apt-get", "-y", "-qq", "update").Run()
	if err != nil {
		logger.Error("Error on 'apt-get update'. %v.", err)
	}

	cmd := "sudo apt list -qq --upgradable 2> /dev/null | wc -l"
	c, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		logger.Error("Error on '%v': %v.", cmd, err)
	}

	if int(c[0]) != 0 {
		logger.Debug("%v updates available.", c)
		return true
	}

	return false
}

// Returns "true" on Dietpi OS, otherwise false
func determineDietPi(df []DistroFile) bool {
	for _, f := range df {
		fileExists := determineFile(f.file)
		if fileExists {
			logger.Debug("Identifier file found: %v.", f.file)
			return true
		}
	}
	return false
}

// Read /etc/os-release, determine distro and return Distro value
func determineDistroByOSRelease(df []DistroFile) (distro Distro) {
	var osr string = df[0].file

	osrmap, osrMapExists := readOSRelease(osr)

	if osrMapExists {
		osrID := cases.Title(language.Und).String(osrmap["ID"])
		logger.Debug("os-release ID: %v", osrID)

		osrIDL := cases.Title(language.Und).String(osrmap["ID_LIKE"])
		logger.Debug("os-release ID_LIKE: %v", osrIDL)

		switch osrID {
		case "Debian":
			return Debian
		case "Raspbian":
			return Raspbian
		default:
			switch osrIDL {
			case "Debian":
				return Debian
			default:
				return Other
			}
		}
	} else {
		return Other
	}
}

// Read /etc/os-release and create key/value map.
// Returns map and a boolean value, if os-release (map) exists.
func readOSRelease(file string) (map[string]string, bool) {
	osmap := make(map[string]string)
	fc, err := os.ReadFile(file)
	if err != nil {
		logger.Warn("/etc/os-release not exists. %v", err)
		return osmap, false
	}

	s := string(fc)

	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			break
		}

		pair := strings.Split(line, "=")
		k := pair[0]
		v := pair[1]

		v = strings.Trim(v, "\"")
		v = strings.Trim(v, "\r")
		v = strings.Trim(v, "\"")
		osmap[k] = v
	}

	logger.Debug("os-release mapped. %v", osmap)
	return osmap, true
}

// Check whether update file is available.
// Returns "true", if input file found, otherwise "false".
func determineFile(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		logger.Debug("File not found. %v", file)
		return false
	}
	return true
}

// Returns necessary update files for checking.
func getDistroFiles(d Distro, usage FileUsage) []DistroFile {
	res := []DistroFile{}
	file := []DistroFile{
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
		{
			distro: DietPi,
			file:   "/boot/dietpi.txt",
			usage:  Identifier,
		},
		{
			distro: All,
			file:   "/etc/os-release",
			usage:  OSRelease,
		},
	}

	for i := range file {
		if file[i].usage == usage && file[i].distro == d {
			res = append(res, file[i])
		}
	}
	return res
}
