package insertionsort

// Сортировка вставками
// https://neerc.ifmo.ru/wiki/index.php?title=Сортировка_вставками

type InsertionSortCommand struct {
	data []int
}

func NewInsertionSortCommand(data []int) *InsertionSortCommand {
	return &InsertionSortCommand{data: data}
}

func (c *InsertionSortCommand) Execute() error {
	n := len(c.data)

	for i := 0; i <= n-1; i++ {
		j := i - 1
		for (j >= 0) && (c.data[j] > c.data[j+1]) {
			c.data[j], c.data[j+1] = c.data[j+1], c.data[j]
			j--
		}
	}

	return nil
}
