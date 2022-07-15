package main

import (
	"github.com/Prithvipal/sailpoint/config"
	"github.com/Prithvipal/sailpoint/jobs"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.GetConfig()
	cfg.DeclFlag()
	logrus.SetLevel(3)
	jobs.Start()
}
