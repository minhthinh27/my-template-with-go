package main

import (
	"flag"
	"my-template-with-go/bootstrap"
	"my-template-with-go/container"
	"my-template-with-go/internal/migration"
	"my-template-with-go/internal/server"
	"my-template-with-go/logger"
)

var flagConf string

func init() {
	flag.StringVar(&flagConf, "conf", "./configs/config.yaml", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	config := bootstrap.InitConfig(flagConf)
	sugar := logger.InitLogger(config)
	provider := container.NewContainer(config, sugar)

	migration.AutoMigrate(provider.DatabaseProvider().GetDBMain(), sugar.GetZapLogger())
	router := server.InitRouter(provider, sugar, config)

	server.Start(router, sugar, config.Server)
}
