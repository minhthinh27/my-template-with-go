package bootstrap

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server Server
	Cache  Cache
	Timer  Timer
}

func InitConfig(configPath string) (Config, error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigFile(fmt.Sprintf("%s/config.yaml", configPath))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error configs file: %w \n", err))
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
