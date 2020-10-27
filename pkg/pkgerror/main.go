// Author: xufei
// Date: 2020/7/23

package main

import (
	"learngo/pkg/pkgerror/repo"

	"go.uber.org/zap"
)

func main() {
	config := zap.NewDevelopmentConfig()
	build, _ := config.Build()
	logger := build.Sugar()

	if err := repo.Get(); err != nil {
		logger.Error("查询错误", err)
	}

	if err := repo.Update(); err != nil {
		logger.Error("更新错误", err)
	}
}
