// rubberduck/config.go

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
)

type Config struct {
	Editor                  string   `json:"editor"`
	EditorOpts              string   `json:"editorOpts"`
	TerminalHistoryEnabled  bool     `json:"terminalHistoryEnabled"`
	TerminalHistoryFile     string   `json:"terminalHistoryFile"`
	TerminalHistoryNumLines int      `json:"terminalHistoryNumLines"`
	GoodnightPrompts        []string `json:"goodnightPrompts"`
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
		Editor:                  "nvim",
		EditorOpts:              "+Goyo",
		TerminalHistoryEnabled:  true,
		TerminalHistoryFile:     ".zsh_history",
		TerminalHistoryNumLines: 4,
		GoodnightPrompts: []string{
			"What did you achieve today?",
			"What could you have done better?",
			"How far did you run today?",
			"What's something you learned?",
			"What are you going to accomplish tomorrow?",
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

func config(configPath string) error {
	conf, err := loadConfiguration(configPath)
	if err != nil {
		return err
	}
	cmd := exec.Command(conf.Editor, configPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	return err
}
