package syschk

import (
	"log"
	"os"
	"os/exec"
)

type OS int

const (
	DietPi OS = iota
	Debian    // Debian based OS
	Other
)

type FileUsage int

const (
	Updates FileUsage = iota
	Identifier
)

type OsFiles struct {
	os    OS
	file  string
	usage FileUsage
}

// Checks for available updates on OS.
// Returns "true" if updates available, otherwise "false".
func SearchForUpdates() bool {
	isDp := determineOs()
	updatesAvbl := false

	if isDp == DietPi {
		updatesAvbl = SearchForUpdatesOnDietpi()
	} else {
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
	err := exec.Command("sudo", "apt-get", "-y", "-qq", "update").Run()
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

		return false // Todo: true!
	}

	return false
}

// Returns "true" on Dietpi OS, otherwise false
func determineOs() OS {
	for _, f := range getOsFiles(DietPi, Identifier) {
		fileExists := determineFile(f.file)
		if fileExists {
			log.Println("Identifier file found:", f.file)
			return DietPi
		}
	}

	// Todo: determine OS by /etc/os-release
	return Debian
}

// Check whether update file is available.
// Returns "true", if input file found, otherwise "false".
func determineFile(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		log.Println("Update file not found:", file)
		return false
	}
	return true
}

// Returns necessary update files for checking.
func getOsFiles(os OS, usage FileUsage) []OsFiles {
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
