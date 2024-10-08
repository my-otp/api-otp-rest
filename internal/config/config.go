package config

type Config struct {
	HttpConfig     `yaml:"http"`
	DatabaseConfig `yaml:"database"`
}

type HttpConfig struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type DatabaseConfig struct {
	Database string `yaml:"database"`
	URL      string `yaml:"url"`
}
