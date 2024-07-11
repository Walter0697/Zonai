package prompt

import (
	"fmt"

	"github.com/Walter0697/zonai/model"
	"github.com/Walter0697/zonai/util"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func ExecuteWrapper() {
	options := []model.SimplePromptItemModel{
		{Name: "docker ps", Action: "PS"},
		{Name: "docker log", Action: "Log"},
		{Name: "docker exec", Action: "Exec"},
		{Name: "Back", Action: "Back"},
	}

	template := model.GetSimpleSelectTemplate("Wrapper")
	searcher := model.GetSimpleSearcher(options)

	prompt := promptui.Select{
		Label:     "Which do command line do you want to call?",
		Items:     options,
		Templates: template,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()

	action := options[i].Action

	switch action {
	case "PS":
		executeDockerPs()
	case "Log":
		executeDockerCommand("Docker Log")
	case "Exec":
		executeDockerCommand("Docker Exec")
	case "Back":
		Execute()
	}
}

func executeDockerPs() {
	deploymentList := util.ReadDeploymentList()
	options := []model.SimplePromptItemModel{}
	for _, project := range deploymentList.List {
		options = append(options, model.SimplePromptItemModel{Name: project.ProjectName, Action: "Build"})
	}
	options = append(options, model.SimplePromptItemModel{Name: "Back", Action: "Back"})

	searcher := model.GetSimpleSearcher(options)
	templates := model.GetSimpleSelectTemplate("Build")

	prompt := promptui.Select{
		Label:     "Which project do you want to see the docker ps?",
		Items:     options,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()
	action := options[i].Action

	if action == "Back" {
		ExecuteWrapper()
		return
	}

	projectName := options[i].Name
	color.Cyan("----------------------------------------")
	color.Cyan("Below is the docker ps for the project")
	color.Cyan("Project: " + projectName)
	color.Cyan("----------------------------------------")

	output, _ := util.DockerPs(projectName, true)
	fmt.Println(output)
}

func executeDockerCommand(commandType string) {
	deploymentList := util.ReadDeploymentList()
	options := []model.SimplePromptItemModel{}
	for _, project := range deploymentList.List {
		options = append(options, model.SimplePromptItemModel{Name: project.ProjectName, Action: "Build"})
	}
	options = append(options, model.SimplePromptItemModel{Name: "Back", Action: "Back"})

	searcher := model.GetSimpleSearcher(options)
	templates := model.GetSimpleSelectTemplate(commandType)

	prompt := promptui.Select{
		Label:     "Which project do you want to see the docker log?",
		Items:     options,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()
	action := options[i].Action

	if action == "Back" {
		ExecuteWrapper()
		return
	}

	projectName := options[i].Name

	var currentProject *model.ProjectParentModel
	for _, project := range deploymentList.List {
		if project.ProjectName == projectName {
			currentProject = &project
			break
		}
	}

	if currentProject == nil {
		color.Red("Unexpected error occurred, " + projectName + " not found")
		Execute()
		return
	}

	executeChildDockerLog(currentProject, commandType)
}

func executeChildDockerLog(project *model.ProjectParentModel, commandType string) {
	options := []model.SimplePromptItemModel{}
	for _, child := range project.List {
		options = append(options, model.SimplePromptItemModel{Name: child.ProjectName, Action: "Log"})
	}
	options = append(options, model.SimplePromptItemModel{Name: "Back", Action: "Back"})

	searcher := model.GetSimpleSearcher(options)
	templates := model.GetSimpleSelectTemplate("Child Project")

	prompt := promptui.Select{
		Label:     "Which child project do you want to see the docker log?",
		Items:     options,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()
	action := options[i].Action

	if action == "Back" {
		executeDockerCommand(commandType)
		return
	}

	childName := options[i].Name
	containerId, targetName, err := util.GetContainerId(project.ProjectName, childName)
	if err == nil {
		fmt.Println("Found image : " + color.GreenString(targetName))
		fmt.Println("Please use the following command to see the log")
		if commandType == "Docker Log" {
			color.Yellow("docker logs " + containerId)
		} else if commandType == "Docker Exec" {
			color.Yellow("docker exec -it " + containerId + " sh")
		}
		return
	}

	color.Red(targetName + " not found, try to check if it is running")
}
