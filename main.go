package main

import (
	"fmt"
	"os"
	"todolist/app"
	"todolist/configs"

	"gopkg.in/yaml.v3"
)

// Main function
func main() {

	// ===================== init configs =============================
	f, err := os.Open("configs/config.yml")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	var conf configs.Configs
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println(err)
	}

	repo := conf.Database

	// readFile(&cfg)
	// readEnv(&cfg)
	fmt.Printf("%+v", conf)

	a := app.App{}
	a.Initialize(repo.Address, repo.DB, repo.Collection, &conf)
	a.Run(conf.Server.Port)
}
