package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Main function
func main() {

	f, err := os.Open("config.yml")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Println(err)
	}

	repo := cfg.Database

	// readFile(&cfg)
	// readEnv(&cfg)
	fmt.Printf("%+v", cfg)

	a := App{}
	a.Initialize(repo.Address, repo.DB, repo.Collection)
	a.Run(cfg.Server.Port)
}
