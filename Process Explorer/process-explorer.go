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

func getProcesses(defaultProcesses bool, searchString string) string {
	processes, err := ps.Processes()

	if err != nil {
		log.Fatal(err)
	}

	var outputString string

	for _, process := range processes {
		duration := processDuration(process)
		processName := strings.Split(process.Executable(), ".exe")[0]
		processPID := process.Pid()
		if !defaultProcesses || defaultProcesses && duration != defaultProcessDuration {
			if len(searchString) == 0 || len(searchString) != 0 && processName == searchString{
				outputFormat := fmt.Sprintf("\n%s: \n - Running for: %s \n - PID:[%d] \n - PPID:[%d]\n", processName, duration, processPID, process.PPid())
				outputString += outputFormat
			}
		}
	}

	return outputString
}

func processDuration(process ps.Process) string {
	_processCreationTime := process.CreationTime()
	_duration := time.Since(_processCreationTime)

	return _duration.String()
}

func main() {

	var processWindow, searchField *walk.TextEdit
	var toggleDefaultsCheckBox *walk.CheckBox
	showDefaultProcesses := false
	searchFieldString := ""

	MainWindow{
		Title:   "Application Stats",
		MinSize: Size{300, 600},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					CheckBox{
						AssignTo: &toggleDefaultsCheckBox,
						Text:     "Hide Defaults",
						Checked:  false,
						OnCheckStateChanged: func() {
							showDefaultProcesses = !showDefaultProcesses
							checkboxValue := strconv.FormatBool(showDefaultProcesses)
							checkboxOutput := fmt.Sprintf("Hide System Processes: %s \n", checkboxValue)
							processWindow.AppendText(checkboxOutput)
						},
					},
					TextEdit{
						AssignTo: &searchField,
					},
					PushButton{
						Text: "Filter",
						OnClicked: func() {
							processWindow.SetText("")
							searchFieldString = searchField.Text()
							returnedProcesses := getProcesses(showDefaultProcesses, searchFieldString)
							for _, applicationString := range strings.Split(returnedProcesses, "\n") {
								processWindow.AppendText(applicationString + "\r\n")
							}
						},
					},
				},
			},
			HSplitter{
				MinSize: Size{300, 570},
				Children: []Widget{
					TextEdit{AssignTo: &processWindow, ReadOnly: true},
				},
			},
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Get Data",
						OnClicked: func() {
							processWindow.SetText("")
							returnedProcesses := getProcesses(showDefaultProcesses, "")
							for _, applicationString := range strings.Split(returnedProcesses, "\n") {
								processWindow.AppendText(applicationString + "\r\n")
							}
						},
					},
				},
			},
		},
	}.Run()

}
