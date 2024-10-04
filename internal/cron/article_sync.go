package cron

import (
	"context"
	"my-template-with-go/bootstrap"
	"my-template-with-go/helper/gdcron"
	"my-template-with-go/internal/biz"
	"my-template-with-go/logger"
	"time"
)

type IArticleCron interface {
	gdcron.ICronJob
}

func NewMailBoxCron(cf bootstrap.Config, zap logger.ILogger, articleSync biz.IArticleUC) (IArticleCron, func(), error) {
	var (
		zone  = cf.Timer.Zone
		sugar = zap.GetZapLogger()
	)

	callback := func(ctx context.Context) error {
		return articleSync.Sync(ctx)
	}

	result := gdcron.NewCronJob("article", "*/10 * * * * *", zone, callback, sugar, time.Minute)
	return result, func() {
		sugar.Info("closing the document nl_cron job")
		result.Stop()
	}, nil
}
