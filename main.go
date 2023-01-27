package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an input file name")
		return
	}
	inputFile := os.Args[1]

	// Read the contents of input file
	input, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Get the current time
	timestamp := time.Now().Format("2006-01-02 15-04-05")
	// Get the action taken
	action := os.Args[2]

	// Write the contents of input file to "output_<timestamp>.txt"
	if action == "update" {
		err = ioutil.WriteFile("output_"+timestamp+".txt", input, 0644)
		if err != nil {
			fmt.Println("Error writing output file:", err)
		}
	}

	//Read the state of the file
	state, err := ioutil.ReadFile("state.txt")
	if err != nil {
		fmt.Println("Error reading state file:", err)
	}
	// Append the contents of the input file and timestamp to the state file
	state = append(state, []byte("\n\nFile: "+inputFile+" Action: "+action+" Timestamp: "+timestamp+"\n"+string(input))...)
	err = ioutil.WriteFile("state.txt", state, 0644)
	if err != nil {
		fmt.Println("Error writing state file:", err)
	}
}
