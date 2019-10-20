// rubberduck/goodmorning.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func getPrompts(configPath string) ([]string, error) {
    var prompts []string
    conf, err := loadConfiguration(configPath)
    if err != nil {
        return prompts, err
    }
    prompts = conf.GoodnightPrompts
	return prompts, err
}

func goodnight(configPath string) error {
	reader := bufio.NewReader(os.Stdin)
	goodnightContent := "### Good Night ###\n\n"

	// Q&A for mood and focus
	fmt.Println("Good night!")
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second / 2)
		fmt.Print(".")
	}
	fmt.Println("But before you hit the hay, a quick touch base?")
	time.Sleep(time.Second / 2)

    prompts, err := getPrompts(configPath)
    if err != nil {
        return err
    }
	for _, prompt := range prompts {
		fmt.Println("\n" + prompt)
		answer, err := reader.ReadString('\n')
		if err != nil {
            return err
		}
		goodnightContent += "> " + prompt + "\n" + answer + "\n"
	}

	// Stamp day's rubberduck
	n := time.Now()
	d, t := initDatetime(n)
	f := initFile(n)
	stamp(f, d, t, configPath)

	// Append goodnight Q&A
	file, err := os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
        return err
	}
	defer file.Close()
	_, err = file.WriteString(goodnightContent)
	if err != nil {
        return err
	}
	file.Sync()
    return err
}
