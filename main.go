package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type Command struct {
	Cmd string `json:"cmd"`
}

func cmdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var cmd Command
	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if cmd.Cmd == "" {
		http.Error(w, "Command not found", http.StatusNotFound)
		return
	}

	cmdOutput, err := execCommand(cmd.Cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, cmdOutput)
}

func execCommand(cmd string) (string, error) {
	cmdParts := strings.Fields(cmd)
	if len(cmdParts) == 0 {
		return "", fmt.Errorf("Empty command")
	}

	cmdOut, err := exec.Command(cmdParts[0], cmdParts[1:]...).Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return string(exitErr.Stderr), err
		}
		return "", err
	}

	return string(cmdOut), nil
}

func main() {
	http.HandleFunc("/api/cmd", cmdHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
