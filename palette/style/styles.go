package style

const (
	Reset = "0"
)

const (
	Bold      = "1"
	Faint     = "2"
	Italic    = "3"
	Underline = "4"
	Blink     = "5"
	FastBlink = "6"
	Inverse   = "7"
	Hidden    = "8"
	Crossed   = "9"
)

var (
	StyleSequence = [10]string{"\x1b[0m", "\x1b[1m", "\x1b[2m", "\x1b[3m", "\x1b[4m", "\x1b[5m", "\x1b[6m", "\x1b[7m", "\x1b[8m", "\x1b[9m"}
)
