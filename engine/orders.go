package engine

import "fmt"

type HoldOrder struct {
	Unit     *Unit
	Position *Province
}

type MoveOrder struct {
	Unit     *Unit
	Position *Province
	Dest     *Province
}

type SupportOrder struct {
	Unit     *Unit
	Position *Province
	Src      *Province
	Dest     *Province
}

type ConvoyOrder struct {
	Unit     *Unit
	Position *Province
	Src      *Province
	Dest     *Province
}

func (this HoldOrder) String() string {
	return fmt.Sprintf("%s %s H", this.Unit.Type, this.Position.Key)
}

func (this MoveOrder) String() string {
	return fmt.Sprintf("%s %s - %s", this.Unit.Type, this.Position.Key, this.Dest.Key)
}

func (this SupportOrder) String() string {
	if this.Src == this.Dest {
		return fmt.Sprintf("%s %s S %s", this.Unit.Type, this.Position.Key, this.Src.Key)
	}
	return fmt.Sprintf("%s %s S %s - %s", this.Unit.Type, this.Position.Key, this.Src.Key, this.Dest.Key)
}

func (this ConvoyOrder) String() string {
	return fmt.Sprintf("%s %s C %s - %s", this.Unit.Type, this.Position.Key, this.Src.Key, this.Dest.Key)
}

func (this UnitType) String() string {
	if this == Army {
		return "A"
	} else if this == Fleet {
		return "F"
	}
	return ""
}
