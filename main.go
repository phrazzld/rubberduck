// rubberduck/main.go

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
			hits := search(duckPath, os.Args[2:])
			formatSearchResults(hits)
		default:
			fmt.Println("Unrecognized command")
			os.Exit(1)
		}
	}
}
