package day5

import (
	"container/list"
	"fmt"
)

type Stack struct {
	stack *list.List
}

func (s Stack) Empty() bool {
	return s.stack.Len() == 0
}

func (s Stack) Peek() (interface{}, error) {
	if s.stack.Len() > 0 {
		return s.stack.Front().Value, nil
	} else {
		return nil, fmt.Errorf("Peek Error: Stack is empty")
	}
}

func (s *Stack) Push(element interface{}) interface{} {
	s.stack.PushFront(element)
	return element
}

func (s *Stack) Pop() (interface{}, error) {
	if s.stack.Len() > 0 {
		element := s.stack.Front()
		return s.stack.Remove(element), nil
	} else {
		return nil, fmt.Errorf("Pop Error: Stack is empty")
	}
}

func (s *Stack) PopMany(total int) ([]interface{}, error) {

	if total > s.stack.Len() {
		return nil, fmt.Errorf("Pop Error: Popping a greater number of items than currently stacked")
	}

	var elements []interface{}

	for i := 0; i < total; i++ {
		element := s.stack.Front()
		elements = append(elements, element.Value)
		s.stack.Remove(element)
	}

	return elements, nil
}

func (s Stack) Size() int {
	return s.stack.Len()
}

func NewStack() Stack {
	return Stack{
		stack: list.New(),
	}
}
