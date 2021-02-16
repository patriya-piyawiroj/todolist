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

// type Config struct {
// 	Server struct {
// 		address string `yaml:"project_addr"`
// 		name    string `yaml:"service_name"`
// 	} `yaml:"server"`
// 	MongoServer struct {
// 		connString  string `yaml:"address"`
// 		localString string `yaml:"local"`
// 	} `yaml:"mongodb"`
// }

// func readFile(cfg *Config) {
// 	f, err := os.Open("config.yml")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer f.Close()

// 	decoder := yaml.NewDecoder(f)
// 	err = decoder.Decode(cfg)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func readEnv(cfg *Config) {
// 	err := envconfig.Process("", cfg)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
