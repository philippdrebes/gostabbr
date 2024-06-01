package engine

import (
	"errors"
	"fmt"
)

type Order interface {
	fmt.Stringer
	GetPosition() *Province
	GetSource() *Province
	GetDestination() *Province
}

type HoldOrder struct {
	Position *Province
}

type MoveOrder struct {
	Position    *Province
	Destination *Province
}

type SupportOrder struct {
	Position    *Province
	Source      *Province
	Destination *Province
}

type ConvoyOrder struct {
	Position    *Province
	Source      *Province
	Destination *Province
}

func (h HoldOrder) String() string {
	return fmt.Sprintf("%s %s H", h.Position.Unit.Type, h.Position.Key)
}

func (h HoldOrder) GetPosition() *Province {
	return h.Position
}

func (h HoldOrder) GetSource() *Province {
	return h.Position
}

func (h HoldOrder) GetDestination() *Province {
	return h.Position
}

func (m MoveOrder) String() string {
	return fmt.Sprintf("%s %s - %s", m.Position.Unit.Type, m.Position.Key, m.Destination.Key)
}

func (m MoveOrder) GetPosition() *Province {
	return m.Position
}

func (m MoveOrder) GetSource() *Province {
	return m.Position
}

func (m MoveOrder) GetDestination() *Province {
	return m.Destination
}

func (m MoveOrder) Move() error {
	if m.Destination.Unit != nil {
		return errors.New("Destination occupied")
	}
	m.Destination.Unit = m.Position.Unit
	m.Position.Unit = nil
	return nil
}

func (s SupportOrder) String() string {
	if s.Source == s.Destination {
		return fmt.Sprintf("%s %s S %s", s.Position.Unit.Type, s.Position.Key, s.Source.Key)
	}
	return fmt.Sprintf("%s %s S %s - %s", s.Position.Unit.Type, s.Position.Key, s.Source.Key, s.Destination.Key)
}

func (s SupportOrder) GetPosition() *Province {
	return s.Position
}

func (s SupportOrder) GetSource() *Province {
	return s.Source
}

func (s SupportOrder) GetDestination() *Province {
	return s.Destination
}

func (c ConvoyOrder) String() string {
	return fmt.Sprintf("%s %s C %s - %s", c.Position.Unit.Type, c.Position.Key, c.Source.Key, c.Destination.Key)
}

func (c ConvoyOrder) GetPosition() *Province {
	return c.Position
}

func (c ConvoyOrder) GetSource() *Province {
	return c.Source
}

func (c ConvoyOrder) GetDestination() *Province {
	return c.Destination
}

func (u UnitType) String() string {
	if u == Army {
		return "A"
	} else if u == Fleet {
		return "F"
	}
	return ""
}
