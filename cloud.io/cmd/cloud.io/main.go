package main

import (
	"context"
	"log"
	"os"

	"github.com/shomali11/slacker"

	"./pluginManager"
	
)



func main() {
	// Create a new Slack client
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	// Create a plugin manager and load the plugins
    pluginManager := pluginManager.PluginManager{}
    pluginManager.LoadPlugin("./plugins/plugin1.so")
    pluginManager.LoadPlugin("./plugins/plugin2.so")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the Slack client
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}