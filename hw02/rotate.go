package hw02

type RotateInterface interface {
	GetDirection() (int, error)
	SetDirection(n int) error
}

type Rotate struct{}

func (rotate *Rotate) Execute(obj RotateInterface) error {
	n, err := obj.GetDirection()
	if err != nil {
		return err
	}

	n = n + 1

	return obj.SetDirection(n)
}
