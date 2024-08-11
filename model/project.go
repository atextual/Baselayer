package model

type Projects map[string]*Project

type Project struct {
	Name             string   `yaml:"name" json:"name"`
	ProjectDirectory string   `yaml:"projectDir" json:"projectDir"`
	SqlDirectory     string   `yaml:"sqlDir" json:"sqlDir"`
	Database         Database `yaml:"database" json:"database"`
}
