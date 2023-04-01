package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(getFoo())

	hasUpdates := determineUpdates()

	if hasUpdates {
		log.Println("Updates available.")
	}
}

func getFoo() string {
	return "Foo"
}

func determineUpdates() bool {
	_, err := os.Stat("testdata/.update_available")
	if err != nil {
		log.Println("No dietpi updates available.")
		_, err := os.Stat("testdata/.apt_updates")
		if err != nil {
			log.Println("No apt updates available.")
			return false
		}
	}
	return true
}
