package commands

import (
	"fmt"
	"github.com/NexonSU/telegram-go-chatbot/app/utils"
	tb "gopkg.in/tucnak/telebot.v2"
	"strings"
)

//Delete Get in DB on /del
func Del(m *tb.Message) {
	if !utils.IsAdminOrModer(m.Sender.Username) {
		if m.Chat.Username != utils.Config.Telegram.Chat {
			return
		}
		_, err := utils.Bot.Reply(m, &tb.Animation{File: tb.File{FileID: "CgACAgIAAx0CQvXPNQABHGrDYIBIvDLiVV6ZMPypWMi_NVDkoFQAAq4LAAIwqQlIQT82LRwIpmoeBA"}})
		if err != nil {
			utils.ErrorReporting(err, m)
			return
		}
		return
	}
	var text = strings.Split(m.Text, " ")
	if len(text) != 2 {
		_, err := utils.Bot.Reply(m, "Пример использования: <code>/del {гет}</code>")
		if err != nil {
			utils.ErrorReporting(err, m)
			return
		}
		return
	}
	result := utils.DB.Delete(&utils.Get{Name: strings.ToLower(text[1])})
	if result.RowsAffected != 0 {
		_, err := utils.Bot.Reply(m, fmt.Sprintf("Гет <code>%v</code> удалён.", text[1]))
		if err != nil {
			utils.ErrorReporting(err, m)
			return
		}
	} else {
		_, err := utils.Bot.Reply(m, fmt.Sprintf("Гет <code>%v</code> не найден.", text[1]))
		if err != nil {
			utils.ErrorReporting(err, m)
			return
		}
	}
}
