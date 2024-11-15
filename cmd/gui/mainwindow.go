package main

import (
	"github.com/therecipe/qt/widgets"
)

type MainWindow struct {
	window *widgets.QMainWindow
}

func NewMainWindow() *MainWindow {
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("gottp")
	window.SetMinimumSize2(400, 100)

	scenarioWidget := widgets.NewQWidget(nil, 0)
	scenarioLayout := widgets.NewQVBoxLayout()
	scenarioWidget.SetLayout(scenarioLayout)

	httpMethodList := widgets.NewQComboBox(nil)

	scenarioLayout.AddWidget(httpMethodList, 0, 0)

	window.SetCentralWidget(scenarioWidget)

	return &MainWindow{window: window}
}

func (w *MainWindow) Show() {
	w.window.Show()
}
