package fivetranio

import "fmt"

type Cursor struct {
	Current string `json:"current_value"`
	Target  string `json:"target_value"`
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

func (s *State) SetCurrentCursor(stepName, val string) {
	c := s.GetCursor(stepName)
	c.Current = val
}

func (s *State) SetTargetCursor(stepName, val string) {
	c := s.GetCursor(stepName)
	c.Target = val
}

func (s *State) IsLowerThanCurrentCursor(stepName string, v interface{}) bool {
	return fmt.Sprintf("%v", v) < s.GetCursor(stepName).Current
}

func (s *State) CloseCursor(stepName string) {
	if s.Cursors != nil {
		s.SetCurrentCursor(stepName, s.GetCursor(stepName).Target)
		s.SetTargetCursor(stepName, "")
	}
}
