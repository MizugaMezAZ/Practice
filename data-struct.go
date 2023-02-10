package main

type stack []rune

func NewStack() stack {
	return stack{}
}

func (s *stack) Push(r rune) {
	*s = append(*s, r)
}

func (s *stack) Peek() rune {
	if len(*s) == 0 {
		return ' '
	}

	return (*s)[len(*s)-1]
}

func (s *stack) Pop() rune {
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r
}

func (s *stack) Empty() bool {
	return len(*s) == 0
}


type queue []int

func NewQueue() queue{
	return queue{}
}
