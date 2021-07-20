package stack

const _INIT_STACK_SIZE = 10

type Stack struct {
	elements []interface{}
	top      int
}

func New() Stack {
	stack := Stack{}
	stack.elements = make([]interface{}, 0, _INIT_STACK_SIZE)
	stack.top = -1
	return stack
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Push(element interface{}) {
	s.elements = append(s.elements, element)
	s.top++
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.top >= 0 {
		element := s.elements[s.top]
		s.elements = s.elements[:s.top]
		s.top--
		return element, true
	}
	return nil, false
}

func (s *Stack) Top() int {
	return s.top
}