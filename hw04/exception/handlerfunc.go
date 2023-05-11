package exception

import (
	"otus-architecture/hw04/additions/queue"
	"otus-architecture/hw04/command"
	"otus-architecture/hw04/command/log"
	"otus-architecture/hw04/command/repeat"
	"otus-architecture/hw04/command/repeatcount"
	"otus-architecture/hw04/object"
)

func AddRepeatCommand(cmd command.CommandInterface, err error) {
	repeatCommand := repeat.NewRepeatCommand(cmd)
	queue.CommandQueue.Add(repeatCommand)
}

func AddRepeatCountCommand(cmd command.CommandInterface, err error) {
	repeatCountCommand := repeatcount.NewRepeatCountCommand(cmd, 1)
	queue.CommandQueue.Add(repeatCountCommand)
}

func AddLogCommand(cmd command.CommandInterface, err error) {
	obj := object.New()
	_ = obj.SetProperty("Error", err)

	logCommand := log.NewLogCommand(log.NewLogAdapter(obj))
	queue.CommandQueue.Add(logCommand)
}
