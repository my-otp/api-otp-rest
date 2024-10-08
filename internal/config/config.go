package config

type Config struct {
	HttpConfig `yaml:"http"`
}

type HttpConfig struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}
