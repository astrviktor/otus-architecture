package exception

import (
	"otus-architecture/hw07/additions/queue"
	"otus-architecture/hw07/command"
	"otus-architecture/hw07/command/log"
	"otus-architecture/hw07/command/repeat"
	"otus-architecture/hw07/command/repeatcount"
	"otus-architecture/hw07/object"
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
