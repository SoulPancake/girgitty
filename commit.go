package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Commit(inputFile, version, stateFile string) {
	// read the input file
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// create a new file with the version name
	newFile := strings.Join([]string{version, inputFile}, "_")
	err = ioutil.WriteFile(newFile, data, 0644)
	if err != nil {
		fmt.Println("Error creating new file:", err)
		return
	}

	// update the state file with the commit information
	state, err := ioutil.ReadFile(stateFile)
	if err != nil {
		fmt.Println("Error reading state file:", err)
		return
	}
	timestamp := CurrentTime()
	state = append(state, []byte("\n\nFile: "+inputFile+" Version: "+version+" Action: commit"+" Timestamp: "+timestamp+"\n")...)
	err = ioutil.WriteFile(stateFile, state, 0644)
	if err != nil {
		fmt.Println("Error writing state file:", err)
		return
	}

	fmt.Println("Commit successful!")
}

func CurrentTime() {
	panic("unimplemented")
}

func commitRunner() {
	inputFile := os.Args[1]
	version := os.Args[2]
	stateFile := version + "_state.txt"

	Commit(inputFile, version, stateFile)
}
