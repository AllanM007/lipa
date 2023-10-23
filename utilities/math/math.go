package math

import "golang.org/x/exp/constraints"

func Add[T constraints.Integer](a , b T) T {
	return a + b
}

func Subtract[T constraints.Integer](a , b T) T {
	return a - b
}

func Multiply[T constraints.Integer](a , b T) T {
	return a * b
}

func Divide[T constraints.Integer](a , b T) T {
	return a / b
}

func Min[T constraints.Ordered](a , b T) T {
	if a > b {
		return b
	}
	return a
}

func Max[T constraints.Ordered](a , b T) T {
	if a > b {
		return a
	}
	return b
}