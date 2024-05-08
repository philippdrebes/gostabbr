package engine

import (
	"errors"
)

type Turn int8
type Phase int8

const (
	Spring Turn = iota
	Fall
	Winter
)

const (
	OrderPhase Phase = iota
	RetreatPhase
	BuildPhase
)

type Country struct {
	Name        string
	HomeCenters []string
	orders      []Order
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
		if s.Phase == OrderPhase {
			s.Phase = RetreatPhase
		} else {
			if s.Turn == Spring {
				s.Turn = Fall
				s.Phase = OrderPhase
			} else {
				s.Turn = Winter
				s.Phase = BuildPhase
			}
		}
	case Winter:
		s.Turn = Spring
		s.Phase = OrderPhase
	default:
		return errors.New("Unsupported Turn")
	}
	return nil
}

func (this *Country) AddOrder(newOrder Order) {
	if this.orders == nil {
		this.orders = []Order{}
	}

	for index, existing := range this.orders {
		if newOrder.GetPosition().Key == existing.GetPosition().Key {
			this.orders[index] = newOrder
			return
		}
	}

	this.orders = append(this.orders, newOrder)
}
