package model

import (
	"strings"

	"github.com/manifoldco/promptui"
)

type PromptItemModel struct {
	Name        string
	Description string
}

type SimplePromptItemModel struct {
	Name   string
	Action string
}

type DeploymentPromptItemModel struct {
	Filename    string
	ProjectName string
	CreateDate  string
	Environment string
	LatestText  string
	Action      string
}

func GetSimpleSearcher(options []SimplePromptItemModel) func(string, int) bool {
	return func(input string, index int) bool {
		option := options[index]
		name := strings.Replace(strings.ToLower(option.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}
}

func GetSimpleSelectTemplate(title string) *promptui.SelectTemplates {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "-> {{ .Name | cyan }}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: title + ": Selected {{ .Name | cyan }}",
	}

	return templates
}

func GetSearcher(options []PromptItemModel) func(string, int) bool {
	return func(input string, index int) bool {
		option := options[index]
		name := strings.Replace(strings.ToLower(option.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}
}

func GetSelectTemplate(title string) *promptui.SelectTemplates {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "-> {{ .Name | cyan }}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: title + ": Selected {{ .Name | cyan }}",
		Details: `
--------- Description ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Description:" | faint }}	{{ .Description }}`,
	}

	return templates
}

func GetDeploymentSearcher(options []DeploymentPromptItemModel) func(string, int) bool {
	return func(input string, index int) bool {
		option := options[index]
		name := strings.Replace(strings.ToLower(option.ProjectName), " ", "", -1)
		env := strings.Replace(strings.ToLower(option.Environment), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input) || strings.Contains(env, input)
	}
}

func GetDeploymentSelectTemplate(title string) *promptui.SelectTemplates {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "-> {{ .Filename | cyan }} {{ .LatestText | yellow }}",
		Inactive: "  {{ .Filename | cyan }} {{ .LatestText | yellow }}",
		Selected: title + ": Selected {{ .Filename | cyan }}",
		Details: `
--------- Description ----------
{{ "Filename:" | faint }}	{{ .Filename }}
{{ "Project Name:" | faint }}	{{ .ProjectName }}
{{ "Environment:" | faint }}	{{ .Environment }}
{{ "Create Date:" | faint }}	{{ .CreateDate }}`,
	}

	return templates
}
