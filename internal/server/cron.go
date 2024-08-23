package server

import (
	"my-template-with-go/helper/nlcron"
	"my-template-with-go/internal/cron"
)

func NewCRONServer(
	mailBoxCron cron.IMailBoxCron,
) (nlcron.ICronApp, func(), error) {
	app := nlcron.NewCronApplication()
	app.Register(mailBoxCron)

	cleanup := func() {
		app.Stop()
	}

	return app, cleanup, nil
}
