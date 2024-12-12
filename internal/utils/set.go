package utils

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{make(map[T]struct{})}
}

func (s *Set[T]) FromSlice(source []T) *Set[T] {
	for _, v := range source {
		s.m[v] = struct{}{}
	}
	return s
}

func (s *Set[T]) Add(values ...T) {
	for _, v := range values {
		s.m[v] = struct{}{}
	}
}

func (s *Set[T]) Remove(value T) {
	delete(s.m, value)
}

func (s *Set[T]) Contains(value T) bool {
	_, ok := s.m[value]
	return ok
}

func (s *Set[T]) Len() int {
	return len(s.m)
}

func (s *Set[T]) Clear() {
	s.m = make(map[T]struct{})
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.m) == 0
}

func (s *Set[T]) GetAll() []T {
	result := make([]T, 0)
	for k := range s.m {
		result = append(result, k)
	}
	return result
}
