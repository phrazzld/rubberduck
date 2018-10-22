package main

// rubberduck
// Make quick timestamped notes from the command line

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
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

func pullConfig() (editor string, goyo string) {
	dat, err := ioutil.ReadFile("config")
	check(err)
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

func rubberduck() {
	// Initialize and format current time
	n := time.Now()
	d, t := initDatetime(n)
	// Initialize the note
	f := initFile(n)
	stamp(f, d, t)
	// Read config
	editor, goyo := pullConfig()
	// Launch editor for the note
	cmd := exec.Command(editor, f, goyo)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	check(err)
}

func config() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Editor: ")
	editor, _ := reader.ReadString('\n')
	editor = strings.Replace(editor, "\n", "", -1)
	editorConf := "EDITOR=" + editor
	goyoConf := "GOYO="
	if editor == "vim" {
		goyoConf += "true"
	}
	conf := editorConf + "\n" + goyoConf + "\n"
	f, err := os.Create("config")
	check(err)
	defer f.Close()
	f.WriteString(conf)
	f.Sync()
}

func main() {
	if len(os.Args) == 1 {
		rubberduck()
	} else if os.Args[1] == "config" {
		// Run config UX
		config()
	} else {
		fmt.Println("Unrecognized command")
		os.Exit(1)
	}
}
