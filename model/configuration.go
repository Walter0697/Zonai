package model

type ProjectConfigurationModel struct {
	OutputImagePath    string
	EnviromentPath     string
	DockerBuildCommand string

	CurrentEnvironment string
}
