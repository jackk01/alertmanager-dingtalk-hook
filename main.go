package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	if len(body) == 0 {
		fmt.Fprint(w, "")
		return
	}
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

func serveHTTP(srv *http.Server) {
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen error, %s", err.Error())
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", httpHandle)
	server := &http.Server{Addr: ":5000", Handler: mux}
	go serveHTTP(server)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.Shutdown(ctx)
}
