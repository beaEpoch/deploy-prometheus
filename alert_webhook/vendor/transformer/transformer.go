package transformer

import (
	"bytes"
	"fmt"
	"model"
	"time"
)

// TransformToMarkdown transform alertmanager notification to dingtalk markdow message
func TransformToMarkdown(notification model.Notification) (markdown *model.Markdown, robotURL string, err error) {

	// groupKey := notification.GroupKey
	// status := notification.Status

	annotations := notification.CommonAnnotations
	robotURL = annotations["robot"]

	var buffer bytes.Buffer

	for _, alert := range notification.Alerts {
		annotations := alert.Annotations
		buffer.WriteString(fmt.Sprintf("[%s]\n", time.Now().Format("20060102-15:04:05")))
		buffer.WriteString(fmt.Sprintf("%s\n%s\n", annotations["summary"], annotations["description"]))
	}

	markdown = &model.Markdown{
		Title: "prometheus",
		Text:  buffer.String(),
	}

	return
}
