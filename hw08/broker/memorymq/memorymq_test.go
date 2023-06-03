package memorymq

import (
	"github.com/stretchr/testify/require"
	"otus-architecture/hw08/broker"
	"testing"
)

func TestMemorySendAndReceive(t *testing.T) {
	args := make([]interface{}, 0)
	args = append(args, 1, 2, 3, nil, 4, "ok")

	message := &broker.Message{
		GameID:      1,
		ObjectID:    2,
		OperationID: 3,
		Args:        args,
	}

	memoryMQ := New()

	err := memoryMQ.Connect()
	require.Equal(t, nil, err)

	err = memoryMQ.Send(message)
	require.Equal(t, nil, err)

	result, err := memoryMQ.Receive()
	require.Equal(t, nil, err)

	require.Equal(t, message, result)

	err = memoryMQ.Close()
	require.Equal(t, nil, err)
}
