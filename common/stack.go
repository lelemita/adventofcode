package common

type Stack[T any] interface {
	Len() int
	IsEmpty() bool
	Push(T)
	Pop() *T
}

type stack[T any] struct {
	data []T
}

func (s *stack[T]) Len() int {
	return len(s.data)
}

func (s *stack[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *stack[T]) Push(obj T) {
	s.data = append(s.data, obj)
}

func (s *stack[T]) Pop() *T {
	if l := len(s.data); l > 0 {
		result := s.data[l-1]
		s.data = s.data[:l-1]
		return &result
	}
	return nil
}

func NewRuneStack() *stack[rune] {
	return &stack[rune]{
		data: []rune{},
	}
}

func NewRuneStackWithData(data []rune) *stack[rune] {
	return &stack[rune]{
		data: data,
	}
}

func NewIntStack() *stack[int] {
	return &stack[int]{
		data: []int{},
	}
}
