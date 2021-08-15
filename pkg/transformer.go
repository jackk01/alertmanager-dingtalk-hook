package pkg

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/jackk01/alertmanager-dingtalk-hook/model"
)

func transformToMarkdown(notification model.Notification) (markdown *model.DingTalkMarkdown, webhook string, err error) {
	webhook = robotURL()
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("### [%s:%d]\n", strings.ToUpper(notification.Status), len(notification.Alerts)))
	buffer.WriteString(fmt.Sprintf("---\n"))

	for _, alert := range notification.Alerts {
		annotations := alert.Annotations
		buffer.WriteString(fmt.Sprintf("##### %s\n > 故障描述: %s\n", annotations["summary"], annotations["description"]))
		buffer.WriteString(fmt.Sprintf("\n> 故障等级：%s\n", alert.Labels["severity"]))
		buffer.WriteString(fmt.Sprintf("\n> 当前状态：%s\n", alert.Status))
		buffer.WriteString(fmt.Sprintf("\n> 开始时间：%s\n", alert.StartsAt.Local().Format("2006-01-02 15:04:05")))
	}

	markdown = &model.DingTalkMarkdown{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Title: fmt.Sprintf("您有%d条监控信息, 请及时查看", len(notification.Alerts)),
			Text:  buffer.String(),
		},
		At: &model.At{
			IsAtAll: false,
		},
	}

	return
}
