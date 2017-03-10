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

type Process struct{
  name string
  duration time.Duration
  pid   int
  ppid  int
}


func getProcesses(defaultProcesses bool, searchString string) []Process {
  processes, err := ps.Processes()

  if err != nil {
    log.Fatal(err)
  }

  var output []Process

  for _, process := range processes {
    duration := processDuration(process)
    name := strings.Split(process.Executable(), ".exe")[0]
    pid := process.Pid()
    ppid := process.PPid()
    if !defaultProcesses || defaultProcesses && duration.String() != defaultProcessDuration{
      if len(searchString) == 0 || len(searchString) != 0 && name == searchString{
        currentProcess := Process{name, duration, pid, ppid}
        output = append(output, currentProcess)
      }
    }
  }

  return output
}

func processDuration(process ps.Process) time.Duration {
  _processCreationTime := process.CreationTime()
  _duration := time.Since(_processCreationTime)

  return _duration
}

func outputToProcessWindow(processWindow *walk.TextEdit, returnedProcesses []Process) {
  fmt.Println(returnedProcesses)
  for _, singleProcess := range returnedProcesses {
    name := singleProcess.name
    duration := singleProcess.duration
    outputString := fmt.Sprintf(" - %s \r -%d", name, duration)
	  for _, applicationString := range strings.Split(outputString, "\n") {
	    processWindow.AppendText(applicationString + "\r\n")
	  }
  }
}

func main() {

  var processWindow, searchField *walk.TextEdit
  var toggleDefaultsCheckBox *walk.CheckBox
  showDefaultProcesses := false
  searchFieldString := ""

  MainWindow{
    Title:   "Go Look At Processes!",
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
              outputToProcessWindow(processWindow,returnedProcesses)
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
            Text: "Get Processes",
            OnClicked: func() {
              processWindow.SetText("")
              returnedProcesses := getProcesses(showDefaultProcesses, "")
              outputToProcessWindow(processWindow,returnedProcesses)
            },
          },
        },
      },
    },
  }.Run()

}
