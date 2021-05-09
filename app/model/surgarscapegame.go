package model

import (
	"math"
	"math/rand"
	"time"

	"github.com/hculpan/go-sdl-lib/component"
)

type Cell struct {
	Sugar    int
	Capacity int

	Occupant *Agent
}

type SugarscapeGame struct {
	component.BaseGame

	BoardWidth  int
	BoardHeight int
	Cycle       int
	Running     bool

	Board  [][]Cell
	Agents map[string]*Agent
}

var Sugarscape *SugarscapeGame

func NewSugarscape(worldWidth, worldHeight int) *SugarscapeGame {
	rand.Seed(time.Now().UnixNano())

	result := SugarscapeGame{
		BoardWidth:  worldWidth,
		BoardHeight: worldHeight,
	}

	Sugarscape = &result

	result.Board = make([][]Cell, worldWidth)
	result.Agents = make(map[string]*Agent)

	result.Reset()

	return &result
}

func NewCell(capacity int) Cell {
	return Cell{
		Sugar:    capacity,
		Capacity: capacity,
		Occupant: nil,
	}
}

func (s *SugarscapeGame) Reset() {
	s.Cycle = 0
	s.Running = false

	s.Agents = map[string]*Agent{}

	for x := 0; x < s.BoardWidth; x++ {
		s.Board[x] = make([]Cell, s.BoardHeight)
		for y := 0; y < s.BoardHeight; y++ {
			s.Board[x][y] = NewCell(0)
		}
	}

	s.SugarPlacements()
	s.AgentPlacement(400)
}

func (s *SugarscapeGame) Start() {
	s.Running = true
}

func (s *SugarscapeGame) Stop() {
	s.Running = false
}

func (s *SugarscapeGame) AgentConsume(agent *Agent) {
	agent.Health += s.Board[agent.X][agent.Y].Sugar
	s.Board[agent.X][agent.Y].Sugar = 0
}

func (s *SugarscapeGame) AgentAge(agent *Agent) {
	agent.Health -= agent.Metabolism
}

func (s *SugarscapeGame) RemoveAgent(agent *Agent) {
	s.Board[agent.X][agent.Y].Occupant = nil
	delete(s.Agents, agent.ID)
}

func (s SugarscapeGame) calculateDestination(dir, x, y int) (int, int) {
	switch dir {
	case 0:
		if y == 0 || s.Board[x][y-1].Occupant != nil {
			return x, y
		}
		return x, y - 1
	case 1:
		if x >= s.BoardWidth-1 || s.Board[x+1][y].Occupant != nil {
			return x, y
		}
		return x + 1, y
	case 2:
		if y >= s.BoardHeight-1 || s.Board[x][y+1].Occupant != nil {
			return x, y
		}
		return x, y + 1
	case 3:
		if x == 0 || s.Board[x-1][y].Occupant != nil {
			return x, y
		}
		return x - 1, y
	}

	return x, y
}

func (s *SugarscapeGame) MoveAgent(agent *Agent) {
	dirmod := rand.Intn(4)
	for cnt := 0; cnt < 4; cnt++ {
		dir := (cnt + dirmod) % 4
		x, y := s.calculateDestination(dir, agent.X, agent.Y)
		if x != agent.X || y != agent.Y {
			s.MoveAgentTo(agent, x, y)
			break
		}
	}
}

func (s *SugarscapeGame) Update() error {
	if !s.Running {
		return nil
	}

	s.Cycle++

	for _, agent := range s.Agents {
		agent.ConsumeFood(s.Board[agent.X][agent.Y].Sugar)
		s.Board[agent.X][agent.Y].Sugar = 0

		agent.Metabolize()

		if agent.Health < 1 {
			s.RemoveAgent(agent)
		} else {
			s.MoveAgent(agent)
		}
	}

	if len(s.Agents) == 0 {
		s.Stop()
	} else {
		s.ReplenishSugar()
	}

	return nil
}

func (s *SugarscapeGame) ReplenishSugar() {
	for x := 0; x < s.BoardWidth; x++ {
		for y := 0; y < s.BoardHeight; y++ {
			if s.Board[x][y].Sugar < s.Board[x][y].Capacity {
				s.Board[x][y].Sugar++
			}
		}
	}
}

func (s *SugarscapeGame) SugarPlacements() {
	var x1, y1, x2, y2 float64 = 20, 80, 80, 20
	for x := 0; x < s.BoardWidth; x++ {
		for y := 0; y < s.BoardHeight; y++ {
			d1 := 4 - distance(x1, y1, float64(x), float64(y))/14
			d2 := 4 - distance(x2, y2, float64(x), float64(y))/14
			if d1 < 0 {
				d1 = 0
			}
			if d2 < 0 {
				d2 = 0
			}
			v := int(math.Round(math.Max(d1, d2)))

			s.Board[x][y].Capacity = v
			s.Board[x][y].Sugar = v
		}
	}
}

func (s *SugarscapeGame) AgentPlacement(number int) {
	for n := 0; n < number; n++ {
		agent := NewAgent()
		s.Agents[agent.ID] = &agent
		s.Board[agent.X][agent.Y].Occupant = &agent
	}
}

func distance(x1, y1, x2, y2 float64) float64 {
	xsqr := math.Pow(x2-x1, 2)
	ysqr := math.Pow(y2-y1, 2)
	return math.Sqrt(xsqr + ysqr)
}

func (s *SugarscapeGame) MoveAgentTo(agent *Agent, x, y int) bool {
	if s.Board[x][y].Occupant != nil {
		return false
	}

	s.Board[agent.X][agent.Y].Occupant = nil
	s.Board[x][y].Occupant = agent
	agent.X, agent.Y = x, y
	return true
}
