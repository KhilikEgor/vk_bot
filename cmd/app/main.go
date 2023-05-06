package main

import (
	"context"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/object"
	"vk_bot/internal/app"
	internalEvents "vk_bot/internal/events"
)

func main() {
	bot := app.New()

	bot.GetFuncList().MessageNew(func(ctx context.Context, object events.MessageNewObject) {
		bot := app.ExtractBotTx(ctx)

		if bot.MessageIsCommand(object) {
			return
		}
		if object.Message.Text == "Start" || object.Message.Text == "Старт" {
			fmt.Printf("%#v\n", *bot)
			fmt.Println(object.Message.Text)

			_, _ = bot.SendMessage(object.Message.FromID,
				"Я бот созданный для игры в камень ножницы бумага."+"\n\n"+
					"Выбирай, камень💩 ножницы✂ бумага📰 или рандом👀", nil)

		}

	})

	bot.HandleCommandForStart("Start", bot)

	bot.HandleCommandForStart("Старт", bot)

	bot.HandleCommandForStart("Заново", bot)

	bot.HandleCommandForСhoice("Рандом", bot)

	bot.HandleCommandForСhoice("Камень", bot)

	bot.HandleCommandForСhoice("Ножницы", bot)

	bot.HandleCommandForСhoice("Бумага", bot)

	bot.GetFuncList().CommandNew("Завершить", func(ctx context.Context, command internalEvents.CommandNew) {
		keyboard := object.NewMessagesKeyboard(true).
			AddRow().AddTextButton("Start", "inline-keyboard", "primary")
		_, _ = bot.SendMessage(command.Object.Message.PeerID, "Ты достойный соперник!"+"\n"+
			"Надеюсь тебе понравилась игра :)"+"\n\n"+
			"До встречи!✋", &api.Params{
			"keyboard": keyboard.ToJSON(),
		})

	})
	bot.Polling()

}
