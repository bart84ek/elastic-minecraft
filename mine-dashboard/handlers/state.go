package handlers

import (
	"io/ioutil"
	"os"
)

const LOCAL_STATE_FILE = "/tmp/elastic_panel.task"

func setState(state string) error {
	f, err := os.Create(LOCAL_STATE_FILE)
	if err != nil {
		return err
	}
	_, err = f.WriteString(state)
	return err
}

func getState() (bool, string) {
	content, err := ioutil.ReadFile(LOCAL_STATE_FILE)
	if err != nil {
		// file not found. normal thing, justr no status
		return false, ""
	}
	return true, string(content)
}

func rmState() {
	os.Remove(LOCAL_STATE_FILE)
}
