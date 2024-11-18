package main

import (
	"github.com/sbreitf1/gottp/internal/gottp"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type ScenarioWidget struct {
	widget *widgets.QWidget
}

func NewScenarioWidget(scenario gottp.Scenario) *ScenarioWidget {
	scenarioWidget := widgets.NewQWidget(nil, 0)
	scenarioLayout := widgets.NewQVBoxLayout()
	scenarioWidget.SetLayout(scenarioLayout)

	{
		urlWidget := widgets.NewQWidget(nil, 0)
		urlLayout := widgets.NewQHBoxLayout()
		urlWidget.SetLayout(urlLayout)

		httpMethodList := widgets.NewQComboBox(nil)
		httpMethodList.AddItem("GET", core.NewQVariant())
		httpMethodList.AddItem("POST", core.NewQVariant())
		httpMethodList.AddItem("PUT", core.NewQVariant())
		httpMethodList.AddItem("DELETE", core.NewQVariant())
		urlLayout.AddWidget(httpMethodList, 0, 0)

		urlTextBox := widgets.NewQLineEdit(nil)
		urlTextBox.SetText(scenario.URL)
		urlLayout.AddWidget(urlTextBox, 0, 0)

		sendButton := widgets.NewQPushButton(nil)
		sendButton.SetText("Send")
		urlLayout.AddWidget(sendButton, 0, 0)

		scenarioLayout.AddWidget(urlWidget, 0, 0)
	}

	{
		requestTabGroup := widgets.NewQTabWidget(nil)

		{
			headersTab := widgets.NewQWidget(nil, 0)
			headersLayout := widgets.NewQHBoxLayout()
			headersTab.SetLayout(headersLayout)

			headersList := widgets.NewQTableView(nil)
			headersModel := core.NewQAbstractListModel(nil)
			headersModel.SetHeaderData(0, core.Qt__Horizontal, core.NewQVariant1("Key"), 0)
			headersModel.SetHeaderData(1, core.Qt__Horizontal, core.NewQVariant1("Value"), 0)
			headersList.SetModel(headersModel)
			headersLayout.AddWidget(headersList, 0, 0)

			requestTabGroup.AddTab(headersTab, "Headers")
		}

		{
			bodyTab := widgets.NewQWidget(nil, 0)
			bodyLayout := widgets.NewQVBoxLayout()
			bodyTab.SetLayout(bodyLayout)

			bodyTextEdit := widgets.NewQTextEdit(nil)
			bodyLayout.AddWidget(bodyTextEdit, 0, 0)

			requestTabGroup.AddTab(bodyTab, "Body")
		}

		scenarioLayout.AddWidget(requestTabGroup, 0, 0)
	}

	{
		responseGroupBox := widgets.NewQGroupBox(nil)
		responseLayout := widgets.NewQVBoxLayout()
		responseGroupBox.SetLayout(responseLayout)
		responseGroupBox.SetTitle("Response")

		responseTextEdit := widgets.NewQTextEdit(nil)
		responseTextEdit.SetReadOnly(true)
		responseLayout.AddWidget(responseTextEdit, 0, 0)

		scenarioLayout.AddWidget(responseGroupBox, 0, 0)
	}

	return &ScenarioWidget{widget: scenarioWidget}
}
