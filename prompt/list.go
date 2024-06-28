package prompt

import (
	"github.com/Walter0697/zonai/model"
	"github.com/Walter0697/zonai/util"
	"github.com/manifoldco/promptui"
)

func ExecuteList() {
	options := []model.PromptItemModel{
		{Name: "All", Description: "List all projects and deployment instructions"},
		{Name: "Project", Description: "List all projects"},
		{Name: "Deployment", Description: "List all deployment instructions"},
		{Name: "Back", Description: "Go back to the main menu"},
	}

	templates := model.GetSelectTemplate("List")
	searcher := model.GetSearcher(options)

	prompt := promptui.Select{
		Label:     "Which list do you want to look at?",
		Items:     options,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()
	name := options[i].Name

	switch name {
	case "All":
		// List all projects and deployment instructions
		util.ListProject()
		util.ListDeployment()
		ExecuteList()
	case "Project":
		// List all projects
		util.ListProject()
		ExecuteList()
	case "Deployment":
		// List all deployment instructions
		util.ListDeployment()
		ExecuteList()
	case "Back":
		// Go back to the main menu
		Execute()
	}
}
