package pkg

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"
)

func sign(ts int64, secret string) string {
	strToHash := fmt.Sprintf("%d\n%s", ts, secret)
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(strToHash))
	data := hmac256.Sum(nil)
	return base64.StdEncoding.EncodeToString(data)
}

func robotURL() string {
	token := os.Getenv("ROBOT_TOKEN")
	secret := os.Getenv("ROBOT_SECRET")
	if token == "" || secret == "" {
		log.Print("env ROBOT_TOKEN or ROBOT_SECRET not found")
		return ""
	}
	timestamp := time.Now().UnixNano() / 1e6
	sign := sign(timestamp, secret)
	return fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%d&sign=%s", token, timestamp, sign)
}
