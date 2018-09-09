package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Devs-On-Discord/DoDdy/bot/commands"
)

const version = "0.0.1"

func main() {
	fmt.Printf("DoDdy %s starting\n", version)

	db := db{}
	db.Init()

	defer db.Close()

	bot := bot{}
	bot.Init()

	defer bot.Close()

	g := &guilds{}
	g.Init(bot.session)

	v := &votes{}
	v.Init(bot.session)

	botCommands := &commands.Commands{}
	botCommands.Init(bot.session)
	botCommands.Validator = commandValidator{guilds: g}
	identifier := commandIdentifier{guilds: g}
	identifier.Init(bot.session)
	botCommands.Identifier = identifier
	botCommands.ResultHandler = commandResultHandler{}
	botCommands.RegisterGroup(guildAdminCommands{guilds: g, votes: v})
	botCommands.RegisterGroup(helpCommands{botCommands})
	botCommands.RegisterGroup(debugCommands{})
	questionCommands := questionCommands{}
	questionCommands.Init(bot.session)
	botCommands.RegisterGroup(questionCommands)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	fmt.Println("DoDdy ready.\nPress Ctrl+C to exit.")
	<-sc
}
