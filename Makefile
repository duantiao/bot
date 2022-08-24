.PHONY: all

all: app

app:
	go vet ./cmd/notifier
	go build -ldflags "-X main.buildTime=`date +%Y%m%d.%H:%M:%S` -X main.buildCommit=`git rev-parse --short=12 HEAD` -X main.buildBranch=`git branch --show-current`" -o ./releases/notifier ./cmd/notifier