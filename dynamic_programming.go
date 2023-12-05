package main

import "reflect"

func DynamicProgramming[T any, U comparable](
	elements T,
	calcCell func(table []U, iterators []int, elements T) U,
) U {
	lengths := calcLengths(elements)
	var table []U

	var makeTable func(iterators []int)
	makeTable = func(iterators []int) {
		if len(lengths) == len(iterators) {
			table = append(table, calcCell(table, iterators, elements))
		} else {
			length := lengths[len(iterators)]
			for i := 0; i < length; i++ {
				makeTable(append(iterators, i))
			}
		}
	}

	makeTable([]int{})

	return table[len(table)-1]
}

func calcLengths[T any](elements T) []int {
	var result []int
	rv := reflect.ValueOf(elements)
	rt := reflect.TypeOf(elements)
	for i := 0; i < rt.NumField(); i++ {
		f := rv.Field(i)

		switch f.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			result = append(result, int(f.Uint())+1)
		case reflect.String:
			result = append(result, len(f.String())+1)
		case reflect.Array, reflect.Slice:
			result = append(result, f.Len()+1)
		default:
			result = append(result, 1)
		}
	}

	return result
}
