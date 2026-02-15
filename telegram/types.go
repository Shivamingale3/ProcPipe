package telegram

type tgUpdate struct {
	UpdateID int64      `json:"update_id"`
	Message  *tgMessage `json:"message"`
}

type tgMessage struct {
	Text string `json:"text"`
	Chat struct {
		ID int64 `json:"id"`
	} `json:"chat"`
}

type updateResponse struct {
	OK     bool       `json:"ok"`
	Result []tgUpdate `json:"result"`
}
