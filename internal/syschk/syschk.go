package syschk

import (
	"log"
	"os"
)

func SearchForUpdates() bool {
	// check update files on Dietpi
	updateFiles := []string{"testdata/.update_available", "testdata/.apt_updates"}

	for _, file := range updateFiles {
		fileExists := determineUpdateFile(file)
		if fileExists {
			log.Println("Update file found:", file)
			return true
		}
	}
	log.Println("No update files found")
	return false
}

func determineUpdateFile(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		log.Println("Update file not found:", file)
		return false
	}
	return true
}
