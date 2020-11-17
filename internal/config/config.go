package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type DBconfig struct {
	Driver string `yaml:"driver"`
}

type Config struct {
	DB      DBconfig `yaml:"db"`
	Version string   `yaml:"version"`
}

func LoadConfig(filename string) (*Config, error) {
	//leo el archivo
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	//contenedo de configuracion
	var c = &Config{}
	//parseo el c en gile
	yaml.Unmarshal(file, c)
	if err != nil {
		return nil, err
	}

	return c, nil

}
