package ds

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

type Node[T any] struct {
	value	T
	next	*Node[T]
}

type Stack[T any] struct {
	top		*Node[T]
	_size	int
}

func (s *Stack[T]) push(value T) {
	node := &Node[T]{
		value: value,
		next: s.top,
	}
	s.top = node
	s._size++
}

func (s *Stack[T]) pop() T {
	if s.top == nil {
		panic("Stack is empty")
	}
	value := s.top.value
	s.top = s.top.next
	s._size--
	return value
}

func (s *Stack[T]) size() int {
	return s.size()
}

type MinStackPair struct {
	value, min	int
}
type MinStack struct {
	s	*Stack[MinStackPair]
}

func (ms *MinStack) push(value int) {
	min := Min(ms.s.top.value.min, value)
	ms.s.push(MinStackPair{value: value, min: min})
}

func (ms *MinStack) pop() int {
	return ms.s.pop().value
}

func (ms *MinStack) size() int {
	return ms.s.size()
}

func main() {
	
}