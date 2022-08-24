package template

var (
	WorkflowCheckTmp = func() string {
		source := `{
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
}`
		return source
	}
)
