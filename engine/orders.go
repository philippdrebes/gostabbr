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
	Position *Province
	Dest     *Province
}

type SupportOrder struct {
	Position *Province
	Src      *Province
	Dest     *Province
}

type ConvoyOrder struct {
	Position *Province
	Src      *Province
	Dest     *Province
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
	return fmt.Sprintf("%s %s - %s", m.Position.Unit.Type, m.Position.Key, m.Dest.Key)
}

func (m MoveOrder) GetPosition() *Province {
	return m.Position
}

func (m MoveOrder) GetSource() *Province {
	return m.Position
}

func (m MoveOrder) GetDestination() *Province {
	return m.Dest
}

func (m MoveOrder) Move() error {
	if m.Dest.Unit != nil {
		return errors.New("Destination occupied")
	}
	m.Dest.Unit = m.Position.Unit
	m.Position.Unit = nil
	return nil
}

func (s SupportOrder) String() string {
	if s.Src == s.Dest {
		return fmt.Sprintf("%s %s S %s", s.Position.Unit.Type, s.Position.Key, s.Src.Key)
	}
	return fmt.Sprintf("%s %s S %s - %s", s.Position.Unit.Type, s.Position.Key, s.Src.Key, s.Dest.Key)
}

func (s SupportOrder) GetPosition() *Province {
	return s.Position
}

func (s SupportOrder) GetSource() *Province {
	return s.Src
}

func (s SupportOrder) GetDestination() *Province {
	return s.Dest
}

func (c ConvoyOrder) String() string {
	return fmt.Sprintf("%s %s C %s - %s", c.Position.Unit.Type, c.Position.Key, c.Src.Key, c.Dest.Key)
}

func (c ConvoyOrder) GetPosition() *Province {
	return c.Position
}

func (c ConvoyOrder) GetSource() *Province {
	return c.Src
}

func (c ConvoyOrder) GetDestination() *Province {
	return c.Dest
}

func (u UnitType) String() string {
	if u == Army {
		return "A"
	} else if u == Fleet {
		return "F"
	}
	return ""
}
