// rubberduck/main.go
// Ad-hoc journaling at the command line.

package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	notesPath := filepath.Join(usr.HomeDir, "rubberducks")
	configPath := filepath.Join(notesPath, "config.json")
	if len(os.Args) == 1 {
		err = rubberduck(configPath)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		switch os.Args[1] {
		case "config":
			err = config(configPath)
			if err != nil {
				fmt.Println(err)
			}
		case "review":
			review(configPath)
		case "reminisce":
			reminisce(configPath)
		case "search":
			hits := search(notesPath, os.Args[2:])
			formatSearchResults(hits)
		case "goodnight":
			err = goodnight(configPath)
			if err != nil {
				fmt.Println(err)
			}
		default:
			fmt.Println("Unrecognized command")
			os.Exit(1)
		}
	}
}
