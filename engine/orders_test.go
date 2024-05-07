package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHoldOrder_String(t *testing.T) {
	unit := &Unit{Country: "USA", Type: Army}
	province := &Province{Key: "NY", Name: "New York"}
	holdOrder := HoldOrder{Unit: unit, Position: province}
	assert.Equal(t, "A NY H", holdOrder.String(), "HoldOrder string should match expected format.")
}

func TestMoveOrder_String(t *testing.T) {
	unit := &Unit{Country: "USA", Type: Fleet}
	position := &Province{Key: "NY", Name: "New York"}
	dest := &Province{Key: "CA", Name: "California"}
	moveOrder := MoveOrder{Unit: unit, Position: position, Dest: dest}
	assert.Equal(t, "F NY - CA", moveOrder.String(), "MoveOrder string should match expected format.")
}

func TestSupportOrder_String(t *testing.T) {
	unit := &Unit{Country: "USA", Type: Army}
	position := &Province{Key: "NY", Name: "New York"}
	src := &Province{Key: "WA", Name: "Washington"}
	dest := &Province{Key: "CA", Name: "California"}
	supportOrderSame := SupportOrder{Unit: unit, Position: position, Src: src, Dest: src}
	supportOrderDiff := SupportOrder{Unit: unit, Position: position, Src: src, Dest: dest}
	assert.Equal(t, "A NY S WA", supportOrderSame.String(), "SupportOrder string (same src and dest) should match expected format.")
	assert.Equal(t, "A NY S WA - CA", supportOrderDiff.String(), "SupportOrder string (different src and dest) should match expected format.")
}

func TestConvoyOrder_String(t *testing.T) {
	unit := &Unit{Country: "USA", Type: Fleet}
	position := &Province{Key: "NY", Name: "New York"}
	src := &Province{Key: "WA", Name: "Washington"}
	dest := &Province{Key: "CA", Name: "California"}
	convoyOrder := ConvoyOrder{Unit: unit, Position: position, Src: src, Dest: dest}
	assert.Equal(t, "F NY C WA - CA", convoyOrder.String(), "ConvoyOrder string should match expected format.")
}

func TestUnitType_String(t *testing.T) {
	assert.Equal(t, "A", Army.String(), "Army should return 'A'.")
	assert.Equal(t, "F", Fleet.String(), "Fleet should return 'F'.")
}
