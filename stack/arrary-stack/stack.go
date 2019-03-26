package arrarystack

type Stack struct {
	Elements []int
	Size     int	// size-1 为栈顶，入栈时先存入，size再加一, 栈顶一般指向下一个元素
}

func New(initSize int) *Stack {
	if initSize < 0 {
		return nil
	}
	return &Stack{
		Elements: make([]int, initSize),
		Size:     0,
	}
}

func (s *Stack) Peek() int {
	if s.Size == 0 {
		return -1
	}
	return s.Elements[s.Size-1]
}

func (s *Stack) Push(data int)  bool {
	if s.Size == len(s.Elements){
		return false
	}
	s.Elements[s.Size] = data
	s.Size ++
	return true
}

func (s *Stack) Pop() (int, bool) {
	if s.Size < 1 {
		return -1, false
	}
	s.Size --
	poped := s.Elements[s.Size]
	return poped, true
}

func (s *Stack) Len() int {
	return s.Size
}

func (s *Stack) IsEmpty() bool {
	return s.Size == 0
}