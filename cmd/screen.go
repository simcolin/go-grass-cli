package cmd

const MOVE_TOP_LEFT = "\033[H"
const CLEAR_SCREEN = "\033[2J"
const RESET = "\033[0m"
const GREEN = "\033[32m"
const BG_GREEN = "\033[42m"
const LIGHT_GREEN = "\033[92m"
const BG_LIGHT_GREEN = "\033[102m"
const BLACK = "\033[30m"
const BG_BLACK = "\033[40m"

type Screen struct {
	width, height int
	buffer        []string
}

func (s *Screen) Write(x, y int, text string) {
	if x >= 0 && x < s.width && y >= 0 && y < s.height {
		s.buffer[y*s.width+x] = text
	}
}

func (s *Screen) Resize(width, height int) {
	s.width = width
	s.height = height
	s.buffer = make([]string, width*height)
}

func (s *Screen) Clear() {
	for i := range s.buffer {
		s.buffer[i] = BG_BLACK
	}
}

func (s *Screen) String() string {
	str := MOVE_TOP_LEFT
	for i := 0; i < len(s.buffer); i++ {
		if i > 0 && s.buffer[i] == s.buffer[i-1] {
			if s.width%2 == 0 && i%s.width == s.width-1 {
				str += "   "
			} else {
				str += "  "
			}
		} else {
			if s.width%2 == 0 && i%s.width == s.width-1 {
				str += s.buffer[i] + "   "
			} else {
				str += s.buffer[i] + "  "
			}
		}
	}
	return str
}

func NewScreen(width, height int) *Screen {
	screen := &Screen{
		width:  width,
		height: height,
		buffer: make([]string, width*height),
	}
	return screen
}
