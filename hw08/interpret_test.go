package hw08

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw08/additions/queue"
	"otus-architecture/hw08/additions/vector"
	"otus-architecture/hw08/broker"
	"otus-architecture/hw08/broker/memorymq"
	interpret "otus-architecture/hw08/command/Interpret"
	"otus-architecture/hw08/ioc"
	"otus-architecture/hw08/object"
	"testing"
)

func TestInterpretMove(t *testing.T) {
	// Create message from Agent and send to Endpoint
	args := make([]interface{}, 0)
	args = append(args, vector.Vector{0, 0}, 8, 0, 2)

	messageFromAgent := &broker.Message{
		GameID:      1,
		ObjectID:    1,
		OperationID: 1,
		Args:        args,
	}

	memoryMQ := memorymq.New()

	err := memoryMQ.Connect()
	require.Equal(t, nil, err)

	err = memoryMQ.Send(messageFromAgent)
	require.Equal(t, nil, err)

	// Endpoint receive message from Agent
	message, err := memoryMQ.Receive()
	require.Equal(t, nil, err)
	require.Equal(t, messageFromAgent, message)

	// Init game before process message
	IoC := ioc.New()
	_, err = IoC.Resolve("Register", "MakeObject", 1, 1).Execute()
	require.Equal(t, nil, err)

	// Check CommandQueue before interpretCommand
	queueLen := queue.CommandQueue.GetLength()
	require.Equal(t, 0, queueLen)

	// Make interpretObject and interpretCommand
	interpretObject := object.New()
	interpretObject.SetProperty("IoC", IoC)
	interpretObject.SetProperty("Message", message)

	interpretAdapter := interpret.NewInterpretdapter(interpretObject)
	interpretCommand := interpret.NewInterpretCommand(interpretAdapter)
	err = interpretCommand.Execute()
	require.Equal(t, nil, err)

	// Check CommandQueue after interpretCommand
	queueLen = queue.CommandQueue.GetLength()
	require.Equal(t, 1, queueLen)

	// Take MoveCommand from CommandQueue, execute and check result
	cmd := queue.CommandQueue.Take()
	obj, err := cmd.Execute()
	require.Equal(t, nil, err)

	position, err := obj.GetProperty("Position")

	require.Equal(t, nil, err)
	require.Equal(t, vector.Vector{2, 0}, position)
}
