package sdk

type BotCore interface {
	SendMessage(channel string, message string) error
	Log(message string)
	SomeMainFunction() string
}

type Plugin interface {
	Name() string
	Execute(data interface{}, core BotCore) error
}
