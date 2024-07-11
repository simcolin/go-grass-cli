package cmd

import (
	"math"
	"math/rand"
)

const WIND_SCALE = 20.0
const WIND_POWER = 0.5

const STRAND_MIN_LENGTH = 4
const STRAND_MAX_LENGTH = 10

var GREEN_PALETTE = []string{
	"\033[48;5;22m",
	"\033[48;5;28m",
	"\033[48;5;34m",
	"\033[48;5;40m",
	"\033[48;5;46m",
	"\033[48;5;83m",
	"\033[48;5;120m",
	"\033[48;5;157m",
}

type Strand struct {
	X          int
	Y          int
	Brightness []float64
	Positions  []int
}

func (s *Strand) Update(tilt float64, wind float64) {
	topStrandDarkness := (wind + (1-wind)*WIND_POWER)
	for i := (0); i < len(s.Brightness); i++ {
		offset := math.Round(tilt * wind * 0.5 * math.Pow(float64(i), 1.5))
		s.Positions[i*2] = s.X + int(offset)
		s.Positions[i*2+1] = s.Y - int(i)
		s.Brightness[i] = topStrandDarkness * (float64(i) / float64(len(s.Brightness)))
	}
}

func (s *Strand) Draw(screen *Screen) {
	for i, darkness := range s.Brightness {
		bg := GREEN_PALETTE[int(darkness*float64(len(GREEN_PALETTE)-1))]
		screen.Write(s.Positions[i*2], s.Positions[i*2+1], bg)
	}
}

func NewStrand(width, height int) Strand {
	length := STRAND_MIN_LENGTH + rand.Int()%(STRAND_MAX_LENGTH-STRAND_MIN_LENGTH)
	return Strand{
		X:          rand.Int() % width,
		Y:          rand.Int() % height,
		Brightness: make([]float64, length),
		Positions:  make([]int, length*2),
	}
}
