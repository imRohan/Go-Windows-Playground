package main

import (
	"fmt"
	"github.com/imRohan/go-ps"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"strings"
	"time"
)

func getProcesses() string {
	processes, err := ps.Processes()

	if err != nil {
		log.Fatal(err)
	}

	var output_string string

	for _, process := range processes {
		output_string += fmt.Sprintf("Process Name: %s, Time Duration: %s \n\n", process.Executable(), processDuration(process))
	}

	return output_string
}

func processDuration(process ps.Process) string {
	_processCreationTime := process.CreationTime()
	_duration := time.Since(_processCreationTime)

	return _duration.String()
}

func main() {

	var applications *walk.TextEdit

	MainWindow{
		Title:   "Application Stats",
		MinSize: Size{500, 300},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &applications, ReadOnly: true},
				},
			},
			PushButton{
				Text: "Refresh",
				OnClicked: func() {
					for _, applicationString := range strings.Split(getProcesses(), "\n") {
						applications.AppendText(applicationString + "\r\n")
					}
				},
			},
		},
	}.Run()

}
