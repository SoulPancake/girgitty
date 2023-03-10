
# GirGitty
![image](https://user-images.githubusercontent.com/70265851/215025499-220f394c-c3bb-46e5-a802-68a96668c976.png)

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

`./GirGitty -f <input file>`

It will create a state.txt file in the same directory where the input file is located. This file will contain the log of all the changes made to the input file with timestamps.

## Example

`./GirGitty -f test.txt`


## The above program can be compiled and run using the following command:

`go run version.go -f <input_file> -v <version>`


This will start monitoring the input file for changes and update the state file with the timestamp whenever a change is detected. The state file will be named as `<version>_state.txt` where `<version>` is the version of the file passed as an argument to the program. The program is using the `fsnotify` package to listen for file changes and it can detect changes like create, update, rename, and delete events.

In order to use GirGitty, you can use the above code as a driver for your application and you can use the `version` flag to keep track of different versions of the same file.

You can also use the following command to switch between different versions of the file:

`go run app.go -f <input_file> -v <version>`

This will create a new file with the version name appended to it and it will update the state file with the commit information.

You can also integrate this functionality into the main program by calling this function when the user wants to commit the changes.

`go run commit.go <input_file> <version>`

This command will start monitoring test.txt file. If the file is modified, renamed, deleted or created, it will update the state.txt file with the respective timestamp and action.

# Progress Features
- Track file changes, including create, update, rename, and delete events.
- Log file changes with timestamps and version information.
- CLI app for easy access to the functionalities.
- Option to have different versions for the same file.
- Logging for different versions.
- Commit functionality in a separate file.
- API endpoints to use the functionalities over a Google Drive folder.
- Versioning used for the Google Drive API.
- Ability to create a new version of the file with the commit message.
- Easy integration with existing projects.

In the future, we plan to integrate GirGitty with Google Drive to allow for seamless version control of files stored on the cloud platform. This will enable users to track changes made to their files and easily revert to previous versions if needed. Additionally, the ability to access and manage versions of files from anywhere with an internet connection will greatly enhance the usability of the tool. We also plan to integrate the commit functionality with Google Drive API, so that users can easily track the changes made in the file and also the user who made the changes. This will also help in collaboration for the same file with different users.

## Contribution
We welcome contributions to this project. If you are interested in contributing, please fork the repository and open a pull request.

## License
GirGitty is released under the MIT License.
