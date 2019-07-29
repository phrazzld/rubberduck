// rubberduck
// Make quick timestamped notes from the command line

package main

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

// Exists takes a filepath string and checks if it exists, returning the appropriate boolean
func Exists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

func initDatetime(t time.Time) (date string, time string) {
	return t.Format("2006 January 2"), t.Format("15:04:05")
}

func createNotesDir(homeDir string) string {
	notesDir := filepath.Join(homeDir, "rubberducks")
	if !Exists(notesDir) {
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
	if !Exists(f) {
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

func pullConfig(confPath string) (editor string, goyo string) {
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

func load(f, confPath string) {
	editor, goyo := pullConfig(confPath)
	// Launch editor for the note
	cmd := exec.Command(editor, f, goyo)
	run(cmd)
}

func search(duckPath string, searchTerm []string) {
	cmd := exec.Command("grep", strings.Join(searchTerm, " "), "-R", duckPath)
	run(cmd)
}

func run(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	check(err)
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
	stamp(f, d, t)
	load(f, confPath)
}

func config(confPath string) {
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
	f, err := os.Create(confPath)
	check(err)
	defer f.Close()
	f.WriteString(conf)
	f.Sync()
}

func review(confPath string) {
	now := time.Now()
	dates := make([]time.Time, 0)
	dates = append(dates, now.AddDate(0, -3, 0))
	dates = append(dates, now.AddDate(0, -1, 0))
	dates = append(dates, now.AddDate(0, 0, -7))
	dates = append(dates, now.AddDate(0, 0, -1))
	for _, date := range dates {
		f := initFile(date)
		if Exists(f) {
			load(f, confPath)
		}
	}
}

func reminisce(confPath string) {
	now := time.Now()
	dates := make([]time.Time, 0)
	for i := -100; i < 0; i++ {
		dates = append(dates, now.AddDate(i, 0, 0))
	}
	dates = append(dates, now.AddDate(0, -6, 0))
	for _, date := range dates {
		f := initFile(date)
		if Exists(f) {
			load(f, confPath)
		}
	}
}

func main() {
	usr, err := user.Current()
	check(err)
	duckPath := filepath.Join(usr.HomeDir, "rubberducks")
	confPath := filepath.Join(duckPath, "config")
	if len(os.Args) == 1 {
		rubberduck(confPath)
	} else {
		switch os.Args[1] {
		case "config":
			config(confPath)
		case "review":
			review(confPath)
		case "reminisce":
			reminisce(confPath)
		case "search":
			search(duckPath, os.Args[2:])
		default:
			fmt.Println("Unrecognized command")
			os.Exit(1)
		}
	}
}
