package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"time"
)

func listen() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an input file name")
		return
	}
	inputFile := os.Args[1]

	// Get the current time
	timestamp := time.Now().Format("2006-01-02 15-04-05")

	// Check if input file exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Println("Input file does not exist")
		return
	}

	// Create a new fsnotify watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error creating fsnotify watcher:", err)
		return
	}
	defer watcher.Close()

	// Add the input file to the watcher
	err = watcher.Add(inputFile)
	if err != nil {
		fmt.Println("Error adding file to fsnotify watcher:", err)
		return
	}

	// Listen for events
	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				// Append the contents of the input file and timestamp to the state file
				state, err := ioutil.ReadFile("state.txt")
				if err != nil {
					fmt.Println("Error reading state file:", err)
				}
				state = append(state, []byte("\n\nFile: "+inputFile+" Action: update"+" Timestamp: "+timestamp+"\n")...)
				err = ioutil.WriteFile("state.txt", state, 0644)
				if err != nil {
					fmt.Println("Error writing state file:", err)
				}
			}
		case err := <-watcher.Errors:
			fmt.Println("Error from fsnotify:", err)
		}
	}
}
