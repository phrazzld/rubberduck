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
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	f.WriteString(conf)
	f.Sync()
}
