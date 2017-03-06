package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"strings"
)

func reverse(input string) string {
	_inputToByte := []byte(input)
	_lengthOfByteRep := len(_inputToByte)
	_reverseString := make([]byte, _lengthOfByteRep)
	_reverseIndex := _lengthOfByteRep - 1
	for i := 0; i < _lengthOfByteRep; i++ {
		_currentElement := _inputToByte[_reverseIndex]
		_reverseString[i] = _currentElement
		_reverseIndex--
	}

	return string(_reverseString)
}

func main() {
	var inputText, reverseText, capitalizeText *walk.TextEdit

	MainWindow{
		Title:   "Text Transform",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inputText},
					TextEdit{AssignTo: &reverseText, ReadOnly: true},
					TextEdit{AssignTo: &capitalizeText, ReadOnly: true},
				},
			},
			PushButton{
				Text: "Transform!",
				OnClicked: func() {
					capitalizeText.SetText(strings.ToUpper(inputText.Text()))
					capitalizeText.SetText(strings.ToUpper(inputText.Text()))
					reverseText.SetText(reverse(inputText.Text()))
				},
			},
		},
	}.Run()
}
