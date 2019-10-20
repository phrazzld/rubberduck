// rubberduck/config.go

package main

import (
	"bufio"
    "io/ioutil"
	"encoding/json"
	"fmt"
	"os"
	"strings"
    "strconv"
)

type Config struct {
	Editor          string `json:"editor"`
    TerminalHistoryEnabled bool `json:"terminalHistoryEnabled"`
    TerminalHistoryFile string `json:"terminalHistoryFile"`
    TerminalHistoryNumLines int `json:"terminalHistoryNumLines"`
	GoodnightPrompts []string `json:"goodnightPrompts"`
}

func (conf *Config) SetEditor(editor string, configPath string) error {
    conf.Editor = editor
    err := writeConfigObjectToFile(*conf, configPath)
    return err
}

func (conf *Config) SetTerminalHistoryEnabled(enabled bool, configPath string) error {
    conf.TerminalHistoryEnabled = enabled
    err := writeConfigObjectToFile(*conf, configPath)
    return err
}

func (conf *Config) SetTerminalHistoryFile(historyFile string, configPath string) error {
    conf.TerminalHistoryFile = historyFile
    err := writeConfigObjectToFile(*conf, configPath)
    return err
}

func (conf *Config) SetTerminalHistoryNumLines(numLines int, configPath string) error {
    conf.TerminalHistoryNumLines = numLines
    err := writeConfigObjectToFile(*conf, configPath)
    return err
}

func (conf *Config) SetGoodnightPrompts(questions []string, configPath string) error {
    conf.GoodnightPrompts = questions
    err := writeConfigObjectToFile(*conf, configPath)
    return err
}

func writeConfigObjectToFile(conf Config, configPath string) error {
    jsonConf, err := json.MarshalIndent(conf, "", "  ")
    if err != nil {
        return err
    }
    err = ioutil.WriteFile(configPath, jsonConf, 0644)
    return err
}

func setDefaultConfiguration(configPath string) error {
    conf := Config{
        Editor: "vim",
        TerminalHistoryEnabled: false,
        TerminalHistoryFile: "",
        TerminalHistoryNumLines: 0,
        GoodnightPrompts: []string{
            "What did you achieve today?",
            "What could you have done better?",
            "What are you going to do tomorrow?",
        },
    }
    err := writeConfigObjectToFile(conf, configPath)
    return err
}

func loadConfiguration(configPath string) (Config, error) {
	var conf Config
    if !Exists(configPath) {
        err := setDefaultConfiguration(configPath)
        if err != nil {
            return conf, err
        }
    }
	configFile, err := os.Open(configPath)
	defer configFile.Close()
	if err != nil {
		return conf, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&conf)
	return conf, err
}

func configEditor(configPath string, reader *bufio.Reader, conf Config) (Config, error) {
	fmt.Print("Editor command (vim, nano, open, ...): ")
	editor, err := reader.ReadString('\n')
	if err != nil {
        return conf, err
	}
    editor = stripNewlines(editor)
    conf.SetEditor(editor, configPath)
    err = writeConfigObjectToFile(conf, configPath)
    return conf, err
}

func stripNewlines(s string) string {
    return strings.ToLower(strings.Replace(s, "\n", "", -1))
}

func configTerminalHistory(configPath string, reader *bufio.Reader, conf Config) (Config, error) {
    fmt.Print("Enable terminal history in journal entries? (y/n): ")
    enableHistory, err := reader.ReadString('\n')
    if err != nil {
        return conf, err
    }
    enableHistory = stripNewlines(enableHistory)
    switch enableHistory {
    case "n":
        conf.SetTerminalHistoryEnabled(false, configPath)
        conf.SetTerminalHistoryFile("", configPath)
        conf.SetTerminalHistoryNumLines(0, configPath)
    case "y":
        conf.SetTerminalHistoryEnabled(true, configPath)
        fmt.Print("What is the name of your history file? (.bash_history, .zsh_history, ...): ")
        historyFile, err := reader.ReadString('\n')
        if err != nil {
            return conf, err
        }
        historyFile = stripNewlines(historyFile)
        conf.SetTerminalHistoryFile(historyFile, configPath)
        fmt.Print("Lines of history to include: ")
        numLinesStr, err := reader.ReadString('\n')
        if err != nil {
            return conf, err
        }
        numLinesStr = stripNewlines(numLinesStr)
        numLines, err := strconv.Atoi(numLinesStr)
        if err != nil {
            return conf, err
        }
        conf.SetTerminalHistoryNumLines(numLines, configPath)
    default:
		fmt.Println("y/n: It means \"Enter the character 'y' or the character 'n'\"")
    }
    err = writeConfigObjectToFile(conf, configPath)
    return conf, err
}

func config(configPath string) error {
	reader := bufio.NewReader(os.Stdin)
    conf, err := loadConfiguration(configPath)
    if err != nil {
        return err
    }
	// Configure editor
	conf, err = configEditor(configPath, reader, conf)
    if err != nil {
        return err
    }
	// Toggle terminal history in stamps
	conf, err = configTerminalHistory(configPath, reader, conf)
    if err != nil {
        return err
    }
    // Save config to file
    err = writeConfigObjectToFile(conf, configPath)
    return err
}
