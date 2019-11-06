// rubberduck/main.go
// Ad-hoc journaling at the command line.

package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	notesPath := filepath.Join(usr.HomeDir, "rubberducks")
	configPath := filepath.Join(notesPath, "config.json")
	if len(os.Args) == 1 {
		err = rubberduck(configPath)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		switch os.Args[1] {
		case "config":
			err = config(configPath)
			if err != nil {
				log.Fatalln(err)
			}
		case "review":
			review(configPath)
		case "reminisce":
			reminisce(configPath)
		case "search":
			hits, err := search(notesPath, os.Args[2:])
			if err != nil {
				log.Fatalln(err)
			}
			formatSearchResults(hits)
		case "goodnight":
			err = goodnight(configPath)
			if err != nil {
				log.Fatalln(err)
			}
		case "retro":
			err = retro(configPath)
			if err != nil {
				log.Fatalln(err)
			}
		default:
			log.Fatalln("Unrecognized command", os.Args[1])
			os.Exit(1)
		}
	}
}
