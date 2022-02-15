package gnt

import (
	"os"
	"strconv"

	tba "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func init() {
	var err error
	if bot, err = tba.NewBotAPI(os.Getenv("GNT_TELEGRAM_TOKEN")); err != nil {
		panic(err)
	}
	if tid, err = strconv.ParseInt(os.Getenv("GNT_TELEGRAM_ID"), 10, 64); err != nil {
		panic(err)
	}
	if tid == 0 {
		panic("GNT_TELEGRAM_ID is not set")
	}
}

var (
	bot *tba.BotAPI
	tid int64
)

func SendText(msg string) {
	mc := tba.NewMessage(5099998680, msg)
	mc.ParseMode = "MarkdownV2"

	_, err := bot.Send(mc)
	if err != nil {
		panic(err)
	}
}

func SendMessage(msg, url string) {
	mc := tba.NewMessage(5099998680, msg)
	mc.ParseMode = "MarkdownV2"
	if url != "" {
		mc.ReplyMarkup = tba.NewInlineKeyboardMarkup(tba.NewInlineKeyboardRow(
			tba.NewInlineKeyboardButtonURL("查看", url),
		))
	}

	_, err := bot.Send(mc)
	if err != nil {
		panic(err)
	}
}
