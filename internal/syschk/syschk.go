package syschk

import (
	"context"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

	log.Printf("Distro %v", d.toString())

	updatesAvbl := false

	switch d {
	case Other:
		log.Printf("OS not defined: %v", d)
	case DietPi:
		updatesAvbl = SearchForUpdatesOnDietpi()
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
// Returns boolean value "true" if updates available, otherwise "false".
func SearchForUpdatesOnDietpi() bool {
	for _, f := range getDistroFiles(DietPi, Updates) {
		fileExists := determineFile(f.file)
		if fileExists {
			log.Println("Update file found:", f.file)
			return true
		}
	}
	log.Println("No update files found")
	return false
}

// Checks whether updates are available with APT package manager.
// Returns boolean value "true" if updates available, otherwise "false".
func SearchForUpdatesWithApt() bool {
	// Set timeout 90 sec for apt-get update
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	err := exec.CommandContext(ctx, "sudo", "apt-get", "-y", "-qq", "update").Run()
	if err != nil {
		log.Printf("Error on 'apt-get update': %s", err)
	}

	cmd := "sudo apt list -qq --upgradable 2> /dev/null | wc -l"
	c, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Printf("Error on '%s': %s", cmd, err)
	}

	if int(c[0]) != 0 {
		log.Printf("%s updates available.", c)

		return true
	}

	return false
}

// Returns "true" on Dietpi OS, otherwise false
func determineDietPi(df []DistroFile) bool {
	for _, f := range df {
		fileExists := determineFile(f.file)
		if fileExists {
			log.Println("Identifier file found:", f.file)
			return true
		}
	}

	return false
}

// Read /etc/os-release, determine distro and return Distro value
func determineDistroByOSRelease(df []DistroFile) (distro Distro) {
	var osr string = df[0].file
	log.Println(osr)

	osrmap, osrMapExists := readOSRelease(osr)

	if osrMapExists {
		osrID := cases.Title(language.Und).String(osrmap["ID"])
		log.Println(osrID)

		osrIDL := cases.Title(language.Und).String(osrmap["ID_LIKE"])
		log.Println(osrIDL)

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
		log.Println("/etc/os-release not exists.")
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

	log.Println(osmap)
	return osmap, true
}

// Check whether update file is available.
// Returns "true", if input file found, otherwise "false".
func determineFile(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		log.Printf("File '%s' not found", file)
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
