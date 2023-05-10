package syschk

import (
	"log"
	"os"
	"os/exec"
)

type UpdateFiles struct {
	os   string
	file string
}

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

// Checks whether updates are available for operating system.
// Returns boolean value "true" if updates available, otherwise "false".
func SearchForUpdates() bool {
	for _, f := range getUpdateFiles() {
		fileExists := determineUpdateFile(f.file)
		if fileExists {
			log.Println("Update file found:", f.file)
			return true
		}
	}
	log.Println("No update files found")
	return false
}

// Returns necessary update files for checking.
func getUpdateFiles() []UpdateFiles {
	return []UpdateFiles{
		{
			os:   "dietpi",
			file: "/run/dietpi/.update_available",
		},
		{
			os:   "dietpi",
			file: "/run/dietpi/.apt_updates",
		},
	}
}

// Check whether update file is available.
// Returns "true", if input file found, otherwise "false".
func determineUpdateFile(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		log.Println("Update file not found:", file)
		return false
	}
	return true
}
