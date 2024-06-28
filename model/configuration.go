package model

type ProjectConfigurationModel struct {
	OutputImagePath    string
	InputImagePath     string
	EnviromentPath     string
	DockerBuildCommand string

	CurrentEnvironment string
}
