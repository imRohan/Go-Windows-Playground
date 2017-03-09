package main

import (
	"fmt"
	"github.com/imRohan/go-ps"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	defaultProcessDuration = "2562047h47m16.854775807s"
)

func getProcesses(defaultProcesses bool) string {
	processes, err := ps.Processes()

	if err != nil {
		log.Fatal(err)
	}

	var output_string string

	for _, process := range processes {
		duration := processDuration(process)
		if defaultProcesses && duration != defaultProcessDuration{
						output_string += fmt.Sprintf("\n%s: \n - Running for: %s \n - PID:[%d] \n - PPID:[%d]\n", process.Executable(), duration, process.Pid(), process.PPid())
		} else if !defaultProcesses {
						output_string += fmt.Sprintf("\n%s: \n - Running for: %s \n - PID:[%d] \n - PPID:[%d]\n", process.Executable(), duration, process.Pid(), process.PPid())
		}
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
	var toggleDefaultsCheckBox *walk.CheckBox
	showDefaultProcesses := false

	MainWindow{
		Title:   "Application Stats",
		MinSize: Size{300, 600},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				MinSize: Size{300, 570},
				Children: []Widget{
					TextEdit{AssignTo: &applications, ReadOnly: true},
				},
			},
			HSplitter{
				Children: []Widget{
					CheckBox{
						AssignTo: &toggleDefaultsCheckBox,
						Text:     "Hide Defaults",
						Checked:  false,
						OnCheckStateChanged: func() {
							showDefaultProcesses = !showDefaultProcesses
							checkboxOutput := fmt.Sprintf("Hide System Processes: %s \n", strconv.FormatBool(showDefaultProcesses))
							applications.AppendText(checkboxOutput)
						},
					},
					PushButton{
						Text: "Get Data",
						OnClicked: func() {
							applications.SetText("")
							for _, applicationString := range strings.Split(getProcesses(showDefaultProcesses), "\n") {
								applications.AppendText(applicationString + "\r\n")
							}
						},
					},
				},
			},
		},
	}.Run()

}
