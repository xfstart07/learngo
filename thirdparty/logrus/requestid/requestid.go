package main

import
(
	log "github.com/sirupsen/logrus")

func main() {
	logw := log.WithFields(log.Fields{
		"requestid": "walrus",
	})

	logw.Info("kk", "kkk")
}