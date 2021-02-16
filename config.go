package main

type Config struct {
	Server struct {
		Address string `yaml:"project_addr"`
		Name    string `yaml:"service_name"`
		Port    string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Address    string `yaml:"address"`
		DB         string `yaml:"db"`
		Collection string `yaml:"collection"`
	} `yaml:"mongodb"`
}
