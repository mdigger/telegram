package telegram

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	botAPIURL = "https://api.telegram.org/"
	agent     = "mdigger-telegram-bot/1.0"
)

// Задает форматирование сообщения
const (
	TypeNone     byte = iota // простой текст
	TypeMarkdown byte = iota // в формате Markdown
	TypeHTML     byte = iota // в формате HTML
)

// Bot описывает бота для Telegram. В качестве значения задается
// токен.
type Bot string

// Send отправляет сообщение в Telegram.
func (b Bot) Send(chatID int64, text string, format byte) error {
	var params = url.Values{}
	params.Set("chat_id", strconv.FormatInt(chatID, 10))
	params.Set("text", text)
	switch format {
	case TypeMarkdown:
		params.Set("parse_mode", "Markdown")
	case TypeHTML:
		params.Set("parse_mode", "HTML")
	}
	var apiURL = botAPIURL + "bot" + string(b) + "/sendMessage"
	req, err := http.NewRequest("POST", apiURL,
		strings.NewReader(params.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", agent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	var httpClient = &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return nil
	}
	var telegramError = new(struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	})
	if err = json.NewDecoder(resp.Body).Decode(telegramError); err == nil {
		return errors.New(telegramError.Description)
	}
	return errors.New(resp.Status)
}

// NewChatBot возвращает функцию для посылки сообщений в чат Telegram.
// В качестве параметров передается токен Telegram, идентификатор чата и
// формат текста.
func NewChatBot(token string, chatID int64, format byte) func(text string) error {
	return func(text string) error {
		return Bot(token).Send(chatID, text, format)
	}
}
