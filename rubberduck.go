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
	confPath, _ := filepath.Abs("config")
	dat, err := ioutil.ReadFile(confPath)
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

func load(f string) {
	editor, goyo := pullConfig()
	// Launch editor for the note
	cmd := exec.Command(editor, f, goyo)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	check(err)
}

func rubberduck() {
	// Initialize and format current time
	n := time.Now()
	d, t := initDatetime(n)
	// Initialize the note
	f := initFile(n)
	confPath, _ := filepath.Abs("config")
	if !exists(confPath) {
		fmt.Println("No config file found! Run `rubberduck config` to create one.")
		os.Exit(1)
	}
	stamp(f, d, t)
	load(f)
}

func config() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Editor command (vim, nano, open, ...): ")
	editor, _ := reader.ReadString('\n')
	editor = strings.Replace(editor, "\n", "", -1)
	editorConf := "EDITOR=" + editor
	goyoConf := "GOYO="
	if editor == "vim+" {
		editorConf = "EDITOR=vim"
		goyoConf += "true"
	}
	conf := editorConf + "\n" + goyoConf + "\n"
	confPath, _ := filepath.Abs("config")
	f, err := os.Create(confPath)
	check(err)
	defer f.Close()
	f.WriteString(conf)
	f.Sync()
}

func review() {
	now := time.Now()
	dates := make([]time.Time, 0)
	for i := -100; i < 0; i++ {
		dates = append(dates, now.AddDate(i, 0, 0))
	}
	dates = append(dates, now.AddDate(0, -6, 0))
	dates = append(dates, now.AddDate(0, -3, 0))
	dates = append(dates, now.AddDate(0, -1, 0))
	dates = append(dates, now.AddDate(0, 0, -7))
	dates = append(dates, now.AddDate(0, 0, -1))

	for _, date := range dates {
		f := initFile(date)
		if exists(f) {
			load(f)
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		rubberduck()
	} else {
		switch os.Args[1] {
		case "config":
			config()
		case "review":
			review()
		default:
			fmt.Println("Unrecognized command")
			os.Exit(1)
		}
	}
}
