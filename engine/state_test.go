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
	unit := &Unit{Type: Army}
	paris := &Province{Key: "Paris", Unit: unit}
	holdOrder := &HoldOrder{Position: paris}
	country.addOrder(holdOrder)
	assert.Equal(t, 1, len(country.orders), "Country should have one order after addition")
	assert.Equal(t, paris, country.orders[0].GetPosition(), "Order position should be Paris")
}

func TestAddOrder_ReplaceOrderWithMoveOrder(t *testing.T) {
	country := &Country{}
	unit := &Unit{Type: Army}
	paris := &Province{Unit: unit, Key: "Paris"}
	holdOrder := &HoldOrder{Position: paris}
	moveOrder := &MoveOrder{Position: paris, Dest: &Province{Key: "Berlin"}}
	country.addOrder(holdOrder)
	country.addOrder(moveOrder)
	assert.Equal(t, 1, len(country.orders), "Country should still have one order after replacement")
	assert.IsType(t, &MoveOrder{}, country.orders[0], "The replaced order should be a MoveOrder")
	assert.Equal(t, "Berlin", country.orders[0].(*MoveOrder).Dest.Key, "Destination of MoveOrder should be Berlin")
}

func TestAddOrder_AddMultipleOrdersDifferentLocations(t *testing.T) {
	country := &Country{}
	unit := &Unit{Type: Army}
	paris := &Province{Key: "Paris", Unit: unit}
	berlin := &Province{Key: "Berlin", Unit: unit}
	holdOrderParis := &HoldOrder{Position: paris}
	holdOrderBerlin := &HoldOrder{Position: berlin}
	country.addOrder(holdOrderParis)
	country.addOrder(holdOrderBerlin)
	assert.Equal(t, 2, len(country.orders), "Country should have two orders for different locations")
	assert.NotEqual(t, country.orders[0].GetPosition(), country.orders[1].GetPosition(), "Orders should be in different provinces")
}

func TestAddOrder_NoDuplicateOnNonMatchingPositions(t *testing.T) {
	country := &Country{}
	unit := &Unit{Type: Army}
	paris := &Province{Key: "Paris", Unit: unit}
	berlin := &Province{Key: "Berlin", Unit: unit}
	holdOrderParis := &HoldOrder{Position: paris}
	moveOrderBerlin := &MoveOrder{Position: berlin, Dest: &Province{Key: "Munich"}}
	country.addOrder(holdOrderParis)
	country.addOrder(moveOrderBerlin)
	assert.Equal(t, 2, len(country.orders), "Country should have two distinct orders when positions don't match")
}

func setupStateWithGraph() *State {
	provinces := map[string]*Province{
		"Paris":     {Key: "Paris"},
		"Berlin":    {Key: "Berlin"},
		"Munich":    {Key: "Munich"},
		"Edinburgh": {Key: "Edinburgh"},
	}
	graph := &Graph{Provinces: provinces}
	state := &State{
		Turn:  Spring,
		Phase: OrderPhase,
		Countries: [7]*Country{
			{Name: "France"},
			{Name: "Germany"},
			{Name: "England"},
		},
		World: graph,
	}

	state.World.AddUnit("France", Army, "Paris")
	state.World.AddUnit("Germany", Army, "Berlin")
	state.World.AddUnit("England", Fleet, "Edinburgh")

	return state
}

func TestAddHoldOrder_ValidInputs(t *testing.T) {
	s := setupStateWithGraph()
	err := s.AddHoldOrder("France", "Paris")
	assert.NoError(t, err, "Adding a valid hold order should not produce an error")
	assert.IsType(t, &HoldOrder{}, s.Countries[0].orders[0], "Order should be a HoldOrder")
}

func TestAddMoveOrder_ValidInputs(t *testing.T) {
	s := setupStateWithGraph()
	err := s.AddMoveOrder("France", "Paris", "Berlin")
	assert.NoError(t, err, "Adding a valid move order should not produce an error")
	assert.IsType(t, &MoveOrder{}, s.Countries[0].orders[0], "Order should be a MoveOrder")
}

func TestAddSupportOrder_ValidInputs(t *testing.T) {
	s := setupStateWithGraph()
	err := s.AddSupportOrder("France", "Paris", "Munich", "Berlin")
	assert.NoError(t, err, "Adding a valid support order should not produce an error")
	assert.IsType(t, &SupportOrder{}, s.Countries[0].orders[0], "Order should be a SupportOrder")
}

func TestAddConvoyOrder_ValidInputs(t *testing.T) {
	s := setupStateWithGraph()
	err := s.AddConvoyOrder("England", "Edinburgh", "Paris", "Munich")
	assert.NoError(t, err, "Adding a valid convoy order should not produce an error")
	assert.IsType(t, &ConvoyOrder{}, s.Countries[2].orders[0], "Order should be a ConvoyOrder")
}

func TestAddOrder_InvalidCountry(t *testing.T) {
	s := setupStateWithGraph()
	err := s.AddHoldOrder("Spain", "Paris")
	assert.Error(t, err, "Should return an error when adding an order to a nonexistent country")
}

func TestAddOrder_InvalidProvince(t *testing.T) {
	s := setupStateWithGraph()
	err := s.AddHoldOrder("France", "Vienna")
	assert.Error(t, err, "Should return an error when adding an order to a nonexistent province")
}

func TestAddOrder_InvalidDestination(t *testing.T) {
	s := setupStateWithGraph()
	err := s.AddMoveOrder("France", "Paris", "Vienna")
	assert.Error(t, err, "Should return an error when adding a move order with a nonexistent destination")
}

func TestAdjudicate_WithOrders(t *testing.T) {
	s := setupStateWithGraph()

	err := s.AddHoldOrder("France", "Paris")
	assert.NoError(t, err, "AddMoveOrder should complete without errors")
	err = s.AddMoveOrder("Germany", "Berlin", "Munich")
	assert.NoError(t, err, "AddMoveOrder should complete without errors")

	err = s.Adjudicate()
	assert.NoError(t, err, "Adjudicate should complete without error with orders")

	par, err := s.World.GetProvince("Paris")
	assert.Nil(t, err)
	assert.NotNil(t, par.Unit)

	ber, err := s.World.GetProvince("Berlin")
	mun, err := s.World.GetProvince("Munich")
	assert.Nil(t, err)
	assert.Nil(t, ber.Unit)
	assert.NotNil(t, mun.Unit)
}
