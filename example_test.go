package telegram_test

import (
	"github.com/mdigger/telegram"
)

func Example() {
	// initializing new chat bot send function
	send := telegram.New(
		"123456789:AABB-010203040506070809", // token
		-100000000,                          // chat ID
		telegram.TypeMarkdownV2)             // message format

	// send message to telegram chat
	if err := send("*Hi!*"); err != nil {
		panic(err)
	}
}
