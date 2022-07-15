package jobs

import (
	"time"

	"github.com/desertbit/timer"
	"github.com/sirupsen/logrus"
)

func Start() {
	logrus.Info("Starting Job")
	timer := timer.NewTimer(0)
	for {
		select {
		case <-timer.C:
			if err := process(); err != nil {
				logrus.Debug("error while executing job so reseting job to start after 5 min")
				timer.Reset(5 * time.Minute)
			} else {
				logrus.Debug("job executed successfully, setting next interval as 7 days")
				timer.Reset(7 * 24 * time.Hour)
			}
		}
	}
}

func process() error {
	logrus.Info("Starting Iteration")
	prs, err := GetPullRequests()
	if err != nil {
		return err
	}
	return SendMail(prs)
}
