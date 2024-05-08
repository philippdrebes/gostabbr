package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPhase(t *testing.T) {
	tests := []struct {
		name          string
		initialTurn   Turn
		initialPhase  Phase
		expectedTurn  Turn
		expectedPhase Phase
	}{
		{"Spring Order to Spring Retreat", Spring, OrderPhase, Spring, RetreatPhase},
		{"Spring Retreat to Fall Order", Spring, RetreatPhase, Fall, OrderPhase},
		{"Fall Order to Fall Retreat", Fall, OrderPhase, Fall, RetreatPhase},
		{"Fall Retreat to Winter Build", Fall, RetreatPhase, Winter, BuildPhase},
		{"Winter Build to Spring Order", Winter, BuildPhase, Spring, OrderPhase},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			state := State{Turn: test.initialTurn, Phase: test.initialPhase}
			state.nextPhase()
			assert.Equal(t, test.expectedTurn, state.Turn, "Turn should transition correctly in "+test.name)
			assert.Equal(t, test.expectedPhase, state.Phase, "Phase should transition correctly in "+test.name)
		})
	}
}

func TestNextPhase_NegativeCases(t *testing.T) {
	tests := []struct {
		name         string
		initialTurn  Turn
		initialPhase Phase
	}{
		{"Invalid Turn", Turn(99), BuildPhase},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			state := State{Turn: test.initialTurn, Phase: test.initialPhase}
			err := state.nextPhase()
			assert.Error(t, err)
		})
	}
}

func TestAddOrder_HoldOrder(t *testing.T) {
	country := &Country{}
	paris := &Province{Key: "Paris"}
	unit := &Unit{Type: Army}
	holdOrder := &HoldOrder{Unit: unit, Position: paris}
	country.AddOrder(holdOrder)
	assert.Equal(t, 1, len(country.orders), "Country should have one order after addition")
	assert.Equal(t, paris, country.orders[0].GetPosition(), "Order position should be Paris")
}

func TestAddOrder_ReplaceOrderWithMoveOrder(t *testing.T) {
	country := &Country{}
	paris := &Province{Key: "Paris"}
	unit := &Unit{Type: Army}
	holdOrder := &HoldOrder{Unit: unit, Position: paris}
	moveOrder := &MoveOrder{Unit: unit, Position: paris, Dest: &Province{Key: "Berlin"}}
	country.AddOrder(holdOrder)
	country.AddOrder(moveOrder)
	assert.Equal(t, 1, len(country.orders), "Country should still have one order after replacement")
	assert.IsType(t, &MoveOrder{}, country.orders[0], "The replaced order should be a MoveOrder")
	assert.Equal(t, "Berlin", country.orders[0].(*MoveOrder).Dest.Key, "Destination of MoveOrder should be Berlin")
}

func TestAddOrder_AddMultipleOrdersDifferentLocations(t *testing.T) {
	country := &Country{}
	paris := &Province{Key: "Paris"}
	berlin := &Province{Key: "Berlin"}
	unit := &Unit{Type: Army}
	holdOrderParis := &HoldOrder{Unit: unit, Position: paris}
	holdOrderBerlin := &HoldOrder{Unit: unit, Position: berlin}
	country.AddOrder(holdOrderParis)
	country.AddOrder(holdOrderBerlin)
	assert.Equal(t, 2, len(country.orders), "Country should have two orders for different locations")
	assert.NotEqual(t, country.orders[0].GetPosition(), country.orders[1].GetPosition(), "Orders should be in different provinces")
}

func TestAddOrder_NoDuplicateOnNonMatchingPositions(t *testing.T) {
	country := &Country{}
	paris := &Province{Key: "Paris"}
	berlin := &Province{Key: "Berlin"}
	unit := &Unit{Type: Army}
	holdOrderParis := &HoldOrder{Unit: unit, Position: paris}
	moveOrderBerlin := &MoveOrder{Unit: unit, Position: berlin, Dest: &Province{Key: "Munich"}}
	country.AddOrder(holdOrderParis)
	country.AddOrder(moveOrderBerlin)
	assert.Equal(t, 2, len(country.orders), "Country should have two distinct orders when positions don't match")
}
