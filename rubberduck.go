// rubberduck/rubberduck.go

package main

import (
	"fmt"
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

func stamp(f, d, t, configPath string) error {
	conf, err := loadConfiguration(configPath)
	if err != nil {
		return err
	}
	numLines := conf.TerminalHistoryNumLines
	historyFile := conf.TerminalHistoryFile
	// Make (date)timestamp string
	var stamp string
	if !Exists(f) {
		stamp = "# " + d
	}
	stamp += "\n\n## " + t
	switch numLines {
	// When set to 0, finish with the timestamp
	case 0:
		stamp += "\n\n\n"
	default:
		// When anything else, append terminal history to the stamp
		// Make terminal history string
		stamp += "\n```"
		// Why not handle negative numbers under the hood when it's so inexpensive to do so?
		if numLines < 0 {
			numLines *= -1
		}
		// Add each event to the stamp
		for _, event := range getTerminalHistory(numLines, historyFile) {
			stamp += "\n" + event
		}
		stamp += "\n```\n\n\n"
	}
	// Open file for writing
	file, err := os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	// Stamp it
	_, err = file.WriteString(stamp)
	if err != nil {
		return err
	}
	file.Sync()
	return err
}

// getTerminalHistory returns the output from `history` as a slice of strings
func getTerminalHistory(n int, historyFile string) []string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	var output strings.Builder
	// Run the "history" bash command
	cmd := exec.Command("cat", filepath.Join(usr.HomeDir, historyFile))
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
	// Only show the last few lines of history
	x := min(len(history), n+1)
	return history[len(history)-x : len(history)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func load(f, configPath string) error {
	conf, err := loadConfiguration(configPath)
	if err != nil {
		return err
	}
	editor := conf.Editor
	// Launch editor for the note
	cmd := exec.Command(editor, f)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	return err
}

func rubberduck(configPath string) error {
	// Initialize and format current time
	n := time.Now()
	d, t := initDatetime(n)
	// Initialize the note
	f := initFile(n)
	// Stamp with timestamp and terminal history
	err := stamp(f, d, t, configPath)
	if err != nil {
		return err
	}
	// Load the note
	err = load(f, configPath)
	if err != nil {
		return err
	}
	return err
}
