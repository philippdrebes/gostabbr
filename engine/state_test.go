package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func InitializeTestGame() (*State, error) {
	austria := &Country{Name: "Austria", HomeCenters: []string{"Vie", "Bud"}}
	italy := &Country{Name: "Italy", HomeCenters: []string{"Rom", "Ven"}}
	turkey := &Country{Name: "Turkey", HomeCenters: []string{}}

	game := &State{
		Turn:      Spring,
		Phase:     OrderPhase,
		Countries: []*Country{austria, italy, turkey},
		World:     initializeTestWorld(),
	}

	for _, c := range game.Countries {
		for _, hc := range c.HomeCenters {
			p, err := game.World.GetProvince(hc)
			if err != nil {
				return nil, err
			}

			p.OwnedBy = hc
		}
	}

	game.World.AddUnit(austria, Army, "Vie")
	game.World.AddUnit(austria, Army, "Bud")

	game.World.AddUnit(italy, Army, "Rom")
	game.World.AddUnit(italy, Army, "Ven")

	game.World.AddUnit(turkey, Fleet, "ION")

	return game, nil
}

func initializeTestWorld() *Graph {
	g := &Graph{Provinces: map[string]*Province{}}

	g.AddProvince("Vie", "Vienna", LandTile, true)
	g.AddProvince("Bud", "Budapest", LandTile, true)
	g.AddProvince("Tri", "Trieste", LandTile, false)
	g.AddProvince("Ven", "Venice", LandTile, true)
	g.AddProvince("Rom", "Rome", LandTile, true)
	g.AddProvince("Tyr", "Tyrolia", LandTile, false)
	g.AddProvince("ADR", "Adriatic Sea", WaterTile, false)
	g.AddProvince("ION", "Ionian Sea", WaterTile, false)

	g.AddEdges("Vie", []string{"Bud", "Tri", "Tyr"})
	g.AddEdges("Bud", []string{"Vie", "Tri"})
	g.AddEdges("Tri", []string{"Vie", "Bud", "Tyr", "Ven", "ADR"})
	g.AddEdges("Ven", []string{"Tri", "Tyr", "Rom", "ADR"})
	g.AddEdges("Rom", []string{"Ven"})
	g.AddEdges("Tyr", []string{"Vie", "Tri", "Ven"})
	g.AddEdges("ADR", []string{"Tri", "Ven"})
	g.AddEdges("ION", []string{"ADR"})

	return g
}

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

func setupStateWithGraph() *State {
	provinces := map[string]*Province{
		"Paris":     {Key: "Paris"},
		"Berlin":    {Key: "Berlin"},
		"Munich":    {Key: "Munich"},
		"Edinburgh": {Key: "Edinburgh"},
	}

	france := &Country{Name: "France"}
	germany := &Country{Name: "Germany"}
	england := &Country{Name: "England"}

	graph := &Graph{Provinces: provinces}
	state := &State{
		Turn:      Spring,
		Phase:     OrderPhase,
		Countries: []*Country{france, germany, england},
		World:     graph,
	}

	state.World.AddUnit(france, Army, "Paris")
	state.World.AddUnit(germany, Army, "Berlin")
	state.World.AddUnit(england, Fleet, "Edinburgh")

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

func TestAdjudicate_HoldAndMoveOrders(t *testing.T) {
	// Hold order and move orders do not conflict
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

func TestCalculateStrength_HoldOrder(t *testing.T) {
	state, err := InitializeTestGame()
	assert.NoError(t, err)

	state.AddHoldOrder("Austria", "Vie")
	country, err := state.GetCountry("Austria")
	assert.NoError(t, err)

	assert.Len(t, country.orders, 1)
	strength := calculateStrength(country.orders[0], state.World)
	assert.Equal(t, 1, strength)
}

func TestCalculateStrength_HoldOrderWithSupport(t *testing.T) {
	state, err := InitializeTestGame()
	assert.NoError(t, err)

	state.AddHoldOrder("Austria", "Vie")
	state.AddSupportOrder("Austria", "Bud", "Vie", "Vie")

	country, err := state.GetCountry("Austria")
	assert.NoError(t, err)

	assert.Len(t, country.orders, 2)
	strength := calculateStrength(country.orders[0], state.World)
	assert.Equal(t, 2, strength)
}

func TestCalculateStrength_MoveOrder(t *testing.T) {
	state, err := InitializeTestGame()
	assert.NoError(t, err)

	state.AddMoveOrder("Austria", "Vie", "Tri")
	country, err := state.GetCountry("Austria")
	assert.NoError(t, err)

	assert.Len(t, country.orders, 1)
	strength := calculateStrength(country.orders[0], state.World)
	assert.Equal(t, 1, strength)
}

func TestCalculateStrength_MoveOrderWithSupport(t *testing.T) {
	state, err := InitializeTestGame()
	assert.NoError(t, err)

	state.AddMoveOrder("Austria", "Vie", "Tri")
	state.AddSupportOrder("Austria", "Bud", "Vie", "Tri")
	country, err := state.GetCountry("Austria")
	assert.NoError(t, err)

	assert.Len(t, country.orders, 2)
	strength := calculateStrength(country.orders[0], state.World)
	assert.Equal(t, 2, strength)
}

func TestCalculateStrength_MoveOrderWithSupportFromOther(t *testing.T) {
	state, err := InitializeTestGame()
	assert.NoError(t, err)

	state.AddMoveOrder("Austria", "Vie", "Tri")
	state.AddSupportOrder("Austria", "Bud", "Vie", "Tri")
	state.AddSupportOrder("Italy", "Ven", "Vie", "Tri")

	austria, err := state.GetCountry("Austria")
	assert.NoError(t, err)

	italy, err := state.GetCountry("Italy")
	assert.NoError(t, err)

	assert.Len(t, austria.orders, 2)
	assert.Len(t, italy.orders, 1)
	strength := calculateStrength(austria.orders[0], state.World)
	assert.Equal(t, 3, strength)
}

func TestCalculateStrength_MoveOrderWithoutSupportFromOther(t *testing.T) {
	state, err := InitializeTestGame()
	assert.NoError(t, err)

	state.AddMoveOrder("Austria", "Vie", "Tri")
	state.AddSupportOrder("Austria", "Bud", "Vie", "Tri")
	state.AddSupportOrder("Italy", "Ven", "Vie", "Bud")

	austria, err := state.GetCountry("Austria")
	assert.NoError(t, err)

	italy, err := state.GetCountry("Italy")
	assert.NoError(t, err)

	assert.Len(t, austria.orders, 2)
	assert.Len(t, italy.orders, 1)
	strength := calculateStrength(austria.orders[0], state.World)
	assert.Equal(t, 2, strength)
}

func TestCalculateStrength_ArmyCannotSupportWaterMove(t *testing.T) {
	state, err := InitializeTestGame()
	assert.NoError(t, err)

	state.AddMoveOrder("Turkey", "ION", "ADR")
	state.AddSupportOrder("Italy", "Ven", "ION", "ADR")

	turkey, err := state.GetCountry("Turkey")
	assert.NoError(t, err)

	italy, err := state.GetCountry("Italy")
	assert.NoError(t, err)

	assert.Len(t, turkey.orders, 1)
	assert.Len(t, italy.orders, 1)
	strength := calculateStrength(turkey.orders[0], state.World)
	assert.Equal(t, 1, strength)
}

func TestIsValidSupportOrder(t *testing.T) {
	paris := &Province{Key: "Paris", Type: LandTile}
	london := &Province{Key: "London", Type: LandTile}
	northSea := &Province{Key: "North Sea", Type: WaterTile}

	tests := []struct {
		name        string
		order       Order
		support     Order
		supportUnit Unit
		expected    bool
	}{
		{
			name: "Valid Army Support on Land",
			order: MoveOrder{
				Position:    paris,
				Destination: london,
			},
			support: SupportOrder{
				Source:      paris,
				Destination: london,
			},
			supportUnit: Unit{Type: Army},
			expected:    true,
		},
		{
			name: "Invalid Army Support to Water",
			order: MoveOrder{
				Position:    paris,
				Destination: northSea,
			},
			support: SupportOrder{
				Source:      paris,
				Destination: northSea,
			},
			supportUnit: Unit{Type: Army},
			expected:    false,
		},
		{
			name: "Valid Fleet Support on Water",
			order: MoveOrder{
				Position:    paris,
				Destination: northSea,
			},
			support: SupportOrder{
				Source:      paris,
				Destination: northSea,
			},
			supportUnit: Unit{Type: Fleet},
			expected:    true,
		},
		{
			name: "Valid Fleet Support to Land",
			order: MoveOrder{
				Position:    paris,
				Destination: london,
			},
			support: SupportOrder{
				Source:      paris,
				Destination: london,
			},
			supportUnit: Unit{Type: Fleet},
			expected:    true,
		},
		{
			name: "Invalid Support Different Destination",
			order: MoveOrder{
				Position:    paris,
				Destination: london,
			},
			support: SupportOrder{
				Source:      paris,
				Destination: northSea,
			},
			supportUnit: Unit{Type: Army},
			expected:    false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := isValidSupportOrder(tc.order, tc.support, tc.supportUnit)
			assert.Equal(t, tc.expected, result, "Test case: %s", tc.name)
		})
	}
}
