package shared

import "errors"

type Token struct {
	Tok_type string
	Value    string
}

// TODO!!!!! Remove ugly repeated code

type NumStack struct {
	self []float64
}

func New_num_stack() *NumStack {
	st := new(NumStack)
	st.self = make([]float64, 0)
	return st
}

func (s *NumStack) Push(t float64) {
	s.self = append(s.self, t)
}

func (s *NumStack) Pop() (t float64) {
	t, s.self = s.self[len(s.self)-1], s.self[:len(s.self)-1]
	return t
}

func (s *NumStack) Peek() (float64, error) {
	if len(s.self) == 0 {
		return 0, errors.New("No such element")
	}
	return s.self[len(s.self)-1], nil
}

func (s *NumStack) GetAtIndex(i int) (float64, error) {
	if i >= len(s.self) {
		return 0, errors.New("No such element")
	}
	return s.self[i], nil
}

// will be used for operator stack as well as the rpn evaluator
type Stack struct {
	self []*Token
}

func New_stack() *Stack {
	st := new(Stack)
	st.self = make([]*Token, 15)
	return st
}

func (s *Stack) Push(t *Token) {
	s.self = append(s.self, t)
}

func (s *Stack) Pop() (t *Token) {
	t, s.self = s.self[len(s.self)-1], s.self[:len(s.self)-1]
	return t
}

func (s *Stack) Peek() (t *Token) {
	if len(s.self) == 0 {
		return nil
	}
	return s.self[len(s.self)-1]
}

func (s *Stack) GetAtIndex(i int) *Token {
	if i >= len(s.self) {
		return nil
	}
	return s.self[i]
}

func New_token(tp, val string) *Token {
	return &Token{Tok_type: tp, Value: val}
}
