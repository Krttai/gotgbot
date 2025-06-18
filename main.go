package main

import (
	"fmt"
	"log"

	"github.com/PaulSonOfLars/gotgbot/v2" // If your bot uses gotgbot v2
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"

	"gotgbot" // This imports your config.go (must be in same module or folder)
)

func main() {
	bot, err := gotgbot.NewBot(gotgbot.BOT_TOKEN, nil)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	fmt.Printf("Bot started as @%s\n", bot.User.Username)

	updater := ext.NewUpdater(nil)
	dispatcher := updater.Dispatcher

	// Command: /start
	dispatcher.AddHandler(handlers.NewCommand("start", func(b *gotgbot.Bot, ctx *ext.Context) error {
		_, err := b.SendMessage(ctx.EffectiveChat.Id, "Hello! I'm alive 🚀")
		return err
	}))

	err = updater.StartPolling(bot, nil)
	if err != nil {
		log.Fatalf("Failed to start polling: %v", err)
	}

	// Run until stopped
	updater.Idle()
}
