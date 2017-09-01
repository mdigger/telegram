package telegram_test

import (
	"log"

	"github.com/mdigger/telegram"
)

func Example() {
	// initializing new chat bot send function
	var chatBot = telegram.NewChatBot(
		"123456789:AABB-010203040506070809", // token
		-100000000,                          // chat ID
		telegram.TypeMarkdown)               // message format
	// send message to telegram chat
	if err := chatBot("*Hi!*"); err != nil {
		log.Fatal(err)
	}
}
