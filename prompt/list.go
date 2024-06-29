package prompt

import (
	"fmt"

	"github.com/Walter0697/zonai/model"
	"github.com/Walter0697/zonai/util"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func ExecuteList() {
	executeListWithSearch(nil)
}

func executeListWithSearch(searchKey *string) {
	if searchKey != nil && *searchKey != "" {
		displayText := color.YellowString(*searchKey)
		fmt.Println("Current Search Key: " + displayText)
	}

	options := []model.PromptItemModel{
		{Name: "All", Description: "List all projects and deployment instructions"},
		{Name: "Project", Description: "List all projects"},
		{Name: "Deployment", Description: "List all deployment instructions"},
		{Name: "Search", Description: "Search for a project or deployment"},
	}

	if searchKey != nil && *searchKey != "" {
		options = append(options, model.PromptItemModel{Name: "Clear Search", Description: "Clear the search key"})
	}

	options = append(options, model.PromptItemModel{Name: "Back", Description: "Go back to the main menu"})

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
		util.ListProject(searchKey)
		util.ListDeployment(searchKey)
		executeListWithSearch(searchKey)
	case "Project":
		// List all projects
		util.ListProject(searchKey)
		executeListWithSearch(searchKey)
	case "Deployment":
		// List all deployment instructions
		util.ListDeployment(searchKey)
		executeListWithSearch(searchKey)
	case "Search":
		// Search for a project or deployment
		executeSearchKey()
	case "Clear Search":
		// Clear the search key
		executeListWithSearch(nil)
	case "Back":
		// Go back to the main menu
		Execute()
	}
}

func executeSearchKey() {
	prompt := promptui.Prompt{
		Label: "Enter search key",
	}

	searchKey, _ := prompt.Run()
	executeListWithSearch(&searchKey)
}
