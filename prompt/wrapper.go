package prompt

import (
	"fmt"
	"strings"

	"github.com/Walter0697/zonai/model"
	"github.com/Walter0697/zonai/util"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func ExecuteWrapper() {
	options := []model.SimplePromptItemModel{
		{Name: "docker ps", Action: "PS"},
		{Name: "docker log", Action: "Log"},
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
		executeDockerLog()
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
		Execute()
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

func executeDockerLog() {
	deploymentList := util.ReadDeploymentList()
	options := []model.SimplePromptItemModel{}
	for _, project := range deploymentList.List {
		options = append(options, model.SimplePromptItemModel{Name: project.ProjectName, Action: "Build"})
	}
	options = append(options, model.SimplePromptItemModel{Name: "Back", Action: "Back"})

	searcher := model.GetSimpleSearcher(options)
	templates := model.GetSimpleSelectTemplate("Docker Log")

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
		Execute()
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

	executeChildDockerLog(currentProject)
}

func executeChildDockerLog(project *model.ProjectParentModel) {
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
		Execute()
		return
	}

	childName := options[i].Name
	output, _ := util.DockerPs(project.ProjectName, false)
	targetName := project.ProjectName + "/" + childName

	outputList := strings.Split(output, "\n")
	for index, line := range outputList {
		if index == 0 {
			// 0 is the header
			continue
		}
		infoSplit := strings.Split(line, "   ")
		if len(infoSplit) >= 2 {
			imageName := infoSplit[1]
			imageInfo := strings.Split(imageName, ":")[0]
			if imageInfo == targetName {
				fmt.Println("Found image : " + color.GreenString(imageInfo))
				containerId := infoSplit[0]
				fmt.Println("Please use the following command to see the log")
				color.Yellow("docker logs " + containerId)

				return
			}
		}
	}

	color.Red(targetName + " not found, try to check if it is running")
}
