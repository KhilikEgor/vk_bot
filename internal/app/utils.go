package app

import (
	"context"
	"vk_bot/internal/utils"
)

func ExtractBotTx(ctx context.Context) *Bot {
	return ctx.Value(utils.ContextKeyBot).(*Bot)
}
