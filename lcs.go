package main

func Lcs(a, b string) uint32 {
	ra := []rune(a)
	rb := []rune(b)

	tableH := len(ra) + 1
	tableW := len(rb) + 1
	table := make([][]uint32, tableH)
	for i := 0; i < tableH; i++ {
		table[i] = make([]uint32, tableW)
	}

	for y := 1; y < tableH; y++ {
		for x := 1; x < tableW; x++ {
			if ra[y-1] == rb[x-1] {
				table[y][x] = table[y-1][x-1] + 1
			} else {
				table[y][x] = Max(table[y-1][x], table[y][x-1])
			}
		}
	}

	return table[tableH-1][tableW-1]
}
