# Простой способ отсылки сообщений в Telegram

1. Регистрируете своего бота через Telegram `@BotFather` и получаете токен.
2. Создаете канал и добавляете в него бота.
3. Выясняете номер канала по ссылке `https://api.telegram.org/bot<YourBOTToken>/getUpdates`: идентификатор будет в виде `chat":{"id":-zzzzzzzzzz`, где `-zzzzzzzzzz` - интересующий нас номер.
4. Далее, все просто:
    ```golang
    // initializing new chat bot send function
    send := telegram.New(
        "123456789:AABB-010203040506070809", // token
        -100000000,                          // chat ID
        telegram.TypeMarkdownV2)               // message format
    // send message to telegram chat
    if err := send("*Hi!*"); err != nil {
        panic(err)
    }
    ```
