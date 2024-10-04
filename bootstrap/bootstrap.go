package bootstrap

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Server   Server
	Database Database
	Cache    Cache
	Timer    Timer
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

	if err := setTimeZone(config.Timer.Zone); err != nil {
		return Config{}, err
	}

	return config, nil
}

func setTimeZone(timeZone string) error {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return errors.New("load time zone failed: " + err.Error())
	}

	time.Local = loc
	return nil
}
