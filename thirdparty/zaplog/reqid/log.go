package main

import "go.uber.org/zap"

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	logger.With(zap.String("reqid", "xxxxx")).Info("kk")
}
