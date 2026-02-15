package telegram

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

// flushPending drains queued updates.
func (c *Client) flushPending() {
	params := url.Values{"timeout": {"0"}, "offset": {strconv.FormatInt(c.offset, 10)}}
	resp, err := c.http.Get(c.apiURL("getUpdates") + "?" + params.Encode())
	if err != nil {
		return
	}
	var result updateResponse
	json.NewDecoder(resp.Body).Decode(&result)
	resp.Body.Close()
	for _, u := range result.Result {
		c.offset = u.UpdateID + 1
	}
}

func (c *Client) PollForReply(ctx context.Context) (string, error) {
	c.flushPending()
	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
		}
		params := url.Values{"timeout": {"60"}, "offset": {strconv.FormatInt(c.offset, 10)}}
		resp, err := c.http.Get(c.apiURL("getUpdates") + "?" + params.Encode())
		if err != nil {
			continue
		}
		var result updateResponse
		json.NewDecoder(resp.Body).Decode(&result)
		resp.Body.Close()
		for _, u := range result.Result {
			c.offset = u.UpdateID + 1
			if u.Message != nil && u.Message.Chat.ID == c.chatID {
				return u.Message.Text, nil
			}
		}
	}
}
