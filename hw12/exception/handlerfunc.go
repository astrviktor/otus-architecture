package exception

import (
	"otus-architecture/hw12/additions/queue"
	"otus-architecture/hw12/command"
	"otus-architecture/hw12/command/log"
	"otus-architecture/hw12/command/repeat"
	"otus-architecture/hw12/command/repeatcount"
	"otus-architecture/hw12/object"
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
	obj.SetProperty("Error", err)

	logCommand := log.NewLogCommand(log.NewLogAdapter(obj))
	queue.CommandQueue.Add(logCommand)
}
