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
