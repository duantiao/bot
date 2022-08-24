package main

import (
	"context"
	"log"
	"testing"
	"github.com/duantiao/bot/pkg/lark"
)

func TestNewNotifier(t *testing.T) {
	var (
		larkRobotClient *lark.Client
	)
	larkRobotClient = lark.NewClient(
		"e355f7e9-a876-4361-99d4-a22782644626",
		"Hb3RDd6dYjbXWudnZ2CyNd",
	)

	// new notifier
	notifier, err := NewNotifier(log.Default(),larkRobotClient)
	if err != nil {
		return
	}

	notifier.NoticeWorkflow(context.Background())
}
