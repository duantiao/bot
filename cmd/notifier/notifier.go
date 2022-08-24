package main

import (
	"context"
	"errors"
	"github.com/duantiao/bot/pkg/lark"
	"github.com/duantiao/bot/pkg/lark/template"
	"log"
)

type Notifier struct {
	logger *log.Logger
	larkRobotClient       *lark.Client
}

func NewNotifier(logger *log.Logger,larkRobotClient *lark.Client) (*Notifier, error) {
	if logger == nil {
		return nil, errors.New("empty logger")
	}
	if larkRobotClient == nil {
		return nil, errors.New("empty larkRobotClient")
	}
	return &Notifier{
		logger: logger,
		larkRobotClient:       larkRobotClient,
	}, nil
}

func (n *Notifier) NoticeWorkflow(ctx context.Context) {
	msgSendResp, err := n.larkRobotClient.Send(ctx, lark.NewInteractiveMessage().
		SetCard(template.WorkflowCheckTmp()))
	if err != nil {
		n.logger.Printf("send lark message failed, body=%s, msgSendResp=%v, error=%s", template.WorkflowCheckTmp(), msgSendResp, err)
	}
	return
}
