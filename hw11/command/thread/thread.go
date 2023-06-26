package thread

import (
	"fmt"
	"otus-architecture/hw11/additions/queue"
	"otus-architecture/hw11/exception"
	"otus-architecture/hw11/object"
	"reflect"
)

type ThreadInterface interface {
	GetQueueChannel() (chan queue.Element, error)
	GetHardStopChannel() (chan struct{}, error)
	GetSoftStopChannel() (chan struct{}, error)
	GetEndChannel() (chan struct{}, error)
}

type ThreadCommand struct {
	obj ThreadInterface
}

func NewThreadCommand(obj ThreadInterface) *ThreadCommand {
	return &ThreadCommand{obj: obj}
}

func (c *ThreadCommand) Execute() (*object.Object, error) {
	queue, err := c.obj.GetQueueChannel()
	if err != nil {
		return nil, err
	}

	hardStopChannel, err := c.obj.GetHardStopChannel()
	if err != nil {
		return nil, err
	}

	softStopChannel, err := c.obj.GetSoftStopChannel()
	if err != nil {
		return nil, err
	}

	endChannel, err := c.obj.GetEndChannel()
	if err != nil {
		return nil, err
	}

	exceptionHandler := exception.NewException()

	go func() {
		fmt.Println("start thread")

		for cmd := range queue {
			fmt.Println(reflect.TypeOf(cmd).String())

			_, err := cmd.Execute()

			if err != nil {
				exceptionHandler.Handle(cmd, err)
			}

			select {
			case <-softStopChannel:
				fmt.Println("soft stop")
				close(queue)
			case <-hardStopChannel:
				fmt.Println("hard stop")
				close(endChannel)
				return
			default:
			}
		}
		fmt.Println("stop thread")
		close(endChannel)
	}()

	return nil, nil
}
