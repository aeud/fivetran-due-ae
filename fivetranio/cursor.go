package fivetranio

type Cursor struct {
	Next string `json:"next_value"`
}

func (s *State) GetCursor(v string) (c *Cursor) {
	if s.Cursors == nil {
		s.Cursors = make(map[string]*Cursor)
	}
	c, ok := s.Cursors[v]
	if !ok {
		c = &Cursor{}
		s.Cursors[v] = c
	}
	return
}

func (s *State) GetCursorNextValue(stepName string) string {
	return s.GetCursor(stepName).Next
}

func (s *State) SetCursorNextValue(stepName, val string) (updated bool) {
	c := s.GetCursor(stepName)
	previousValue := c.Next
	c.Next = val
	updated = previousValue != val
	return
}
