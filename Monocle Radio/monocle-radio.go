package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var webWindow *walk.WebView

	MainWindow{
		Title:   "Monocle Radio",
		MinSize: Size{600, 400},
		Layout: VBox{
			MarginsZero: true,
		},
		Children: []Widget{
			WebView{
				AssignTo: &webWindow,
				Name:     "Monocle Radio",
				URL:      "https://monocle.com/radio/",
			},
		},
	}.Run()
}
