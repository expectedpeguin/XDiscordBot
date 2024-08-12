package sdk
import "github.com/bwmarrin/discordgo"

type BotCore interface {
	SendEmbedMessage(channel string, message *discordgo.MessageEmbed, button ...discordgo.Button) error
	SendMessage(channel string, message string) error
	Log(message string)
	SomeMainFunction() string
}

type Plugin interface {
	Name() string
	Execute(data interface{}, core BotCore) error
}
