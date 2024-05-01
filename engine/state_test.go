package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextPhase(t *testing.T) {
	tests := []struct {
		name          string
		initialTurn   Turn
		initialPhase  Phase
		expectedTurn  Turn
		expectedPhase Phase
	}{
		{"Spring Order to Spring Retreat", Spring, Order, Spring, Retreat},
		{"Spring Retreat to Fall Order", Spring, Retreat, Fall, Order},
		{"Fall Order to Fall Retreat", Fall, Order, Fall, Retreat},
		{"Fall Retreat to Winter Build", Fall, Retreat, Winter, Build},
		{"Winter Build to Spring Order", Winter, Build, Spring, Order},
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

func TestNextPhaseNegativeCases(t *testing.T) {
	tests := []struct {
		name         string
		initialTurn  Turn
		initialPhase Phase
	}{
		{"Invalid Turn", Turn(99), Build},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			state := State{Turn: test.initialTurn, Phase: test.initialPhase}
			err := state.nextPhase()
			assert.Error(t, err)
		})
	}
}
