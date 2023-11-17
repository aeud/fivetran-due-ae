package fivetranio

import (
	"encoding/json"
	"fmt"
	"log"
)

type State struct {
	Debug           bool               `json:"debug"` // Debug will be used to shorten the pagination. It will only read the first page.
	CurrentStep     string             `json:"current_step"`
	StepProgression string             `json:"step_progression"`
	NextPageNumber  int                `json:"next_page_number"`
	Cursors         map[string]*Cursor `json:"cursors,omitempty"`
	RemainingSteps  []string           `json:"remaining_steps"`
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

func (s *State) NextPageWithCursor(totalPageNumber int) (*State, bool, error) {
	ns := &State{
		Debug:           s.Debug,
		CurrentStep:     s.CurrentStep,
		Cursors:         s.Cursors,
		StepProgression: fmt.Sprintf("Remaining: %d", totalPageNumber),
		NextPageNumber:  1,
		RemainingSteps:  s.RemainingSteps,
	}
	return ns, true, nil
}

func (s *State) NextPage(pageNumber, totalPageNumber int) (*State, bool, error) {
	ns := &State{
		Debug:           s.Debug,
		CurrentStep:     s.CurrentStep,
		Cursors:         s.Cursors,
		StepProgression: fmt.Sprintf("%d/%d", pageNumber-1, totalPageNumber),
		NextPageNumber:  pageNumber,
		RemainingSteps:  s.RemainingSteps,
	}
	return ns, true, nil
}

func (s *State) NextStep() (newState *State, hasMore bool, err error) {
	if len(s.RemainingSteps) == 0 {
		newState = &State{
			Debug:          s.Debug,
			CurrentStep:    "",
			Cursors:        s.Cursors,
			NextPageNumber: 0,
			RemainingSteps: []string{},
		}
		hasMore = false
	} else {
		newState = &State{
			Debug:          s.Debug,
			CurrentStep:    s.RemainingSteps[0],
			Cursors:        s.Cursors,
			NextPageNumber: 1,
			RemainingSteps: s.RemainingSteps[1:],
		}
		hasMore = true
	}
	return
}

func (s *State) Reset(steps []string) error {
	allSteps := steps
	if allSteps == nil || len(allSteps) < 1 {
		return fmt.Errorf("you should define at least one entitiy to collect in the secrets")
	}
	s.CurrentStep = allSteps[0]
	s.NextPageNumber = 1
	s.RemainingSteps = allSteps[1:]
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
func (s *State) ResetAllCursors() {
	s.Cursors = map[string]*Cursor{}
}
