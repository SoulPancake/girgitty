# GirGitty
A simple version control system for tracking and logging file changes in Go.

## Features
- Logs file changes, including create, update, rename and delete events
- Keeps a timestamp for each change
- CLI for easy integration with existing programs
- Utilizes fsnotify for efficient file monitoring

## Installation
1. Download the source code from this repository
2. Run `go build` to build the program
3. Run `./GirGitty -f <input file>` to start monitoring the file

## Usage
GirGitty can be used by passing the file to be monitored as a command line argument. 

./GirGitty -f <input file>

It will create a state.txt file in the same directory where the input file is located. This file will contain the log of all the changes made to the input file with timestamps.

## Example

./GirGitty -f test.txt


This command will start monitoring test.txt file. If the file is modified, renamed, deleted or created, it will update the state.txt file with the respective timestamp and action.

## Contribution
We welcome contributions to this project. If you are interested in contributing, please fork the repository and open a pull request.

## License
GirGitty is released under the MIT License.