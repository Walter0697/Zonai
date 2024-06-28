package prompt

import (
	"strings"

	"github.com/Walter0697/zonai/model"
	"github.com/Walter0697/zonai/util"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func ExecuteDelete() {
	options := []model.SimplePromptItemModel{
		{Name: "Project", Action: "Project"},
		{Name: "Deployment", Action: "Deployment"},
		{Name: "Back", Action: "Back"},
	}

	template := model.GetSimpleSelectTemplate("Delete")
	searcher := model.GetSimpleSearcher(options)

	prompt := promptui.Select{
		Label:     "Which do you want to delete?",
		Items:     options,
		Templates: template,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()

	action := options[i].Action

	switch action {
	case "Project":
		executeDeleteFromList(action)
	case "Deployment":
		executeDeleteFromList(action)
	case "Back":
		Execute()
	}
}

func executeDeleteFromList(delete_type string) {
	options := []model.ProjectPromptItemModel{}

	var list []model.ProjectParentModel
	if delete_type == "Project" {
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

	options = append(options, model.ProjectPromptItemModel{ProjectName: "Back", ChildList: "", Action: "Back"})

	template := model.GetProjectSelectTemplate(delete_type)
	searcher := model.GetProjectSearcher(options)

	prompt := promptui.Select{
		Label:     "Which " + delete_type + " are you going to delete?",
		Items:     options,
		Templates: template,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()

	action := options[i].Action

	if action == "Back" {
		ExecuteDelete()
		return
	}

	var parentModel *model.ProjectParentModel
	for _, item := range list {
		if item.ProjectName == options[i].ProjectName {
			parentModel = &item
			break
		}
	}

	if parentModel == nil {
		color.Red("Unexpected error occurred, " + options[i].ProjectName + " not found")
		ExecuteDelete()
		return
	}

	executeDeleteChild(delete_type, parentModel)
}

func executeDeleteChild(delete_type string, parentProject *model.ProjectParentModel) {
	options := []model.SimplePromptItemModel{
		{Name: "All", Action: "All"},
	}

	for _, item := range parentProject.List {
		options = append(options, model.SimplePromptItemModel{Name: item.ProjectName, Action: ""})
	}

	options = append(options, model.SimplePromptItemModel{Name: "Back", Action: "Back"})

	template := model.GetSimpleSelectTemplate(parentProject.ProjectName)
	searcher := model.GetSimpleSearcher(options)

	prompt := promptui.Select{
		Label:     "Which " + delete_type + " do you want to delete?",
		Items:     options,
		Templates: template,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()

	action := options[i].Action

	if action == "Back" {
		executeDeleteFromList(delete_type)
		return
	}

	var projectList model.ProjectList
	var updatedList model.ProjectList
	if delete_type == "Project" {
		projectList = util.ReadProjectList()
	} else {
		projectList = util.ReadDeploymentList()
	}

	if action == "All" {
		updatedList = util.RemoveWholeParentProject(projectList, parentProject.ProjectName)
	} else {
		updatedList = util.RemoveProject(projectList, parentProject.ProjectName, options[i].Name)
	}

	if delete_type == "Project" {
		util.SaveProjectList(updatedList)
	} else {
		util.SaveDeploymentList(updatedList)
	}
}
