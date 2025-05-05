package color

const (
	FgBlack   = "30"
	FgRed     = "31"
	FgGreen   = "32"
	FgYellow  = "33"
	FgBlue    = "34"
	FgMagenta = "35"
	FgCyan    = "36"
	FgWhite   = "37"
)

const (
	BgBlack   = "40"
	BgRed     = "41"
	BgGreen   = "42"
	BgYellow  = "43"
	BgBlue    = "44"
	BgMagenta = "45"
	BgCyan    = "46"
	BgWhite   = "47"
)

const (
	FgBrightBlack   = "90"
	FgBrightRed     = "91"
	FgBrightGreen   = "92"
	FgBrightYellow  = "93"
	FgBrightBlue    = "94"
	FgBrightMagenta = "95"
	FgBrightCyan    = "96"
	FgBrightWhite   = "97"
)

const (
	BgBrightBlack   = "100"
	BgBrightRed     = "101"
	BgBrightGreen   = "102"
	BgBrightYellow  = "103"
	BgBrightBlue    = "104"
	BgBrightMagenta = "105"
	BgBrightCyan    = "106"
	BgBrightWhite   = "107"
)

var (
	FgColor16Sequence       = [8]string{"\x1b[30m", "\x1b[31m", "\x1b[32m", "\x1b[33m", "\x1b[34m", "\x1b[35m", "\x1b[36m", "\x1b[37m"}
	BgColor16Sequence       = [8]string{"\x1b[40m", "\x1b[41m", "\x1b[42m", "\x1b[43m", "\x1b[44m", "\x1b[45m", "\x1b[46m", "\x1b[47m"}
	FgBrightColor16Sequence = [8]string{"\x1b[90m", "\x1b[91m", "\x1b[92m", "\x1b[93m", "\x1b[94m", "\x1b[95m", "\x1b[96m", "\x1b[97m"}
	BgBrightColor16Sequence = [8]string{"\x1b[100m", "\x1b[101m", "\x1b[102m", "\x1b[103m", "\x1b[104m", "\x1b[105m", "\x1b[106m", "\x1b[107m"}
)
