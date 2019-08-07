// rubberduck/config.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func config(confPath string) {
	reader := bufio.NewReader(os.Stdin)
	// Configure editor
	fmt.Print("Editor command (vim, nano, open, ...): ")
	editor, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	editor = strings.Replace(editor, "\n", "", -1)
	editorConf := "EDITOR=" + editor
	// Configure Goyo (secretly)
	goyoConf := "GOYO="
	if editor == "vim+" {
		editorConf = "EDITOR=vim"
		goyoConf += "true"
	}
	// Toggle terminal history in stamps
	fmt.Print("Include terminal history in journal entries? (y/n): ")
	includeHistory, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	includeHistory = strings.ToLower(strings.Replace(includeHistory, "\n", "", -1))
	historyConf := "HISTORY="
	switch includeHistory {
	case "n":
		historyConf += "0"
	case "y":
		// Ask how much history to use
		fmt.Print("Number of lines of history to include: ")
		numLines, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		historyConf += strings.Replace(numLines, "\n", "", -1)
	default:
		fmt.Println("y/n: It means \"Enter the character 'y' or the character 'n'\"")
	}
	// Assemble the config file
	conf := editorConf + "\n" + goyoConf + "\n" + historyConf + "\n"
	f, err := os.Create(confPath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	f.WriteString(conf)
	f.Sync()
}
