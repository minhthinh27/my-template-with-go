package biz

import (
	"my-template-with-go/container"
	"my-template-with-go/internal/service"
	"my-template-with-go/logger"
)

type mailboxUC struct {
	zap       logger.ILogger
	container container.IContainerProvider
}

func (b *mailboxUC) ProcessSync() {
	//TODO implement me
	panic("implement me")
}

func NewMailBoxUC(
	zap logger.ILogger,
	container container.IContainerProvider,
) service.IMailboxUC {
	return &mailboxUC{
		zap:       zap,
		container: container,
	}
}
