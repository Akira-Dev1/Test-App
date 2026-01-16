package bot

import (
    tgbotapi "gopkg.in/telegram-bot-api.v4"

    "telegram_client/handlers"
)

func (a *App) Router() {
    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates, _ := a.Bot.GetUpdatesChan(u)

    for update := range updates {

        if update.CallbackQuery != nil {
            handlers.HandleCallback(a.Bot, update.CallbackQuery, &update)
            continue
        }

        if update.Message != nil {
            handlers.HandleMessage(a.Bot, update.Message, &update)
        }
    }
}
