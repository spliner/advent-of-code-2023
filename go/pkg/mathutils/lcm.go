package mathutils

import "golang.org/x/exp/constraints"

func LeastCommonMultiple[T constraints.Integer](slice []T) T {
	if len(slice) == 0 {
		return 0
	}
	if len(slice) == 1 {
		return slice[0]
	}

	result := leastCommonMultiple(slice[0], slice[1])
	for _, n := range slice[2:] {
		result = leastCommonMultiple(result, n)
	}

	return result
}

func leastCommonMultiple[T constraints.Integer](a, b T) T {
	return a * b / greatestCommonDivisor(a, b)
}

func greatestCommonDivisor[T constraints.Integer](a, b T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}
