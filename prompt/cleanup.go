package prompt

import (
	"github.com/Walter0697/zonai/model"
	"github.com/Walter0697/zonai/util"
	"github.com/manifoldco/promptui"
)

func ExecuteCleanup() {
	options := []model.SimplePromptItemModel{
		{Name: "All", Action: "All"},
		{Name: "Input Image Folder", Action: "Input"},
		{Name: "Output Image Folder", Action: "Output"},
		{Name: "Back", Action: "Back"},
	}

	template := model.GetSimpleSelectTemplate("Clean Up")
	searcher := model.GetSimpleSearcher(options)

	prompt := promptui.Select{
		Label:     "Which do you want to clean up?",
		Items:     options,
		Templates: template,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()

	action := options[i].Action

	switch action {
	case "All":
		util.CleanupAll()
	case "Input":
		util.CleanupInputFolder()
	case "Output":
		util.CleanupOutputFolder()
	case "Back":
		Execute()
	}
}
