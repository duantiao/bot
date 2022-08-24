package lark

import (
	"context"
	"testing"
)

func TestClient_Send(t *testing.T) {
	client := NewClient(
		"e355f7e9-a876-4361-99d4-a22782644626",
		"Hb3RDd6dYjbXWudnZ2CyNd",
	)
	resp, err := client.Send(context.Background(), NewInteractiveMessage().SetCard(`{
  "config": {
    "wide_screen_mode": true
  },
  "header": {
    "template": "red",
    "title": {
      "tag": "plain_text",
      "content": "请注意查看工作任务"
    }
  },
  "elements": [
    {
      "tag": "div",
      "fields": [
        {
          "is_short": true,
          "text": {
            "tag": "lark_md",
            "content": "1. [Jira](https://jira.voiladev.xyz/secure/Dashboard.jspa)"
          }
        }
      ]
    }
  ]
}`))
	if err != nil {
		t.Errorf("resp=%v, err=%s", resp, err)
	}
}
