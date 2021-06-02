package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jackk01/alertmanager-dingtalk-hook/model"
	"github.com/jackk01/alertmanager-dingtalk-hook/pkg"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func httpHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprint(w, "")
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	log.Print(string(body))

	var notification model.Notification
	// json to struct
	if err := json.Unmarshal(body, &notification); err != nil {
		log.Printf("unmarshal data error, %s", err.Error())
		return
	}

	// send messages to dingtalk
	if err := pkg.Send(notification); err != nil {
		log.Printf("send messages error, %s", err.Error())
		return
	}

	fmt.Fprint(w, `{"message": "send to dingtalk successful"}`)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", httpHandle)
	server := &http.Server{
		Addr:    ":5000",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
