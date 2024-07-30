package config

var config *Config

func C() *Config {
	if config == nil {
		config = ReadDBConf(Filename)
	}

	return config
}
