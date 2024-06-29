package prompt

import (
	"fmt"
	"strings"
	"time"

	"github.com/Walter0697/zonai/model"
	"github.com/Walter0697/zonai/util"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
	"github.com/manifoldco/promptui"
)

func ExecuteBuild() {
	projectList := util.ReadProjectList()
	options := []model.SimplePromptItemModel{}
	for _, project := range projectList.List {
		options = append(options, model.SimplePromptItemModel{Name: project.ProjectName, Action: "Build"})
	}
	options = append(options, model.SimplePromptItemModel{Name: "Back", Action: "Back"})

	searcher := model.GetSimpleSearcher(options)
	templates := model.GetSimpleSelectTemplate("Build")

	prompt := promptui.Select{
		Label:     "Which project are you going to build?",
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
	for _, project := range projectList.List {
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

	executeChildBuild(currentProject)
}

func executeChildBuild(project *model.ProjectParentModel) {
	fmt.Println("Current Child Projects available:")
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"#", "Child Name", "Flag", "Project Path"})
	for i, child := range project.List {
		tw.AppendRow([]interface{}{i + 1, child.ProjectName, child.Flag, child.ProjectPath})
	}
	util.Divider()
	fmt.Println(tw.Render())
	util.Divider()

	options := []model.SimplePromptItemModel{
		{Name: "All", Action: "All"},
		{Name: "Selected", Action: "Selected"},
		{Name: "Back", Action: "Back"},
	}
	template := model.GetSimpleSelectTemplate("Build")
	searcher := model.GetSimpleSearcher(options)

	prompt := promptui.Select{
		Label:     "Which child project are you going to build?",
		Items:     options,
		Templates: template,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()

	action := options[i].Action
	switch action {
	case "All":
		buildFlags := util.GetAllProjectFlags(project)
		executeCompress(project, buildFlags)
	case "Selected":
		executeSelectBuild(project)
	case "Back":
		ExecuteBuild()
	}
}

func executeSelectBuild(project *model.ProjectParentModel) {
	fmt.Println("To select a project, please type their flags and use space as separator")
	var flagExp []string
	maxLen := 3
	if len(project.List) < 3 {
		maxLen = len(project.List)
	}
	for i := 0; i < maxLen; i++ {
		flagExp = append(flagExp, project.List[i].Flag)
	}
	flagExpStr := strings.Join(flagExp, " ")
	fmt.Println("For example: " + flagExpStr)
	prompt := promptui.Prompt{
		Label: "Flags",
	}

	result, _ := prompt.Run()

	selectedFlags := strings.Split(result, " ")
	for _, flag := range selectedFlags {
		found := false
		for _, project := range project.List {
			if project.Flag == flag {
				found = true
				break
			}
		}
		if !found {
			color.Red("Flag " + flag + " is not a valid flag")
			ExecuteBuild()
			return
		}
	}

	executeCompress(project, selectedFlags)
}

func executeCompress(project *model.ProjectParentModel, flags []string) {
	options := []model.SimplePromptItemModel{
		{Name: "Compress to .gz", Action: "Compress"},
		{Name: "Do not compress", Action: "Not Compress"},
		{Name: "Back", Action: "Back"},
	}
	template := model.GetSimpleSelectTemplate("Compression")
	searcher := model.GetSimpleSearcher(options)

	prompt := promptui.Select{
		Label:     "Do you want to compress the Docker Images?",
		Items:     options,
		Templates: template,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()

	action := options[i].Action

	switch action {
	case "Compress":
		executeEnvironment(project, flags, true)
	case "Not Compress":
		executeEnvironment(project, flags, false)
	case "Back":
		executeChildBuild(project)
	}
}

func executeEnvironment(project *model.ProjectParentModel, flags []string, compressFlag bool) {
	configuration := util.ReadConfiguration()
	currentEnvironment := configuration.CurrentEnvironment
	envDisplay := color.YellowString("(" + currentEnvironment + ")")
	options := []model.SimplePromptItemModel{
		{Name: "Build with all environments", Action: "All"},
		{Name: "Build with current environment " + envDisplay, Action: "Current"},
		{Name: "Back", Action: "Back"},
	}

	template := model.GetSimpleSelectTemplate("environment")
	searcher := model.GetSimpleSearcher(options)

	prompt := promptui.Select{
		Label:     "What environment do you want to build with?",
		Items:     options,
		Templates: template,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()

	action := options[i].Action

	switch action {
	case "All":
		buildProject(project, flags, compressFlag, true)
	case "Current":
		buildProject(project, flags, compressFlag, false)
	case "Back":
		executeCompress(project, flags)
	}
}

func buildProject(project *model.ProjectParentModel, flags []string, compressFlag bool, wholeFlag bool) {
	configuration := util.ReadConfiguration()
	history := util.ReadBuildHistory()
	now := time.Now().Format("2006-01-02")

	if wholeFlag {
		env_list := util.GetAllEnvironments(&configuration, project)
		for _, env := range env_list {
			util.BuildProjectWithImageList(project, flags, &configuration, &history, now, compressFlag, env)
		}
	} else {
		currentEnvironment := configuration.CurrentEnvironment
		util.BuildProjectWithImageList(project, flags, &configuration, &history, now, compressFlag, currentEnvironment)
	}

	util.SaveBuildHistory(history)
}
