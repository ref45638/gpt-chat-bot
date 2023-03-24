package main

import (
	"context"
	"fmt"
	"log"

	"github.com/PullRequestInc/go-gpt3"
)

type Chat struct {
	ctx       context.Context
	client    gpt3.Client
	messsages []gpt3.ChatCompletionRequestMessage

	listener chan string
	stop     chan bool
}

func (c *Chat) OpenListener() (chan string, chan bool) {
	c.listener = make(chan string)
	c.stop = make(chan bool)

	return c.listener, c.stop
}

func (c *Chat) Chat(user string, message string) {
	c.messsages = append(c.messsages, gpt3.ChatCompletionRequestMessage{Role: user, Content: message})
	go c.chatCompletionStream()
}

func (c *Chat) NewChat() {
	c.messsages = []gpt3.ChatCompletionRequestMessage{}
}

func (c *Chat) chatCompletionStream() {
	fmt.Println(c.messsages[len(c.messsages)-1].Content)

	var answer string

	err := c.client.ChatCompletionStream(c.ctx, gpt3.ChatCompletionRequest{
		Model:     gpt3.GPT3Dot5Turbo,
		Messages:  c.messsages,
		MaxTokens: 2000,
	}, func(resp *gpt3.ChatCompletionStreamResponse) {
		content := resp.Choices[0].Delta.Content

		fmt.Print(content)
		c.listener <- content
		answer += content

		if resp.Choices[0].FinishReason == "stop" {
			fmt.Println()
			c.stop <- true
			c.messsages = append(c.messsages, gpt3.ChatCompletionRequestMessage{Role: "assistant", Content: answer})
		}
	})
	if err != nil {
		log.Fatalln(err)
	}
}
