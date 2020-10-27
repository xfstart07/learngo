.PHONY: push

# Golang Flags
GOPATH ?= $(shell go env GOPATH)
GO=go

# vet: # run go vet
# 	@echo Run go vet
# 	go vet

push: # git push
	@echo Push
	git add .
	git commit -m 'update'
	git push origin master
