package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

//Environment environment
type Environment struct {
	Version     string `envconfig:"VERSION" default:"0.1.1"`
	DbName      string `envconfig:"DBNAME" default:"{{.Dbname}}"`
	DbUsername  string `envconfig:"DBUSENAME" default:"{{.Dbusername}}"`
	DbPassword  string `envconfig:"DBPASSWORD" default:"{{.Dbpassword}}"`
	DbHost      string `envconfig:"DBHOST" default:"{{.Dbhost}}"`
	DbPort      string `envconfig:"DBPORT" default:"{{.Dbport}}"`
	ServiceName string `envconfig:"SERVICENAME" default:"{{.Projectname}}"`
}

//NewConfig function
func NewConfig() Environment {
	var e Environment
	err := envconfig.Process("myapp", &e)
	if err != nil {
		log.Fatal(err.Error())
	}
	return e
}
