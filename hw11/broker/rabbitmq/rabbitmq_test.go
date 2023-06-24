package rabbitmq

// For test need a RabbitMQ (Docker)

//func TestRabbitmqSendAndReceive(t *testing.T) {
//	args := make([]interface{}, 0)
//	args = append(args, 1, 2, 3, nil, 4, "ok")
//
//	message := &broker.Message{
//		GameID:      "1",
//		ObjectID:    "2",
//		OperationID: "3",
//		Args:        nil,
//	}
//
//	rmq := New()
//
//	err := rmq.Connect()
//	require.Equal(t, nil, err)
//
//	err = rmq.Send(message)
//	require.Equal(t, nil, err)
//
//	result, err := rmq.Receive()
//	require.Equal(t, nil, err)
//
//	require.Equal(t, message, result)
//
//	err = rmq.Close()
//	require.Equal(t, nil, err)
//}
