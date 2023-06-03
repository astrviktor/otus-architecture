package broker

type BrockerInterface interface {
	Connect() error
	Send(message Message) error
	Receive() (Message, error)
	Close() error
}

type Message struct {
	GameID      int
	ObjectID    int
	OperationID int
	Args        []interface{}
}
