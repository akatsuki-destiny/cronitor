package main

import (
	"cronitor/pkg/config"
	"cronitor/pkg/scripts"
	"cronitor/pkg/utils"
	"github.com/robfig/cron/v3"
)

func main() {

	// Init env configs
	config.InitEnvConfigs()

	// create cron
	c := cron.New()

	// Add reset login attempt counts cron job
	_, err := c.AddFunc("0 */2 * * *", scripts.ResetLoginAttemptCounts)
	utils.FatalErr("Error while adding cron job", err)

	c.Start()

	// Wait until the application is stopped
	select {}
}
