package cron

import (
	"fmt"
	"my-template-with-go/bootstrap"
	"my-template-with-go/helper/nlcron"
	"my-template-with-go/logger"
)

type IMailBoxCron interface {
	nlcron.ICronJob
}

func NewMailBoxCron(
	cf bootstrap.Config,
	zap logger.ILogger,
) (IMailBoxCron, func(), error) {
	var (
		zone  = cf.Timer.GetZone()
		sugar = zap.GetZapLogger()
	)

	callback := func() {
		fmt.Println("run")
	}

	result := nlcron.NewCronJob("mailbox", "*/1 * * * *", zone, callback, sugar)
	return result, func() {
		sugar.Info("closing the document nlcron job")
		result.Stop()
	}, nil
}
