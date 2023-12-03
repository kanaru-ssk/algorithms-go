package main

func DynamicProgramming[T, U comparable](
	items []U,
	capacity uint32,
	calcCell func(y, x uint32, table [][]T, items []U) T,
) T {
	tableH := uint32(len(items) + 1)
	tableW := capacity + 1
	table := make([][]T, tableH)

	for y := uint32(0); y < tableH; y++ {
		for x := uint32(0); x < tableW; x++ {
			cell := calcCell(y, x, table, items)
			table[y] = append(table[y], cell)
		}
	}

	return table[tableH-1][tableW-1]
}
