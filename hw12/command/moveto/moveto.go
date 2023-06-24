package moveto

import "otus-architecture/hw12/additions/queue"

type MoveToInterface interface {
	GetMoveToQueue() (chan queue.Element, error)
	GetQueue() (chan queue.Element, error)
	GetRunChannel() (chan struct{}, error)
}

type MoveToCommand struct {
	obj MoveToInterface
}

func NewMoveToCommand(obj MoveToInterface) *MoveToCommand {
	return &MoveToCommand{obj: obj}
}

func (c *MoveToCommand) Execute() error {
	moveTo, err := c.obj.GetMoveToQueue()
	if err != nil {
		return err
	}

	queue, err := c.obj.GetQueue()
	if err != nil {
		return err
	}

	run, err := c.obj.GetRunChannel()
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case elem := <-queue:
				moveTo <- elem
			case <-run:
				return
			}
		}
	}()

	return nil
}
