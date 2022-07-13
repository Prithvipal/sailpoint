package main

import (
	"fmt"

	"github.com/Prithvipal/sailpoint/config"
	"github.com/Prithvipal/sailpoint/jobs"
)

func main() {
	cfg := config.GetConfig()

	cfg.DeclFlag()
	fmt.Println(cfg.Git.Owner)
	fmt.Println(cfg.Git.URL)
	jobs.Start()
}
