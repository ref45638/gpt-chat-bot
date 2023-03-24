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
