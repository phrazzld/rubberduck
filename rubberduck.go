package main

// rubberduck
// Make quick timestamped notes from the command line

import (
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func exists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func initDatetime(t time.Time) (date string, time string) {
	return t.Format("2006 January 2"), t.Format("15:04:05")
}

func createNotesDir(homeDir string) string {
	notesDir := filepath.Join(homeDir, "rubberducks")
	if !exists(notesDir) {
		err := os.MkdirAll(notesDir, 0755)
		check(err)
	}
	return notesDir
}

func initFile(t time.Time) string {
	usr, err := user.Current()
	check(err)
	dir := createNotesDir(usr.HomeDir)
	file := t.Format("20060102") + ".md"
	return filepath.Join(dir, file)
}

func stamp(f string, d string, t string) {
	// Make (date)timestamp string
	var stamp string
	if !exists(f) {
		stamp = "# " + d
	}
	stamp += "\n\n## " + t + "\n"
	// Open file for writing
	file, err := os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer file.Close()
	// Stamp it
	_, err = file.WriteString(stamp)
	check(err)
	file.Sync()
}

func rubberduck(editor string, file string, goyo string) {
	cmd := exec.Command(editor, file, goyo)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	check(err)
}

func main() {
	n := time.Now()
	d, t := initDatetime(n)

	f := initFile(n)
	stamp(f, d, t)

	// Launch $TERM running $EDITOR for file
	rubberduck("vim", f, "+Goyo")
}
