package main

func DynamicProgramming[T comparable](tableH, tableW uint32, calcCell func(table [][]T, y, x uint32) T) T {
	table := make([][]T, tableH)

	for y := uint32(0); y < tableH; y++ {

		for x := uint32(0); x < tableW; x++ {
			cell := calcCell(table, y, x)
			table[y] = append(table[y], cell)
		}
	}

	return table[tableH-1][tableW-1]
}
