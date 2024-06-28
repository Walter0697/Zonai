package prompt

import (
	"strings"

	"github.com/Walter0697/zonai/model"
	"github.com/Walter0697/zonai/util"
	"github.com/manifoldco/promptui"
)

func ExecuteAdd() {
	options := []model.SimplePromptItemModel{
		{Name: "Project", Action: "Project"},
		{Name: "Deployment", Action: "Deployment"},
		{Name: "Back", Action: "Back"},
	}

	template := model.GetSimpleSelectTemplate("Add")
	searcher := model.GetSimpleSearcher(options)

	prompt := promptui.Select{
		Label:     "Which do you want to add?",
		Items:     options,
		Templates: template,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()

	action := options[i].Action

	switch action {
	case "Project":
		executeAddToList(action)
	case "Deployment":
		executeAddToList(action)
	case "Back":
		Execute()
	}
}

func executeAddToList(add_type string) {
	options := []model.ProjectPromptItemModel{}

	var list []model.ProjectParentModel
	if add_type == "Project" {
		list = util.ReadProjectList().List
	} else {
		list = util.ReadDeploymentList().List
	}

	for _, item := range list {
		var childList []string
		for _, child := range item.List {
			childList = append(childList, child.ProjectName)
		}

		childStr := strings.Join(childList, ", ")
		options = append(options, model.ProjectPromptItemModel{ProjectName: item.ProjectName, ChildList: childStr, Action: ""})
	}

	options = append(options, model.ProjectPromptItemModel{ProjectName: "New Project", ChildList: "", Action: "New"})
	options = append(options, model.ProjectPromptItemModel{ProjectName: "Back", ChildList: "", Action: "Back"})

	template := model.GetProjectSelectTemplate(add_type)
	searcher := model.GetProjectSearcher(options)

	prompt := promptui.Select{
		Label:     "Which " + add_type + " are you going to add?",
		Items:     options,
		Templates: template,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()

	action := options[i].Action

	if action == "Back" {
		ExecuteAdd()
		return
	}

	if action == "New" {
		executeNewParent(add_type)
		return
	}

	projectName := options[i].ProjectName
	executeNewChild(add_type, projectName)
}

func executeNewParent(add_type string) {
	prompt := promptui.Prompt{
		Label: "Enter the project name",
	}

	projectName, _ := prompt.Run()

	executeNewChild(add_type, projectName)
}

func executeNewChild(add_type string, parentName string) {
	prompt := promptui.Prompt{
		Label: "Enter the child " + add_type + " name",
	}

	childName, _ := prompt.Run()

	promptPath := promptui.Prompt{
		Label: "Enter the project path",
	}

	projectPath, _ := promptPath.Run()

	if add_type == "Project" {
		projectList := util.ReadProjectList()
		updatedList := util.AddProject(projectList, parentName, childName, projectPath)
		util.SaveProjectList(updatedList)
	} else {
		projectList := util.ReadDeploymentList()
		updatedList := util.AddProject(projectList, parentName, childName, projectPath)
		util.SaveDeploymentList(updatedList)
	}

}
