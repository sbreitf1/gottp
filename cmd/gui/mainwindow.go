package main

import (
	"github.com/sbreitf1/gottp/internal/gottp"
	"github.com/therecipe/qt/widgets"
)

type MainWindow struct {
	window *widgets.QMainWindow
}

func NewMainWindow() *MainWindow {
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("gottp")
	window.SetMinimumSize2(400, 300)

	{
		scenarioWidget := NewScenarioWidget(gottp.Scenario{
			Name:   "Example",
			Method: "GET",
			URL:    "https://sbreitf1.de",
		})

		window.SetCentralWidget(scenarioWidget.widget)
	}

	return &MainWindow{window: window}
}

func (w *MainWindow) Show() {
	w.window.Show()
}
