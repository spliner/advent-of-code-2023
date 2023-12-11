package set

var marker struct{}

type Set[T comparable] struct {
	items map[T]struct{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]struct{})}
}

func NewFromSlice[T comparable](slice []T) *Set[T] {
	set := New[T]()
	for _, item := range slice {
		set.Add(item)
	}
	return set
}

func (s *Set[T]) Length() int {
	return len(s.items)
}

func (s *Set[T]) IsEmpty() bool {
	return s.Length() == 0
}

func (s *Set[T]) Add(item T) bool {
	if _, ok := s.items[item]; ok {
		return false
	}

	s.items[item] = marker
	return true
}

func (s *Set[T]) Remove(item T) bool {
	if _, ok := s.items[item]; !ok {
		return false
	}

	delete(s.items, item)
	return true
}

func (s *Set[T]) Contains(item T) bool {
	_, ok := s.items[item]
	return ok
}

func (s *Set[T]) All() []T {
	slice := make([]T, len(s.items))
	var i int
	for k := range s.items {
		slice[i] = k
		i++
	}

	return slice
}
