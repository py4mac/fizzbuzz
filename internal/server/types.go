package server

import (
	log "github.com/sirupsen/logrus"
)

var (
	logger = log.WithFields(log.Fields{
		"module": "server",
	})
)
