package bots

import (
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

type TelegramBot struct {
	bot *telebot.Bot
}

func (tb *TelegramBot) Initialize(botToken, webAppURL string) error {
	if botToken == "" || webAppURL == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN or WEB_APP_URL is not set")
		return nil
	}

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  botToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatalf("Failed to create Telegram bot: %v", err)
		return err
	}

	tb.bot = bot

	tb.bot.Handle("/start", func(m *telebot.Message) {
		replyMarkup := &telebot.ReplyMarkup{}
		webAppBtn := replyMarkup.URL("Open Web App", webAppURL)
		replyMarkup.Inline(replyMarkup.Row(webAppBtn))

		_, err := tb.bot.Send(m.Sender, "Welcome! Click below to open the web app.", replyMarkup)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	})

	return nil
}

func (tb *TelegramBot) Start() {
	go tb.bot.Start()
}
