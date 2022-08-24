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
    "template": "yellow",
    "title": {
      "tag": "plain_text",
      "content": "上线后出现问题的爬虫网站"
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
            "content": "1. [Aldo](https://dataflow.voiladev.xyz/#/crawl/tasks?page=1&count=50&siteId=740a70f485afb9d035f01c0e86be529c&triger=TrigerTypeTimer)"
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
