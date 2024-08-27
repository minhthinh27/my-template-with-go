package server

import (
	"my-template-with-go/bootstrap"
	"my-template-with-go/container"
	"my-template-with-go/helper/nl_cron"
	"my-template-with-go/internal/biz"
	"my-template-with-go/internal/cron"
	"my-template-with-go/internal/data"
	"my-template-with-go/logger"
)

func NewCRONServer(
	provider container.IContainerProvider,
	zap logger.ILogger,
	cf bootstrap.Config,
) (nl_cron.ICronApp, func(), error) {
	articleRepo := data.NewArticleRepo(provider.DatabaseProvider())
	articleUseCase := biz.NewArticleUseCase(articleRepo)

	articleCron, cleanup, err := cron.NewMailBoxCron(cf, zap, articleUseCase)
	if err != nil {
		cleanup()
		return nil, nil, err
	}

	app := nl_cron.NewCronApplication()
	app.Register(articleCron)

	cleanApp := func() {
		app.Stop()
	}

	return app, cleanApp, nil
}
