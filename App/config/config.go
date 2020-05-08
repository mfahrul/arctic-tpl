package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

//Environment environment
type Environment struct {
	Version     string `envconfig:"VERSION" default:"0.1.1"`
	DbName      string `envconfig:"DBNAME" default:"giftanoDb"`
	DbUsername  string `envconfig:"DBUSENAME" default:"root"`
	DbPassword  string `envconfig:"DBPASSWORD" default:"Giftano_Id"`
	DbHost      string `envconfig:"DBHOST" default:"localhost"`
	DbPort      string `envconfig:"DBPORT" default:"27017"`
	ServiceName string `envconfig:"SERVICENAME" default:"core"`
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
