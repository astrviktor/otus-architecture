package queue

import (
	"errors"
	"fmt"
	"otus-architecture/hw12/command"
)

type Element command.CommandInterface

type Queue struct {
	Elements []Element
	Size     int
}

func New(size int) *Queue {
	return &Queue{
		Elements: nil,
		Size:     size,
	}
}

func (q *Queue) Add(elem Element) {
	if q.GetLength() == q.Size {
		fmt.Println("Overflow")
		return
	}
	q.Elements = append(q.Elements, elem)
}

func (q *Queue) Take() Element {
	if q.IsEmpty() {
		fmt.Println("UnderFlow")
		return nil
	}
	element := q.Elements[0]
	if q.GetLength() == 1 {
		q.Elements = nil
		return element
	}
	q.Elements = q.Elements[1:]
	return element // Slice off the element once it is dequeued.
}

func (q *Queue) GetLength() int {
	return len(q.Elements)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q *Queue) Peek() (Element, error) {
	if q.IsEmpty() {
		return nil, errors.New("empty queue")
	}
	return q.Elements[0], nil
}

const QueueMaxSize = 100

var CommandQueue *Queue
