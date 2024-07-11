package cmd

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/ojrac/opensimplex-go"
)

type GameState struct {
	Width, Height       int
	Screen              *Screen
	Noise               opensimplex.Noise
	WindPosition        Vector3
	WindDirection       Vector3
	WindDirectionTarget Vector3
	Strands             []Strand
}

const STRAND_PER_PIXEL = 0.05

func NewGameState(width, height int) *GameState {
	strandCount := int(STRAND_PER_PIXEL * float64(width*height))
	return &GameState{
		Width:               width,
		Height:              height,
		Screen:              NewScreen(width, height),
		Noise:               opensimplex.NewNormalized(rand.Int63()),
		WindPosition:        Vector3{0, 0, 0},
		WindDirection:       Vector3{0, 0, 0},
		WindDirectionTarget: Vector3{2, 2, 0.1},
		Strands:             make([]Strand, strandCount),
	}
}

func (g *GameState) Resize(width, height int) {
	g.Width = width
	g.Height = height
	g.Screen.Resize(width, height)
	strandCount := int(STRAND_PER_PIXEL * float64(width*height))
	g.Strands = make([]Strand, strandCount)
	g.Setup()
}

func (g *GameState) Setup() {
	for i := range g.Strands {
		g.Strands[i] = NewStrand(g.Width, g.Height)
	}
	sort.SliceStable(g.Strands, func(i, j int) bool {
		return g.Strands[i].Y < g.Strands[j].Y
	})
}

func (g *GameState) Loop(deltaTime float64) {
	g.Screen.Clear()
	for _, strand := range g.Strands {
		noiseValue := g.Noise.Eval2(
			float64(strand.X)/WIND_SCALE+float64(g.WindPosition.X),
			float64(strand.Y)/WIND_SCALE+float64(g.WindPosition.Y),
		)
		strand.Update(0.5, noiseValue)
		strand.Draw(g.Screen)
	}
	fmt.Print(g.Screen)

	g.WindDirection = g.WindDirection.Lerp(g.WindDirectionTarget, 0.01)
	g.WindPosition = g.WindPosition.Add(g.WindDirection.Scale(deltaTime * 1.5))
}
