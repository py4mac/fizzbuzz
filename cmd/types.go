package cmd

import (
	log "github.com/sirupsen/logrus"
)

func init() {
}

var (
	logger = log.WithFields(log.Fields{
		"module": "cmd",
	})
)
