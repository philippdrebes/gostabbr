package engine

import (
	"errors"
	"fmt"
	"log"
	"reflect"
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
	Countries [7]*Country
	World     *Graph
}

func (s *State) GetCountry(country string) (*Country, error) {
	for _, c := range s.Countries {
		if c == nil {
			continue
		}
		if c.Name == country {
			return c, nil
		}
	}
	return nil, errors.New("Country does not exist")
}

func (s *State) AddHoldOrder(country string, position string) error {
	c, err := s.GetCountry(country)
	if err != nil {
		return err
	}

	pos, err := s.World.GetProvince(position)
	if err != nil {
		return err
	}

	err = s.addOrder(c, &HoldOrder{Position: pos})
	if err != nil {
		return err
	}

	return nil
}

func (s *State) AddMoveOrder(country, position, destination string) error {
	c, err := s.GetCountry(country)
	if err != nil {
		return err
	}

	pos, err := s.World.GetProvince(position)
	if err != nil {
		return err
	}

	dest, err := s.World.GetProvince(destination)
	if err != nil {
		return err
	}
	err = s.addOrder(c, &MoveOrder{Position: pos, Dest: dest})
	if err != nil {
		return err
	}

	return nil
}

func (s *State) AddSupportOrder(country, position, source, destination string) error {
	c, err := s.GetCountry(country)
	if err != nil {
		return err
	}

	pos, err := s.World.GetProvince(position)
	if err != nil {
		return err
	}

	src, err := s.World.GetProvince(source)
	if err != nil {
		return err
	}

	dest, err := s.World.GetProvince(destination)
	if err != nil {
		return err
	}
	err = s.addOrder(c, &SupportOrder{Position: pos, Src: src, Dest: dest})
	if err != nil {
		return err
	}

	return nil
}

func (s *State) AddConvoyOrder(country, position, source, destination string) error {
	c, err := s.GetCountry(country)
	if err != nil {
		return err
	}

	pos, err := s.World.GetProvince(position)
	if err != nil {
		return err
	}

	src, err := s.World.GetProvince(source)
	if err != nil {
		return err
	}

	dest, err := s.World.GetProvince(destination)
	if err != nil {
		return err
	}
	err = s.addOrder(c, &ConvoyOrder{Position: pos, Src: src, Dest: dest})
	if err != nil {
		return err
	}

	return nil
}

func (s *State) addOrder(country *Country, newOrder Order) error {
	if country.orders == nil {
		country.orders = []Order{}
	}

	for index, existing := range country.orders {
		if newOrder.GetPosition().Key == existing.GetPosition().Key {
			country.orders[index] = newOrder
			return nil // todo: do not return
		}
	}

	country.orders = append(country.orders, newOrder)

	return nil
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

func (s *State) Adjudicate() error {
	log.Println("Adjudication starting...")
	log.Println("Collecting orders")
	orders := []Order{}
	for _, country := range s.Countries {
		if country == nil {
			continue
		}
		log.Printf("Collecting %ss orders", country.Name)
		for _, order := range country.orders {
			if order == nil {
				continue
			}
			orders = append(orders, order)
		}
	}

	log.Printf("Processing %d orders in total", len(orders))
	for _, order := range orders {
		log.Printf("Processing order %s", order)
		switch o := order.(type) {
		case *HoldOrder:
			log.Printf("%s", o)
		case *MoveOrder:
			log.Printf("%s", o)
			err := o.Move()
			if err != nil {
				return err
			}
		default:
			return errors.New(fmt.Sprintf("Type %s is not supported", reflect.TypeOf(o)))
		}
	}

	// return s.nextPhase()
	return nil
}

func calculateStrength(order Order, world *Graph) int {
	var p *Province

	if _, ok := order.(MoveOrder); ok {
		// if move: get neighbors of destination and check for supporting orders
		p = order.GetDestination()
	} else {
		// if hold, support or convoy: get neighbors of source and check for supporting orders
		p = order.GetSource()
	}

	neighbors, err := world.GetNeighborsWithUnits(p)
	if err != nil {
		return 1
	}

	fmt.Printf("Found %d neighbors with units on them", len(neighbors))

	return 1
}
