package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
)

var (
	inputFile string
	version   string
	stateFile string
)

func init() {
	flag.StringVar(&inputFile, "f", "", "input file to be monitored")
	flag.StringVar(&version, "v", "master", "version of the file")
	flag.Parse()
	stateFile = version + "_state.txt"
}

func versions() {
	// Check if input file exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Println("Error: Input file does not exist.")
		os.Exit(1)
	}

	// Initialize watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer watcher.Close()

	done := make(chan bool)

	// Watch input file for changes
	err = watcher.Add(inputFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Event loop
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				timestamp := time.Now().Format("2006-01-02 15:04:05")
				if event.Op&fsnotify.Write == fsnotify.Write {
					// Append the contents of the input file and timestamp to the state file
					state, err := ioutil.ReadFile(stateFile)
					if err != nil {
						fmt.Println("Error reading state file:", err)
					}
					state = append(state, []byte("\n\nFile: "+inputFile+" Version: "+version+" Action: update"+" Timestamp: "+timestamp+"\n")...)
					err = ioutil.WriteFile(stateFile, state, 0644)
					if err != nil {
						fmt.Println("Error writing state file:", err)
					}
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					// Append the contents of the input file and timestamp to the state file
					state, err := ioutil.ReadFile(stateFile)
					if err != nil {
						fmt.Println("Error reading state file:", err)
					}
					state = append(state, []byte("\n\nFile: "+inputFile+" Version: "+version+" Action: create"+" Timestamp: "+timestamp+"\n")...)
					err = ioutil.WriteFile(stateFile, state, 0644)
					if err != nil {
						fmt.Println("Error writing state file:", err)
					}
				}
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					// Append the contents of the input file and timestamp to the state file
					state, err := ioutil.ReadFile(stateFile)
					if err != nil {
						fmt.Println("Error reading state file:", err)
					}
					state = append(state, []byte("\n\nFile: "+inputFile+" Version: "+version+" Action: rename"+" Timestamp: "+timestamp+"\n")...)
					err = ioutil.WriteFile(stateFile, state, 0644)
					if err != nil {
						fmt.Println("Error writing state file:", err)
					}
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					// Append the contents of the input file and timestamp to the state file
					state, err := ioutil.ReadFile(stateFile)
					if err != nil {
						fmt.Println("Error reading state file:", err)
					}
					state = append(state, []byte("\n\nFile: "+inputFile+" Version: "+version+" Action: delete"+" Timestamp: "+timestamp+"\n")...)
					err = ioutil.WriteFile(stateFile, state, 0644)
					if err != nil {
						fmt.Println("Error writing state file:", err)
					}
				}
			case err := <-watcher.Errors:
				fmt.Println("Error:", err)
			}
		}
	}()
	<-done
}
