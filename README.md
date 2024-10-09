# Go module made to ease handling Telegram Bot API updates

## Installation
```
go get https://github.com/ivn-ln/go-tgbot-handler
```


## Example

Example of a simple bot with a /help command and a plaintext handler.

```
package main

import (
	token "tg-bot/API"
	
	tgbothandler "https://github.com/ivn-ln/go-tgbot-handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, handler, err := tgbothandler.NewTGBot(token.TelegramBotToken, false)
	if err != nil {
		panic(err)
	}
	help := func(update tgbotapi.Update) {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Useful information")
		bot.Send(msg)
	}
	handler.Callbacks.PlaintextHandler = func(update tgbotapi.Update) {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Try /help for command list!")
		bot.Send(msg)
	}
	handler.RegisterCommand("help", help)
	handler.StartPolling(bot)
}
```