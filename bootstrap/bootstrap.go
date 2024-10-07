package bootstrap

import (
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   Server
	Database Database
	Cache    Cache
	Timer    Timer
}

var (
	once           sync.Once
	configInstance Config
)

func InitConfig(configPath string) Config {
	once.Do(func() {
		viper.SetConfigType("yaml")
		viper.AddConfigPath(configPath)
		viper.SetConfigFile(configPath)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			log.Fatalf("Error parsing config file: %s", err)
		}

		setTimeZone(configInstance.Timer.Zone)
	})

	return configInstance
}

func setTimeZone(timeZone string) {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		log.Fatalf("Error loading timezone, %s", err)
	}

	time.Local = loc
}
