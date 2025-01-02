package telegram

import (
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

// Bot is exported to potentially allow other packages to use it
var Bot *telebot.Bot

// InitBot initializes the Telegram bot
func InitBot(botToken, webAppURL string) error {
	if botToken == "" || webAppURL == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN or REACT_APP_URL is not set in .env")
		return nil
	}

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  botToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
		return err
	}

	bot.Handle("/start", func(m *telebot.Message) {
		replyMarkup := &telebot.ReplyMarkup{}
		webAppBtn := replyMarkup.URL("Open Web App", webAppURL)
		replyMarkup.Inline(replyMarkup.Row(webAppBtn))

		_, err := bot.Send(m.Sender, "Welcome! Click below to open the web app.", replyMarkup)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	})

	Bot = bot

	go bot.Start()

	return nil
}
