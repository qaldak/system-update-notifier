package sysos

import (
	"log"
	"os"
)

func CheckForUpdates() bool {
	// check update files on Dietpi
	updateFiles := []string{"testdata/.update_available", "testdata/.apt_updates"}

	for _, file := range updateFiles {
		fileExists := determineUpdateFiles(file)
		if fileExists {
			log.Println("Update file found:", file)
			return true
		}
	}
	log.Println("No update files found")
	return false
}

func determineUpdateFiles(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		log.Println("Update file not found:", file)
		return false
	}
	return true
}