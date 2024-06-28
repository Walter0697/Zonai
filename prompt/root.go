package prompt

import (
	"fmt"
	"os"

	"github.com/Walter0697/zonai/model"
	"github.com/manifoldco/promptui"
)

const ExitText = "Thank you for using Zonai! Have a nice day! ٩(╹ꇴ ╹๑)۶"

func Execute() {
	options := []model.PromptItemModel{
		{Name: "Deploy", Description: "Deploy a project using a compressed .gz file"},
		{Name: "Build", Description: "Build a project to a compressed .gz file"},
		{Name: "Add", Description: "Add a new project or deployment instruction to the list"},
		{Name: "Delete", Description: "Delete a project or deployment instruction from the list"},
		{Name: "List", Description: "List all projects and deployment instructions"},
		{Name: "Environment", Description: "Manage the environment settings"},
		{Name: "Configuration", Description: "Configure the project settings"},
		{Name: "Command Line Tools", Description: "How to use command line tools"},
		{Name: "Version", Description: "Show the version of the program"},
		{Name: "Exit", Description: "Exit the program"},
	}

	templates := model.GetSelectTemplate("Main Menu")
	searcher := model.GetSearcher(options)

	prompt := promptui.Select{
		Label:     "Please select the following action:",
		Items:     options,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()
	name := options[i].Name

	switch name {
	case "Deploy":
		ExecuteDeploy()
	case "Build":
		ExecuteBuild()
	case "Add":
		ExecuteAdd()
	case "Delete":
		ExecuteDelete()
	case "List":
		ExecuteList()
	case "Environment":
		ExecuteEnvironment()
	case "Configuration":
		ExecuteConfig()
	case "Command Line Tools":
		ExecuteCli()
	case "Version":
		ExecuteVersion()
	case "Exit":
		fmt.Println(ExitText)
		os.Exit(1)
	}

	fmt.Println(ExitText)
}
