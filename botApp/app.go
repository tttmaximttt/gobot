package botApp

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"github.com/tttmaximttt/gobot/config"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"io"
	"strconv"
)

func New(config config.Config) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	fmt.Println("WEBHOOK_URL>> ", config.WebhookURL)
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(config.WebhookURL))
	return bot, err
}

func Run(bot tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		fmt.Println("HERE")
 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		u1, err := uuid.FromBytes([]byte("tytryujhgtfrgfdr"))

		fmt.Println(err)

		fmt.Println(">>>> TEXT", update.Message.Text)
		fmt.Println(">>>> COMMAND", update.Message.Command())
		fmt.Println(">>>> CHAT ID", update.Message.Chat.ID)
		fmt.Println(">>>> CHAT u1", u1)
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(update.Message.Chat.ID, 10))
		fmt.Printf("%x", h.Sum(nil))

		if update.Message.Command() != "" {
			// TODO handle comand
		} else if update.Message.Text == "" {
			// TODO handle text
		}

		msg.Text = "I don't know that commands"
		bot.Send(msg)
	}
}
