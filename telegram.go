// Package telegram предоставляет простой способ отправки сообщений в Telegram.
//
// Для получения токена необходимо воспользоваться роботом Telegram @BotFather.
// Там можно в интерактивном режиме создать своего бота и получить токен для
// его использования.
//
// Далее, создаете свой канал в Telegram и, используя его номер, отправляете
// в него сообщения. Для получения номера канала воспользуйтесь одним из
// следующих советов:
// https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id/38388851#38388851
package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	botAPIURL = "https://api.telegram.org/bot%s/sendMessage"
	agent     = "mdigger-telegram-bot/1.0"
)

// Client is used to send HTTP messages.
var Client = &http.Client{
	Timeout: time.Second * 5,
}

// Задает форматирование сообщения
const (
	TypeNone     byte = iota // простой текст
	TypeMarkdown             // в формате Markdown
	TypeHTML                 // в формате HTML
)

// Bot описывает бота для Telegram. В качестве значения задается
// токен.
type Bot string

// Send отправляет сообщение в Telegram.
func (b Bot) Send(chatID int64, text string, format byte) error {
	params := make(url.Values, 3)
	params.Set("chat_id", strconv.FormatInt(chatID, 10))
	params.Set("text", text)
	switch format {
	case TypeMarkdown:
		params.Set("parse_mode", "Markdown")
	case TypeHTML:
		params.Set("parse_mode", "HTML")
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(botAPIURL, b),
		strings.NewReader(params.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", agent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil
	}

	var telegramError struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&telegramError); err == nil {
		return errors.New(telegramError.Description)
	}

	return errors.New(resp.Status)
}

// New возвращает функцию для посылки сообщений в чат Telegram.
// В качестве параметров передается токен Telegram, идентификатор чата и
// формат текста.
func New(token string, chatID int64, format byte) func(text string) error {
	return func(text string) error {
		return Bot(token).Send(chatID, text, format)
	}
}
