package main

import (
    "github.com/shomali11/slacker"
)

type MyPlugin struct {}

func (p *MyPlugin) RegisterCommands(bot *slacker.Slacker) {
    bot.Command("mycommand", func(request slacker.Request, response slacker.ResponseWriter) {
        // Handle the mycommand command here
		definition := &slacker.CommandDefinition{
			Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
				response.Reply("pong")
			},
		}
		
		bot.Command("ping", definition)

    })
}




