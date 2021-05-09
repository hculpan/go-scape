package model

import (
	"fmt"
	"math/rand"
)

type Agent struct {
	ID         string
	X          int
	Y          int
	Health     int
	Vision     int
	Metabolism int
	Age        int
}

func NewAgent() Agent {
	result := Agent{}

	cnt := 0
	for {
		cnt++
		x, y := rand.Intn(Sugarscape.BoardWidth), rand.Intn(Sugarscape.BoardHeight)
		if Sugarscape.Board[x][y].Occupant == nil {
			result.X, result.Y = x, y
			break
		} else if cnt > 100 {
			panic("unable to place agent")
		}
	}

	result.Health = rand.Intn(10) + 20
	result.Vision = rand.Intn(4) + 1
	result.Metabolism = rand.Intn(4) + 1
	result.ID = fmt.Sprintf("agent-%05d", len(Sugarscape.Agents)+1)

	return result
}

func (a Agent) ToString() string {
	return fmt.Sprintf("ID: %s, Pos: [%d, %d], Health: %d, Vision: %d, Metabolism: %d", a.ID, a.X, a.Y, a.Health, a.Vision, a.Metabolism)
}

func (a *Agent) Metabolize() {
	a.Health -= a.Metabolism
	a.Age++
}

func (a *Agent) ConsumeFood(amount int) {
	a.Health += amount
}
