package util

import (
	"fmt"

	"github.com/Walter0697/zonai/model"
	"github.com/jedib0t/go-pretty/table"
)

func ListHistory(list []model.DeploymentItemModel) {
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"#", "File Name", "Build Time", "Project Name", "Project Path", "Image Tag"})
	for i, item := range list {
		if len(item.ImageList) != 0 {
			firstImage := item.ImageList[0]
			tw.AppendRow([]interface{}{i + 1, item.FileName, item.BuildTime, firstImage.ProjectName, firstImage.ProjectPath, firstImage.ImageTag})
		} else {
			tw.AppendRow([]interface{}{i + 1, item.FileName, item.BuildTime, "", ""})
		}

		for index, image := range item.ImageList {
			if index == 0 {
				continue
			}
			tw.AppendRow([]interface{}{"", "", item.BuildTime, image.ProjectName, image.ProjectPath, image.ImageTag})
		}
	}

	fmt.Println(tw.Render())
}

func DisplayHistory(project_name *string) {
	historyList := ReadDeploymentHistory()
	var filteredList []model.DeploymentItemModel

	if project_name != nil {
		for _, item := range historyList.List {
			hasProject := false
			var imageList []model.DeploymentImageItem
			for _, image := range item.ImageList {
				if image.ProjectName == *project_name {
					imageList = append(imageList, image)
					hasProject = true
				}
			}
			if hasProject {
				item.ImageList = imageList
				filteredList = append(filteredList, item)
			}
		}
	} else {
		filteredList = historyList.List
	}

	Divider()
	fmt.Println("Deployment History")
	ListHistory(filteredList)
	Divider()
}
