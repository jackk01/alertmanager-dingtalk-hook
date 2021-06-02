package pkg

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jackk01/alertmanager-dingtalk-hook/model"
)

func Send(notification model.Notification) (err error) {
	markdown, webhook, err := transformToMarkdown(notification)
	if err != nil {
		return
	}

	data, err := json.Marshal(markdown)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("create an new request error, %s", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Print("send an http request error, please try again.")
		return
	}
	defer resp.Body.Close()

	return
}
