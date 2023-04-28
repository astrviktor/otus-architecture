package selectionsort

// Сортировка выбором
// https://neerc.ifmo.ru/wiki/index.php?title=Сортировка_выбором

type SelectionSortCommand struct {
	data []int
}

func NewSelectionSortCommand(data []int) *SelectionSortCommand {
	return &SelectionSortCommand{data: data}
}

func (c *SelectionSortCommand) Execute() error {
	n := len(c.data)

	for i := 0; i <= n-2; i++ {
		for j := i + 1; j <= n-1; j++ {
			if c.data[i] > c.data[j] {
				c.data[i], c.data[j] = c.data[j], c.data[i]
			}
		}
	}

	return nil
}
