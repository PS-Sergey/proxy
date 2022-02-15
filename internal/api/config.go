package api

type Config struct {
	ProxyUrl  string `yaml:"proxyUrl"`
	Port      string `yaml:"port"`
	CacheSize string `yaml:"cacheSize"`
}

func NewConfig() *Config {
	return &Config{}
}
