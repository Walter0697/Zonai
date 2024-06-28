package prompt

import (
	"fmt"

	"github.com/Walter0697/zonai/model"
	"github.com/Walter0697/zonai/util"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func ExecuteConfig() {
	options := []model.SimplePromptItemModel{
		{Name: "Show", Action: "Show"},
		{Name: "Setup", Action: "Setup"},
		{Name: "Back", Action: "Back"},
	}
	template := model.GetSimpleSelectTemplate("Configuration")
	searcher := model.GetSimpleSearcher(options)

	prompt := promptui.Select{
		Label:     "Please select an options?",
		Items:     options,
		Templates: template,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()

	action := options[i].Action

	switch action {
	case "Show":
		util.ShowConfiguration()
		ExecuteConfig()
	case "Setup":
		executeSetup()
	case "Back":
		Execute()
	}
}

func executeSetup() {
	options := []model.SimplePromptItemModel{
		{Name: "Output Image Path", Action: "output"},
		{Name: "Input Image Path", Action: "input"},
		{Name: "Environment Path", Action: "env"},
		{Name: "Docker Build Command", Action: "docker"},
		{Name: "Back", Action: "Back"},
	}
	template := model.GetSimpleSelectTemplate("Setup")
	searcher := model.GetSimpleSearcher(options)

	prompt := promptui.Select{
		Label:     "Please select an options",
		Items:     options,
		Templates: template,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()

	action := options[i].Action

	displayText := ""
	configuration := util.ReadConfiguration()
	switch action {
	case "output":
		displayText = configuration.OutputImagePath
	case "input":
		displayText = configuration.InputImagePath
	case "env":
		displayText = configuration.EnviromentPath
	case "docker":
		displayText = configuration.DockerBuildCommand
	case "Back":
		ExecuteConfig()
		return
	}

	text := color.YellowString(displayText)
	fmt.Println("Current value: ", text)
	promptText := promptui.Prompt{
		Label: "Please enter the new value",
	}
	result, _ := promptText.Run()

	switch action {
	case "output":
		configuration.OutputImagePath = result
	case "input":
		configuration.InputImagePath = result
	case "env":
		configuration.EnviromentPath = result
	case "docker":
		configuration.DockerBuildCommand = result
	}

	util.SaveConfiguration(configuration)

	color.Cyan("Set up %s...", options[i].Name)

	ExecuteConfig()
}
