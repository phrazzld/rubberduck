// rubberduck/rubberduck.go

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

func initDatetime(t time.Time) (date string, time string) {
	return t.Format("2006 January 2"), t.Format("15:04:05")
}

func createNotesDir(homeDir string) string {
	notesDir := filepath.Join(homeDir, "rubberducks")
	if !Exists(notesDir) {
		err := os.MkdirAll(notesDir, 0755)
		if err != nil {
			fmt.Println(err)
		}
	}
	return notesDir
}

func initFile(t time.Time) string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	dir := createNotesDir(usr.HomeDir)
	file := t.Format("20060102") + ".md"
	return filepath.Join(dir, file)
}

func stamp(f string, d string, t string) {
	// Make (date)timestamp string
	var stamp string
	if !Exists(f) {
		stamp = "# " + d
	}
	stamp += "\n\n## " + t
	// Make terminal history string
	stamp += "\n```"
	// Add each event to the stamp
	for _, event := range getTerminalHistory() {
		stamp += "\n" + event
	}
	stamp += "\n```\n\n\n"
	// Open file for writing
	file, err := os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// Stamp it
	_, err = file.WriteString(stamp)
	if err != nil {
		fmt.Println(err)
	}
	file.Sync()
}

// getTerminalHistory returns the output from `history` as a slice of strings
func getTerminalHistory() []string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	var output strings.Builder
	// Run the "history" bash command
	cmd := exec.Command("cat", filepath.Join(usr.HomeDir, ".bash_history"))
	cmd.Stdin = os.Stdin
	// Write the output to our strings.Builder
	cmd.Stdout = &output
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	// Convert output to a string
	// Then split the string on newlines
	history := strings.Split(output.String(), "\n")
	// Only show the last few lines of .bash_history
	// Even that's too noisy, but I want the reminder to do more with the info
	x := min(len(history), 5)
	return history[len(history)-x : len(history)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func pullConfig(confPath string) (editor string, goyo string) {
	dat, err := ioutil.ReadFile(confPath)
	if err != nil {
		fmt.Println(err)
	}
	conf := string(dat)
	configs := strings.Split(strings.Replace(conf, "\n", "=", -1), "=")
	for i, val := range configs {
		switch val {
		case "EDITOR":
			editor = configs[i+1]
		case "GOYO":
			if configs[i+1] == "true" {
				goyo = "+Goyo"
			}
		}
	}
	return editor, goyo
}

func load(f, confPath string) {
	editor, goyo := pullConfig(confPath)
	// Launch editor for the note
	cmd := exec.Command(editor, f, goyo)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func rubberduck(confPath string) {
	// Initialize and format current time
	n := time.Now()
	d, t := initDatetime(n)
	// Initialize the note
	f := initFile(n)
	if !Exists(confPath) {
		fmt.Println("No config file found! Run `rubberduck config` to create one.")
		os.Exit(1)
	}
	// Stamp with timestamp and terminal history
	stamp(f, d, t)
	// Load the note
	load(f, confPath)
}
