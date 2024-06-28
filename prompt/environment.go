package prompt

import (
	"fmt"

	"github.com/Walter0697/zonai/model"
	"github.com/Walter0697/zonai/util"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func ExecuteEnvironment() {
	configuration := util.ReadConfiguration()
	currentEnvrionment := color.YellowString(configuration.CurrentEnvironment)
	fmt.Println("Current Environment is " + currentEnvrionment)

	options := []model.SimplePromptItemModel{
		{Name: "Change Environment", Action: "Change"},
		{Name: "Back", Action: "Back"},
	}
	template := model.GetSimpleSelectTemplate("Environment")
	searcher := model.GetSimpleSearcher(options)

	prompt := promptui.Select{
		Label:     "What do you want to do?",
		Items:     options,
		Templates: template,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()

	action := options[i].Action
	switch action {
	case "Change":
		executeChangeEnvironment()
	case "Back":
		Execute()
	}
}

func executeChangeEnvironment() {
	prompt := promptui.Prompt{
		Label: "Enter the environment name",
	}

	environment, _ := prompt.Run()

	configuration := util.ReadConfiguration()
	configuration.CurrentEnvironment = environment
	util.SaveConfiguration(configuration)

	color.Cyan("Environment changed to " + environment)
	ExecuteEnvironment()
}
