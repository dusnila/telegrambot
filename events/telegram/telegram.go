package telegram

import "telegramBot/cliens/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	// storage
}
