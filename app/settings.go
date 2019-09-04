package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jmoiron/sqlx"
)

const (
	VERSION        = "0.1.0"
	TELEGRAM_TOKEN = "[telegram-token]"
	BASE_URL       = "[https://website-url]"
)

var (
	DB   *sqlx.DB
	Bot  *tgbotapi.BotAPI
	Conf map[string]string
)
