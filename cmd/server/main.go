package main

import (
	"log"
	"my-template-with-go/bootstrap"
	"my-template-with-go/container"
	"my-template-with-go/logger"
)

var configPath = "./configs"

func main() {
	config, err := bootstrap.InitConfig(configPath)
	if err != nil {
		log.Fatalln(err)
	}

	sugar, err := logger.InitLogger(config)
	if err != nil {
		sugar.GetZapLogger().Fatal(err)
	}

	_, err = container.NewContainer(config, sugar)
	if err != nil {
		sugar.GetZapLogger().Fatal(err)
	}

}

//func run(engine *echo.Echo, sugarLogger logger.ILogger, config bootstrap.Server) {
//	var (
//		sugarLog = sugarLogger.GetZapLogger()
//		timeOut  = time.Duration(config.Http.Timeout) * time.Second
//		address  = config.Http.Address
//	)
//
//	// start and wait for stop signal
//	ctx, stop := signal.NotifyContext(
//		context.Background(),
//		os.Interrupt,
//		syscall.SIGINT,
//		syscall.SIGTERM,
//		syscall.SIGQUIT,
//	)
//	defer stop()
//
//	srv := &http.Server{
//		Addr:              fmt.Sprintf(":%s", address),
//		Handler:           engine,
//		ReadHeaderTimeout: timeOut,
//		ReadTimeout:       timeOut,
//		WriteTimeout:      timeOut,
//	}
//
//	go func() {
//		sugarLog.Infof("start server on: %s", srv.Addr)
//		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
//			sugarLog.Fatal(err)
//		}
//	}()
//	<-ctx.Done()
//	stop()
//
//	var cancel context.CancelFunc
//	ctx, cancel = context.WithTimeout(context.Background(), timeOut)
//	defer cancel()
//	if err := srv.Shutdown(ctx); err != nil {
//		sugarLog.Fatal(err)
//	}
//}
