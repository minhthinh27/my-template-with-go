package cron

import (
	"my-template-with-go/bootstrap"
	"my-template-with-go/helper/nl_cron"
	"my-template-with-go/internal/biz"
	"my-template-with-go/logger"
)

type IArticleCron interface {
	nl_cron.ICronJob
}

func NewMailBoxCron(
	cf bootstrap.Config,
	zap logger.ILogger,
	articleSync biz.IArticleUC,
) (IArticleCron, func(), error) {
	var (
		zone  = cf.Timer.Zone
		sugar = zap.GetZapLogger()
	)

	callback := func() { articleSync.Sync() }

	result := nl_cron.NewCronJob("article", "*/10 * * * * *", zone, callback, sugar)
	return result, func() {
		sugar.Info("closing the document nl_cron job")
		result.Stop()
	}, nil
}
