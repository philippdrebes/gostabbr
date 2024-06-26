package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHoldOrder_String(t *testing.T) {
	usa := &Country{Name: "USA"}
	unit := &Unit{Country: usa, Type: Army}
	province := &Province{Key: "NY", Name: "New York", Unit: unit}
	holdOrder := HoldOrder{Position: province}
	assert.Equal(t, "A NY H", holdOrder.String(), "HoldOrder string should match expected format.")
}

func TestMoveOrder_String(t *testing.T) {
	usa := &Country{Name: "USA"}
	unit := &Unit{Country: usa, Type: Fleet}
	position := &Province{Unit: unit, Key: "NY", Name: "New York"}
	dest := &Province{Key: "CA", Name: "California"}
	moveOrder := MoveOrder{Position: position, Destination: dest}
	assert.Equal(t, "F NY - CA", moveOrder.String(), "MoveOrder string should match expected format.")
}

func TestSupportOrder_String(t *testing.T) {
	usa := &Country{Name: "USA"}
	unit := &Unit{Country: usa, Type: Army}
	position := &Province{Unit: unit, Key: "NY", Name: "New York"}
	src := &Province{Key: "WA", Name: "Washington"}
	dest := &Province{Key: "CA", Name: "California"}
	supportOrderSame := SupportOrder{Position: position, Source: src, Destination: src}
	supportOrderDiff := SupportOrder{Position: position, Source: src, Destination: dest}
	assert.Equal(t, "A NY S WA", supportOrderSame.String(), "SupportOrder string (same src and dest) should match expected format.")
	assert.Equal(t, "A NY S WA - CA", supportOrderDiff.String(), "SupportOrder string (different src and dest) should match expected format.")
}

func TestConvoyOrder_String(t *testing.T) {
	usa := &Country{Name: "USA"}
	unit := &Unit{Country: usa, Type: Fleet}
	position := &Province{Unit: unit, Key: "NY", Name: "New York"}
	src := &Province{Key: "WA", Name: "Washington"}
	dest := &Province{Key: "CA", Name: "California"}
	convoyOrder := ConvoyOrder{Position: position, Source: src, Destination: dest}
	assert.Equal(t, "F NY C WA - CA", convoyOrder.String(), "ConvoyOrder string should match expected format.")
}

func TestUnitType_String(t *testing.T) {
	assert.Equal(t, "A", Army.String(), "Army should return 'A'.")
	assert.Equal(t, "F", Fleet.String(), "Fleet should return 'F'.")
}
