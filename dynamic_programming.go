package main

type Item struct {
	Id, Cost, Value uint32
}

func DynamicProgramming(capacity uint32, items []Item) uint32 {
	numItem := len(items)

	// 最大公約数で割ってtableの大きさを最適化
	costs := make([]uint32, numItem)
	for i := range costs {
		costs[i] = items[i].Cost
	}
	gcd := GcdVec(append(costs, capacity))
	capacity /= gcd
	for i := range items {
		items[i] = Item{Id: items[i].Id, Cost: items[i].Cost / gcd, Value: items[i].Value}
	}

	table := make([][]uint32, numItem)
	for i := range table {
		table[i] = make([]uint32, capacity+1)
	}

	for i := 1; i < numItem; i++ {
		for j := uint32(0); j <= capacity; j++ {
			if items[i].Cost > j {
				// items[i]が入らない場合 : 上のマスを代入
				table[i][j] = table[i-1][j]
			} else {
				// items[i]が入る場合 : 上のマスと比較して大きい方を代入
				table[i][j] = Max(
					table[i-1][j-items[i].Cost]+items[i].Value,
					table[i-1][j],
				)
			}
		}
	}

	return table[numItem-1][capacity]
}
