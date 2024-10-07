package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

const addr = "quote.wfgroup.com.hk:8083"

func main() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	u := url.URL{Scheme: "wss", Host: addr, Path: "/socket.io/", RawQuery: "token=applepieapplepieapplepieapplepie&EIO=3&transport=websocket"}
	log.Printf("connecting to %s", u.String())
	websocket.DefaultDialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	header := http.Header{}
	header.Add("Host", "quote.wfgroup.com.hk:8083")
	header.Add("Pragma", "no-cache")
	header.Add("Cache-Control", "no-cache")
	header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/")
	header.Add("Origin", "https://www.wfbullion.com")
	header.Add("Accept-Encoding", "gzip, deflate, br, zstd")
	header.Add("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")

	c, rsp, err := websocket.DefaultDialer.Dial(u.String(), header)
	fmt.Println(err)
	bodyBytes, err := io.ReadAll(rsp.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal("dial:-", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				err := c.WriteMessage(websocket.TextMessage, []byte("40/bquote,"))
				if err != nil {
					log.Println("write:", err)
					return
				}
			case <-interrupt:
				log.Println("interrupt")

				// Cleanly close the connection by sending a close message and then
				// waiting (with timeout) for the server to close the connection.
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					log.Println("write close:", err)
					return
				}
				select {
				case <-done:
				case <-time.After(time.Second):
				}
				return
			}
		}

	}()

	defer close(done)
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		msg := string(message)
		if strings.HasPrefix(msg, "42/bquote,") {
			body := strings.TrimPrefix(msg, "42/bquote,")
			asArray := []interface{}{}
			er := json.Unmarshal([]byte(body), &asArray)
			if er != nil {
				fmt.Println(er)
				continue
			}
			if len(asArray) == 2 {
				asMap := asArray[1].(map[string]interface{})
				products := asMap["products"].(map[string]interface{})
				hkg := products["HKG="].(map[string]interface{})
				for k, v := range hkg {
					fmt.Println(k, v)
				}
			}
		}
	}
}
