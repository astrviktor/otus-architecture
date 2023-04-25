package command

type CommandInterface interface {
	Execute() error
}
