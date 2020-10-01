package configs

import(
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Template string `yaml:"template"`
	DefaultName string `yaml:"defaultName"`
	Version string `yaml:"version"`
}

var Configuration Config = Config{
	Template: "Hello, %s!",
	DefaultName: "Stranger",
	Version: "Unknown",
}

func Configure(configurationFile string) error {
	file, err := os.Open(configurationFile)
	if err != nil {
    	return err;
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&Configuration)
	return err
}