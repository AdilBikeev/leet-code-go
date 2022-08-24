package DataStructure

type Stack[T any] struct {
	values []T
}

func NewStack[T any](values []T) *Stack[T] {
	return &Stack[T]{
		values: values,
	}
}

func (this *Stack[T]) Size() int {
	return len(this.values)
}

func (this *Stack[T]) Push(val T) {
	this.values = append(this.values, val)
}

func (this *Stack[T]) Top() T {
	return this.values[len(this.values)-1]
}

func (this *Stack[T]) Pop() T {
	n := len(this.values)
	top := this.values[n-1]
	this.values = this.values[:n]
	return top
}
