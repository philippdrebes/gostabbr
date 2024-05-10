package engine

import "fmt"

type Order interface {
	GetPosition() *Province
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

func (m MoveOrder) String() string {
	return fmt.Sprintf("%s %s - %s", m.Position.Unit.Type, m.Position.Key, m.Dest.Key)
}

func (m MoveOrder) GetPosition() *Province {
	return m.Position
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

func (c ConvoyOrder) String() string {
	return fmt.Sprintf("%s %s C %s - %s", c.Position.Unit.Type, c.Position.Key, c.Src.Key, c.Dest.Key)
}

func (c ConvoyOrder) GetPosition() *Province {
	return c.Position
}

func (u UnitType) String() string {
	if u == Army {
		return "A"
	} else if u == Fleet {
		return "F"
	}
	return ""
}
