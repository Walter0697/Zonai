package prompt

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/Walter0697/zonai/model"
	"github.com/Walter0697/zonai/util"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func ExecuteDeploy() {
	configuration := util.ReadConfiguration()
	if configuration.InputImagePath == "" {
		command := color.YellowString("zonai deploy")
		fmt.Println("Input Image Path is not set, if you want to use this function, please set up input image path in configuration first")
		fmt.Println("Or, deploy using the following command: " + command)
		Execute()
		return
	}

	inputImagePath := configuration.InputImagePath
	// browse all files in the input image path
	files, err := os.ReadDir(inputImagePath)
	if err != nil {
		panic(err)
	}

	options := []model.DeploymentPromptItemModel{}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filename := file.Name()
		if path.Ext(filename) != ".gz" {
			continue
		}

		fileInfo := strings.Split(filename, "_")
		projectName := fileInfo[0]
		environmentInfo := fileInfo[len(fileInfo)-1]
		environment := strings.Split(environmentInfo, ".")[0]
		date := strings.Join(fileInfo[1:len(fileInfo)-1], "_")

		options = append(options, model.DeploymentPromptItemModel{
			Filename:    filename,
			ProjectName: projectName,
			Environment: environment,
			CreateDate:  date,
			LatestText:  "",
			Action:      "",
		})
	}

	// sort options with date time
	sort.Slice(options, func(i, j int) bool {
		aCreateDate, _ := time.Parse("2006-01-02_15_04_05", options[i].CreateDate)
		bCreateDate, _ := time.Parse("2006-01-02_15_04_05", options[j].CreateDate)
		return aCreateDate.After(bCreateDate)
	})

	if len(options) != 0 {
		options[0].LatestText = color.YellowString("(Latest)")
	}

	options = append(options, model.DeploymentPromptItemModel{
		Filename:    "Back",
		ProjectName: "",
		Environment: "",
		CreateDate:  "",
		LatestText:  "",
		Action:      "Back",
	})

	template := model.GetDeploymentSelectTemplate("Deploy")
	searcher := model.GetDeploymentSearcher(options)

	prompt := promptui.Select{
		Label:     "Which project are you going to deploy?",
		Items:     options,
		Templates: template,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, _ := prompt.Run()

	action := options[i].Action
	if action == "Back" {
		Execute()
		return
	}

	filename := options[i].Filename
	filePath := path.Join(inputImagePath, filename)

	now := time.Now().Format("2006-01-02 15:04:05")
	var deploymentItem model.DeploymentItemModel
	deploymentItem.FileName = filename
	deploymentItem.BuildTime = now
	deploymentItem.ImageList = []model.DeploymentImageItem{}

	loadedImageList := util.LoadAllImagesFromGz(filePath, inputImagePath)
	pathList := []string{}
	for _, imageTag := range loadedImageList {
		destination, imageItem := util.FindComposeAndEdit(imageTag)
		if destination != "" {
			pathList = append(pathList, destination)
		}

		if imageItem != nil {
			deploymentItem.ImageList = append(deploymentItem.ImageList, *imageItem)
		}
	}

	deploymentHistory := util.ReadDeploymentHistory()
	deploymentHistory.List = append(deploymentHistory.List, deploymentItem)
	util.SaveDeploymentHistory(deploymentHistory)

	deployInstruction := color.YellowString("docker-compose up -d")
	fmt.Println("--> To deploy, run " + deployInstruction + " in the following paths")
	for _, path := range pathList {
		color.Green("cd " + path)
	}

}
