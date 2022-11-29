package tg

type TelegramAPI struct {
	Ok     bool             `json:"ok"`
	Result []TelegramResult `json:"result"`
}

type TelegramResult struct {
	UpdateId int         `json:"update_id"`
	Message  MessageData `json:"message"`
}

type MessageData struct {
	MessageId int         `json:"message_id"`
	From      MessageFrom `json:"from"`
	Chat      Chatting    `json:"chat"`
	Date      int         `json:"date"`
	Text      string      `json:"text"`
}

type MessageFrom struct {
	Id           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type Chatting struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}
