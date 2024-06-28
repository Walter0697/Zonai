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
