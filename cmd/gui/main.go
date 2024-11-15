package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	app.SetParent(nil)
	app.SetApplicationName("gottp")

	mainWindow := NewMainWindow()
	mainWindow.Show()

	widgets.QApplication_Exec()
}
