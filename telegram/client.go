package telegram

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const apiBase = "https://api.telegram.org/bot"

// Client is a minimal Telegram Bot API client.
type Client struct {
	token  string
	chatID int64
	http   *http.Client
	offset int64
}

// NewClient creates a Telegram API client.
func NewClient(token string, chatID int64) *Client {
	return &Client{
		token:  token,
		chatID: chatID,
		http:   &http.Client{},
	}
}

func (c *Client) apiURL(method string) string {
	return fmt.Sprintf("%s%s/%s", apiBase, c.token, method)
}

// SendMessage sends an HTML-formatted message to the configured chat.
func (c *Client) SendMessage(text string) error {
	params := url.Values{
		"chat_id":    {strconv.FormatInt(c.chatID, 10)},
		"text":       {text},
		"parse_mode": {"HTML"},
	}
	resp, err := c.http.PostForm(c.apiURL("sendMessage"), params)
	if err != nil {
		return fmt.Errorf("telegram send: %w", err)
	}
	resp.Body.Close()
	return nil
}
