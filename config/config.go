package config

type PromConfig struct {
	Services []Service `yaml:"services"`
	Day      string    `yaml:"day`
}

type Service struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
}
