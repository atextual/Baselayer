package models

type Projects map[string]*Project

type Project struct {
	Name             string
	ProjectDirectory string   `yaml:"projectDir"`
	SqlDirectory     string   `yaml:"sqlDir"`
	Database         Database `yaml:"database"`
}
