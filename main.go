package main

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/exp/inotify"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
)

var watchpath string = "/"
var command string
var exts []string = []string{".go", ".js", ".php", ".py", ".test"}

func main() {
	if len(os.Args) == 3 {
		watchpath = os.Args[1]
		command = os.Args[2]
	} else {
		fmt.Println("Usage: reloader <path> <command> &")
		return
	}
	// log.Println("watch:", watchpath)
	// log.Println("command:", command)
	watcher, err := inotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.Watch(watchpath)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case ev := <-watcher.Event:
			//DEBUG:log.Println("event:", ev)
			if ev.Mask == syscall.IN_CLOSE_WRITE ||
				ev.Mask == syscall.IN_DELETE {
				if hasExt(path.Ext(ev.Name)) {
					flds := strings.Fields(command)
					_, err := execShell(flds[0], flds[1:])
					if err != nil {
						log.Println(err)
					}
					// log.Println(execShell(flds[0], flds[1:]))
					// if s := strings.Index(command, " "); s > -1 {
					// 	log.Println(command[:s])
					// 	log.Println(command[s+1:])
					// 	log.Println(execShell(command[:s], []string{command[s+1:]}))
					// }
				}
			}
		case err := <-watcher.Error:
			log.Println("error:", err)
		}
	}
}
func hasExt(ext string) bool {
	for _, e := range exts {
		if e == ext {
			return true
		}
	}
	return false
}

//get output from shell command
func execShell(cmd string, args []string) (string, error) {
	sh := exec.Command(cmd, args...)
	var out bytes.Buffer
	var errString bytes.Buffer
	sh.Stdout = &out
	sh.Stderr = &errString
	err := sh.Run()
	if err != nil || len(errString.String()) > 0 {
		errStr := "Error in command: " + errString.String()
		errStr += "(command: " + cmd + " " + strings.Join(args, " ") + ")"
		return "", errors.New(errStr)
	}
	return out.String(), nil
}
