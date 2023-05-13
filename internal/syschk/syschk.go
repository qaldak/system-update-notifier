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
	DietPi Distro = iota
	Raspbian
	Debian // Debian bases OS by ID_Like
	Other  // Undefined distro/OS
)

func (d Distro) toString() string {
	return []string{"DietPi", "Raspbian", "Debian", "Other"}[d]
}

type FileUsage int

const (
	Updates FileUsage = iota
	Identifier
)

type OsFiles struct {
	os    Distro
	file  string
	usage FileUsage
}

// Checks for available updates on OS.
// Returns "true" if updates available, otherwise "false".
func SearchForUpdates() bool {
	d := determineDistro()
	log.Printf("Distro %v", d.toString())

	updatesAvbl := false

	switch d {
	case DietPi:
		updatesAvbl = SearchForUpdatesOnDietpi()
	case Debian:
		updatesAvbl = SearchForUpdatesWithApt()
	default:
		log.Printf("OS not defined: %v", d)
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
	for _, f := range getOsFiles(DietPi, Updates) {
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
func determineDistro() Distro {
	for _, f := range getOsFiles(DietPi, Identifier) {
		fileExists := determineFile(f.file)
		if fileExists {
			log.Println("Identifier file found:", f.file)
			return DietPi
		}
	}

	osrmap, osrMapExists := getOSRelease()

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
func getOSRelease() (map[string]string, bool) {
	osmap := make(map[string]string)
	fc, err := os.ReadFile("/etc/os-release")

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
func getOsFiles(os Distro, usage FileUsage) []OsFiles {
	res := []OsFiles{}
	osFiles := []OsFiles{
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
		{
			os:    DietPi,
			file:  "/boot/dietpi.txt",
			usage: Identifier,
		},
	}

	for i := range osFiles {
		if osFiles[i].usage == usage && osFiles[i].os == os {
			res = append(res, osFiles[i])
		}
	}
	return res
}
