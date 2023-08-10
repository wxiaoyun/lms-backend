package ternary

//nolint:revive
func TernaryOperater[T any](condition bool, trueVal T, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

//nolint:revive
func LazyTernaryOperator[T any](condition bool, trueVal, falseVal func() T) T {
	if condition {
		return trueVal()
	}
	return falseVal()
}

type TernaryIf[T any] struct {
	condition bool
}

type TernaryThen[T any] struct {
	condition bool
	trueVal   T
}

type LazyTernaryThen[T any] struct {
	condition bool
	trueVal   func() T
}

func If[T any](condition bool) *TernaryIf[T] {
	return &TernaryIf[T]{condition: condition}
}

func (i *TernaryIf[T]) Then(trueVal T) *TernaryThen[T] {
	return &TernaryThen[T]{condition: i.condition, trueVal: trueVal}
}

func (t *TernaryThen[T]) Else(falseVal T) T {
	return TernaryOperater[T](t.condition, t.trueVal, falseVal)
}

func (i *TernaryIf[T]) LazyThen(trueVal func() T) *LazyTernaryThen[T] {
	return &LazyTernaryThen[T]{condition: i.condition, trueVal: trueVal}
}

func (t *LazyTernaryThen[T]) LazyElse(falseVal func() T) T {
	return LazyTernaryOperator[T](t.condition, t.trueVal, falseVal)
}
