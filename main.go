package main
import (
    "discord-bot/bot"
    "discord-bot/config"
    "log"
)
funcmain() {
    err := config.ReadConfig()
    if err != nil {
        log.Fatal(err)
        return
    }
    bot.Run()
    <-make(chanstruct{})
    return
}
