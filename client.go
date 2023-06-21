package slack

import (
	"bytes"
	"fmt"
	"net/http"
)

type Client struct {
	webhookURL string
}

func NewClient(webhookURL string) *Client {
	return &Client{
		webhookURL: webhookURL,
	}
}

func (c *Client) Send(msg Message) error {
	if msg.Footer != "" {
		msg = msg.AddBlock(NewDividerBlock()).AddBlock(NewContext().AddElement(NewMrkdwnText(msg.Footer)))
	}
	if err := send(c.webhookURL, msg.String()); err != nil {
		return err
	}
	return nil
}

func send(endpoint string, msg string) error {
	b := bytes.NewBuffer([]byte(msg))
	res, err := http.Post(endpoint, "application/json", b)
	if err != nil {
		return err
	}

	if res.StatusCode >= 299 {
		return fmt.Errorf("Error on message: %s\n", res.Status)
	}
	return nil
}
