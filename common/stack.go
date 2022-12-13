package common

type runeStack struct {
	Data []rune
}

func NewRuneStack() runeStack {
	return runeStack{
		Data: []rune{},
	}
}

func (s *runeStack) IsEmpty() bool {
	return len(s.Data) == 0
}

func (s *runeStack) Push(obj rune) {
	s.Data = append(s.Data, obj)
}

func (s *runeStack) Pop() *rune {
	if l := len(s.Data); l > 0 {
		result := s.Data[l-1]
		s.Data = s.Data[:l-1]
		return &result
	}
	return nil
}

func (s *runeStack) CopyData() []rune {
	return s.Data
}
