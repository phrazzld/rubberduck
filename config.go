// rubberduck/config.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getEditorConfig(reader *bufio.Reader) string {
	fmt.Print("Editor command (vim, nano, open, ...): ")
	editor, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	editor = strings.Replace(editor, "\n", "", -1)
	editorConfig := "EDITOR=" + editor
	// Configure Goyo (secretly)
	if editor == "vim+" {
		editorConfig = "EDITOR=vim\nGOYO=true"
	}
	return editorConfig
}

func getHistoryLinesConfig(reader *bufio.Reader) string {
	historyLinesConfig := "HISTORY_LINES="
	// Get amount of history to include by default
	fmt.Print("Number of lines of history to include: ")
	numLines, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	historyLinesConfig += strings.Replace(numLines, "\n", "", -1)
	return historyLinesConfig
}

func getHistoryFileConfig(reader *bufio.Reader) string {
	historyFileConfig := "HISTORY_FILE="
	// Get history file name
	fmt.Print("Name of terminal history file (.bash_history, .zsh_history, ...): ")
	historyFile, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	historyFileConfig += strings.Replace(historyFile, "\n", "", -1)
	return historyFileConfig
}

func disableHistoryConfig() string {
	historyConfig := "HISTORY_FILE=NONE"
	historyConfig += "\nHISTORY_LINES=0"
	return historyConfig
}

func enableHistoryConfig(reader *bufio.Reader) string {
	historyConfig := getHistoryFileConfig(reader) + "\n"
	historyConfig += getHistoryLinesConfig(reader) + "\n"
	return historyConfig
}

func getShellHistoryConfig(reader *bufio.Reader) string {
	historyConfig := ""
	fmt.Print("Include terminal history in journal entries? (y/n): ")
	includeHistory, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	includeHistory = strings.ToLower(strings.Replace(includeHistory, "\n", "", -1))
	switch includeHistory {
	case "n":
		historyConfig += disableHistoryConfig()
	case "y":
		historyConfig += enableHistoryConfig(reader)
	default:
		fmt.Println("y/n: It means \"Enter the character 'y' or the character 'n'\"")
	}
	return historyConfig
}

func config(configPath string) {
	reader := bufio.NewReader(os.Stdin)
	// Configure editor
	editorConfig := getEditorConfig(reader)
	// Toggle terminal history in stamps
	historyConfig := getShellHistoryConfig(reader)
	config := editorConfig + "\n" + historyConfig + "\n"
	f, err := os.Create(configPath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	f.WriteString(config)
	f.Sync()
}
