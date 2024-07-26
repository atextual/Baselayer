package models

type Database struct {
	Id             int    `db:"id" json:"id"`
	Name           string `db:"name" json:"name"`
	NormalisedName string `db:"normalised_name" json:"normalisedName"`
	Driver         string `yaml:"driver" db:"driver" json:"driver"`
	Username       string `yaml:"username" db:"username" json:"username"`
	Password       string `yaml:"password" db:"password" json:"password"`
	Database       string `yaml:"database" db:"database,omitempty" json:"database"`
	Port           int    `yaml:"port" db:"port" json:"port"`
}
