package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var chat *Chat

func main() {
	chat = NewChat()

	router := gin.Default()
	router.POST("/chat", stream)
	router.POST("/new", func(c *gin.Context) {
		chat.NewChat()
		c.JSON(200, gin.H{
			"success": true,
		})
	})
	router.Run(":8080")
}

func NewChat() *Chat {
	godotenv.Load()

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("Missing API KEY")
	}

	c := &Chat{
		ctx: context.Background(),
		client: gpt3.NewClient(apiKey, gpt3.WithHTTPClient(
			&http.Client{
				Timeout: time.Duration(10000 * time.Second),
			},
		)),
	}

	return c
}

func stream(c *gin.Context) {
	prompt := c.PostForm("prompt")

	listener, stop := chat.OpenListener()
	chat.Chat("user", prompt)

	clientGone := c.Request.Context().Done()
	c.Stream(func(w io.Writer) bool {
		select {
		case message := <-listener:
			c.SSEvent("message", message)
			return true
		case <-stop:
			return false
		case <-clientGone:
			return false
		}
	})
}

// 請用 html,css與JS寫一個律師事務所網頁，使用Bootstrap 排版，包括一個sliders放上律師的大頭貼以及姓名，公司是睿見法律事務所， 免費法律諮詢專線：  免費市話直撥：0809-080-229  手機快速直撥：0906-898-660  LINE官方帳號：@SNE7146R， 客服時間(週一至週五)：  09:00～12:00 / 13:30～21:00  夜間/假日：0906-898-660

// 				以下是我提供的律師資訊，包括大頭貼以及姓名
// 				律師1:
// 				姓名: 陳儀文 律師
// 				大頭貼: https://reurl.cc/gZGjrN
// 				律師2:
// 				姓名: 張育嘉 律師
// 				大頭貼: https://reurl.cc/qkVjaN
// 				律師3:
// 				姓名: 黃瓊瑩 律師
// 				大頭貼: https://reurl.cc/5MvyXz
// 				律師4:
// 				姓名: 陳軾霖 律師
// 				大頭貼: https://reurl.cc/b7VjeM
// 				律師5:
// 				姓名: 趙若竹 律師
// 				大頭貼: https://reurl.cc/zA1jo6
