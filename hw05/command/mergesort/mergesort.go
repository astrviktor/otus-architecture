package mergesort

// Сортировка слиянием
// https://neerc.ifmo.ru/wiki/index.php?title=Сортировка_слиянием

type MergeSortCommand struct {
	data []int
}

func NewMergeSortCommand(data []int) *MergeSortCommand {
	return &MergeSortCommand{data: data}
}

func (c *MergeSortCommand) Execute() error {
	sorted := mergeSort(c.data)

	copy(c.data, sorted)

	return nil
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
