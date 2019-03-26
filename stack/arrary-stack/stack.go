package arrarystack

type Stack struct {
	elements []int
	size     int	// size-1 为栈顶，入栈时先存入，size再加一, 栈顶一般指向下一个元素
}

func New(initSize int) *Stack {
	if initSize < 0 {
		return nil
	}
	return &Stack{
		elements: make([]int, initSize),
		size:     0,
	}
}

func (s *Stack) peek() int {
	if s.size == 0 {
		return -1
	}
	return s.elements[s.size-1]
}

func (s *Stack) push(data int)  bool {
	if s.size == len(s.elements){
		return false
	}
	s.elements[s.size] = data
	s.size ++
	return true
}

func (s *Stack) pop() (int, bool) {
	if s.size < 1 {
		return -1, false
	}
	s.size --
	poped := s.elements[s.size]
	return poped, true
}

func (s *Stack) len() int {
	return s.size
}
