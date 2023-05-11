package thread

import (
	"fmt"
	"otus-architecture/hw07/additions/queue"
	"otus-architecture/hw07/exception"
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

func (c *ThreadCommand) Execute() error {
	queue, err := c.obj.GetQueueChannel()
	if err != nil {
		return err
	}

	hardStopChannel, err := c.obj.GetHardStopChannel()
	if err != nil {
		return err
	}

	softStopChannel, err := c.obj.GetSoftStopChannel()
	if err != nil {
		return err
	}

	endChannel, err := c.obj.GetEndChannel()
	if err != nil {
		return err
	}

	exceptionHandler := exception.NewException()

	go func() {
		fmt.Println("start thread")

		for cmd := range queue {
			fmt.Println(reflect.TypeOf(cmd).String())

			err := cmd.Execute()

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

	return nil
}
