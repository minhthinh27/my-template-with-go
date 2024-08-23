package nlcron

import (
	"context"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"strings"
	"time"
)

type ICronJob interface {
	Start()
	Stop() context.Context
	Log(error)
}

type job struct {
	name     string
	callback func()
	schedule string
	sugar    *zap.SugaredLogger
	cron     *cron.Cron
}

func NewCronJob(name string, timer string, timezone string, callback func(), sugar *zap.SugaredLogger) ICronJob {
	loc := time.Now().Location()
	if zone, err := time.LoadLocation(timezone); err == nil {
		loc = zone
	} else {
		sugar.Error(err)
	}

	cron := cron.New(
		cron.WithSeconds(),
		cron.WithLocation(loc),
	)

	callLog := func() {
		sugar.Infof("======= Job %s starting invoke at %s =======", name, time.Now().String())
		callback()
		sugar.Infof("======= Job %s completed invoke at %s =======", name, time.Now().String())
	}

	timer = formatTimer(timer)
	_, err := cron.AddFunc(timer, callLog)
	if err != nil {
		sugar.Error(err)
		return nil
	}

	return &job{
		name:     name,
		schedule: timer,
		callback: callLog,
		sugar:    sugar,
		cron:     cron,
	}
}

func (j *job) Start() {
	if j.cron != nil {
		j.cron.Start()
	}
}

func (j *job) Stop() context.Context {
	if j.cron != nil {
		return j.cron.Stop()
	}

	return nil
}

func (j *job) Log(err error) {
	if j.sugar != nil {
		j.sugar.Error(err)
	}
}

func formatTimer(timer string) string {
	values := strings.Split(strings.Trim(timer, " "), " ")
	for len(values) < 6 {
		values = append([]string{"0"}, values...)
	}

	return strings.Join(values, " ")
}
