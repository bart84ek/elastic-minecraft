package handlers

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

func mineCmd() string {
	cmd := os.Getenv("MINE_CMD")
	if len(cmd) == 0 {
		return "mine"
	}
	return cmd
}

func listMines() ([]byte, error) {
	return exec.Command(mineCmd(), "list").CombinedOutput()
}

func stopMine() {
	setState("stopping")
	out, err := exec.Command(mineCmd(), "stop").CombinedOutput()
	if err != nil {
		log.Printf("Error stopping mine: %v\n%s", err, out)
		rmState()
		return
	}
	rmState()
	log.Println("Mine stopped")
}

func startMine() {
	setState("starting")
	out, err := exec.Command(mineCmd(), "start").CombinedOutput()
	if err != nil {
		log.Printf("Error starting mine: %v\n%s", err, out)
		rmState()
		return
	}
	rmState()
	log.Println("Mine started")
}

func mineStatus() (string, error) {
	if found, localState := getState(); found {
		return localState, nil
	}

	out, err := exec.Command(mineCmd(), "list").CombinedOutput()
	if err != nil {
		return "", err
	}

	minesCnt, err := minesCount(out)
	if err != nil {
		return "", err
	}

	if minesCnt > 0 {
		return "running", nil
	}

	return "stopped", nil
}

func minesCount(listOutput []byte) (int, error) {
	re := regexp.MustCompile(`Found (\d+) mine\(s\)`)
	matches := re.FindAllSubmatch([]byte(listOutput), -1)
	if len(matches) == 0 || len(matches[0]) <= 1 {
		return 0, nil
	}
	minesCnt, err := strconv.Atoi(string(matches[0][1]))
	if err != nil {
		return 0, err
	}
	return minesCnt, nil
}
