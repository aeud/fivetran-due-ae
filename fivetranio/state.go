package fivetranio

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type State struct {
	CurrentStep    string             `json:"current_step"`
	NextPageNumber int                `json:"next_page_number"`
	Cursors        map[string]*Cursor `json:"cursors,omitempty"`
	RemainingSteps []string           `json:"remaining_steps"`
	PageSize       int                `json:"page_size"`
}

func NewStateFromReader(r io.Reader, allSteps []string) (*State, error) {
	var v State
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return nil, err
	}
	if v.CurrentStep == "" {
		if err := v.Reset(allSteps, 0); err != nil {
			return nil, err
		}
	}
	return &v, nil
}
func NewStateFromJson(data []byte, allSteps []string) (*State, error) {
	var v State
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	if v.CurrentStep == "" {
		if err := v.Reset(allSteps, 0); err != nil {
			return nil, err
		}
	}
	return &v, nil
}
func NewStateFromJsonOrPanic(data []byte, allSteps []string) *State {
	s, err := NewStateFromJson(data, allSteps)
	if err != nil {
		panic(err)
	}
	return s
}
func (s *State) MarshalForce() []byte {
	bs, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return bs
}

func (s *State) LogContent() {
	log.Printf("%s\n", s.MarshalForce())
}

func (s *State) NextPage(pageNumber int) (*State, bool, error) {
	ns := &State{
		CurrentStep:    s.CurrentStep,
		Cursors:        s.Cursors,
		NextPageNumber: pageNumber,
		RemainingSteps: s.RemainingSteps,
		PageSize:       s.PageSize,
	}
	return ns, true, nil
}

func (s *State) NextStep() (newState *State, hasMore bool, err error) {
	if len(s.RemainingSteps) == 0 {
		newState = &State{
			CurrentStep:    "",
			Cursors:        s.Cursors,
			NextPageNumber: 0,
			RemainingSteps: []string{},
			PageSize:       s.PageSize,
		}
		hasMore = false
	} else {
		newState = &State{
			CurrentStep:    s.RemainingSteps[0],
			Cursors:        s.Cursors,
			NextPageNumber: 1,
			RemainingSteps: s.RemainingSteps[1:],
			PageSize:       s.PageSize,
		}
		hasMore = true
	}
	return
}

func (s *State) IncrementPage() (*State, bool, error) {
	if len(s.RemainingSteps) == 0 {
		return nil, false, nil
	}
	if s.CurrentStep == "" {
		return s.NextStep()
	}
	return &State{
		CurrentStep:    s.CurrentStep,
		Cursors:        s.Cursors,
		NextPageNumber: s.NextPageNumber + 1,
		RemainingSteps: s.RemainingSteps,
		PageSize:       s.PageSize,
	}, true, nil
}

func (s *State) Reset(steps []string, pageSize int) error {
	allSteps := steps
	if allSteps == nil || len(allSteps) < 1 {
		return fmt.Errorf("you should define at least one entitiy to collect in the secrets")
	}
	s.CurrentStep = allSteps[0]
	s.NextPageNumber = 1
	s.RemainingSteps = allSteps[1:]
	s.PageSize = pageSize
	return nil
}

func (s *State) GetAllCursors() []*Cursor {
	if s.Cursors == nil {
		return []*Cursor{}
	}
	allCursors := make([]*Cursor, len(s.Cursors))
	i := 0
	for _, c := range s.Cursors {
		allCursors[i] = c
		i++
	}
	return allCursors
}
