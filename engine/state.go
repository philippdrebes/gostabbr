package engine

import "errors"

type Turn int8
type Phase int8

const (
	Spring Turn = iota
	Fall
	Winter
)

const (
	Order Phase = iota
	Retreat
	Build
)

type Country struct {
	Name        string
	HomeCenters []string
}

type State struct {
	Turn      Turn
	Phase     Phase
	Countries [7]Country
	World     *Graph
}

func (s *State) nextPhase() error {
	switch s.Turn {
	case Spring, Fall:
		if s.Phase == Order {
			s.Phase = Retreat
		} else {
			if s.Turn == Spring {
				s.Turn = Fall
				s.Phase = Order
			} else {
				s.Turn = Winter
				s.Phase = Build
			}
		}
	case Winter:
		s.Turn = Spring
		s.Phase = Order
	default:
		return errors.New("Unsupported Turn")
	}
	return nil
}
