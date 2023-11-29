package main

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI
var chatId int64

var badWords = []string{
	"мат",
}

var answer string = "Общайтесь без мата, пожалуйста)"

func connWithTg() {
	var err error
	if bot, err = tgbotapi.NewBotAPI("<Token>"); err != nil {
		panic("cant conn with Tg")
	}
}

func isMessageBad(update *tgbotapi.Update) bool {
	if update.EditedMessage == nil {
		msgInLowerCase := strings.ToLower(update.Message.Text)
		for _, word := range badWords {
			if strings.Contains(msgInLowerCase, word) {
				return true
			}
		}
	} else {
		EditedMsgInLowerCase := strings.ToLower(update.EditedMessage.Text)
		for _, eWord := range badWords {
			if strings.Contains(EditedMsgInLowerCase, eWord) {
				return true
			}
		}

	}
	return false
}

func sendAnswer(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(chatId, answer)
	if update.EditedMessage == nil {
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	} else {
		msg.ReplyToMessageID = update.EditedMessage.MessageID
		bot.Send(msg)
	}
}

func main() {

	connWithTg()

	updateConfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.ChannelPost == nil {
			if update.EditedMessage == nil {
				if update.Message != nil || update.Message.Text == "/start" {
					chatId = update.Message.Chat.ID
				}
				if isMessageBad(&update) {
					sendAnswer(&update)
				}
			} else {
				chatId = update.EditedMessage.Chat.ID
				if isMessageBad(&update) {
					sendAnswer(&update)
				}
			}
		} else {
			continue
		}
	}
}
