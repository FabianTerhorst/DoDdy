package botcommands

import "github.com/bwmarrin/discordgo"
import "github.com/Devs-On-Discord/DoDdy/commands"
import (
	"github.com/Devs-On-Discord/DoDdy/guilds"
)

func setAnnouncementsChannel(session *discordgo.Session, commandMessage *discordgo.MessageCreate, args []string) commands.CommandResultMessage {
	channelId := commandMessage.ChannelID
	channel, err := session.Channel(channelId)
	if err != nil {
		return &commands.CommandError{Message: "Announcement channel couldn't be identified " + err.Error(), Color: 0xb30000}
	}
	err = guilds.SetAnnouncementsChannel(channel.GuildID, channelId)
	if err != nil {
		return &commands.CommandError{Message: err.Error(), Color: 0xb30000}
	}
	return &commands.CommandReply{Message: "Announcement channel set to " + channel.Name, Color: 0x00b300}
}

func postAnnouncement(session *discordgo.Session, commandMessage *discordgo.MessageCreate, args []string) commands.CommandResultMessage {
	if len(args) < 1 {
		return &commands.CommandError{Message: "Announcement message missing", Color: 0xb30000}
	}
	announcement := args[0]
	channels, err := guilds.GetAnnouncementChannels()
	if err != nil {
		return &commands.CommandError{Message: err.Error(), Color: 0xb30000}
	}
	for _, channelID := range channels {
		session.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
			Color: 0x00b300,
			Title: announcement,
		})
	}
	return &commands.CommandReply{Message: "Announcement posted", Color: 0x00b300}
}

func setup(session *discordgo.Session, commandMessage *discordgo.MessageCreate, args []string) commands.CommandResultMessage {
	channelId := commandMessage.ChannelID
	channel, err := session.Channel(channelId)
	if err != nil {
		return &commands.CommandError{Message: "Server couldn't be identified " + err.Error(), Color: 0xb30000}
	}
	guildId := channel.GuildID
	guild, err := session.Guild(guildId)
	if err != nil {
		return &commands.CommandError{Message: "Server couldn't be identified " + err.Error(), Color: 0xb30000}
	}
	err = guilds.Create(guildId, guild.Name)
	if err != nil {
		return &commands.CommandError{Message: err.Error(), Color: 0xb30000}
	}
	return &commands.CommandReply{Message: "setup", Color: 0x00b300}
}