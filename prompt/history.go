package prompt

import (
	"github.com/Walter0697/zonai/model"
	"github.com/Walter0697/zonai/util"
	"github.com/manifoldco/promptui"
)

func ExecuteHistory() {
	options := []model.SimplePromptItemModel{
		{Name: "List All History", Action: "List"},
		{Name: "List History by Project", Action: "ListByProject"},
		{Name: "Back", Action: "Back"},
	}

	template := model.GetSimpleSelectTemplate("History")
	searcher := model.GetSimpleSearcher(options)

	prompt := promptui.Select{
		Label:     "Select an option?",
		Items:     options,
		Templates: template,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()

	action := options[i].Action
	switch action {
	case "List":
		util.DisplayHistory(nil)
	case "ListByProject":
		executeHistoryByProject()
	case "Back":
		Execute()
	}
}

func executeHistoryByProject() {
	prompt := promptui.Prompt{
		Label: "Enter the project name",
	}

	projectName, _ := prompt.Run()

	util.DisplayHistory(&projectName)
	ExecuteHistory()
}
