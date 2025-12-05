package clr

// ANSI escape codes for text colors
const (
	TBlack   = "\033[30m"
	TRed     = "\033[31m"
	TGreen   = "\033[32m"
	TYellow  = "\033[33m"
	TBlue    = "\033[34m"
	TMagenta = "\033[35m"
	TCyan    = "\033[36m"
	TWhite   = "\033[37m"
	TReset   = "\033[0m"
)

// ANSI escape codes for background colors
const (
	BBlack   = "\033[40m"
	BRed     = "\033[41m"
	BGreen   = "\033[42m"
	BYellow  = "\033[43m"
	BBlue    = "\033[44m"
	BMagenta = "\033[45m"
	BCyan    = "\033[46m"
	BWhite   = "\033[47m"
)

// Generic helper to apply any ANSI code
func colorize(text, colorCode string) string  { return colorCode + text + TReset }
func Bg(text string, colorCode string) string { return colorCode + text + "\033[0m" }

// --------------------
// TEXT COLOR FUNCTIONS
// --------------------

func TextBlack(text string) string   { return colorize(text, TBlack) }
func TextRed(text string) string     { return colorize(text, TRed) }
func TextGreen(text string) string   { return colorize(text, TGreen) }
func TextYellow(text string) string  { return colorize(text, TYellow) }
func TextBlue(text string) string    { return colorize(text, TBlue) }
func TextMagenta(text string) string { return colorize(text, TMagenta) }
func TextCyan(text string) string    { return colorize(text, TCyan) }
func TextWhite(text string) string   { return colorize(text, TWhite) }

// ------------------------
// BACKGROUND COLOR FUNCTIONS
// ------------------------

func BgBlack(text string) string   { return colorize(text, BBlack) }
func BgRed(text string) string     { return colorize(text, BRed) }
func BgGreen(text string) string   { return colorize(text, BGreen) }
func BgYellow(text string) string  { return colorize(text, BYellow) }
func BgBlue(text string) string    { return colorize(text, BBlue) }
func BgMagenta(text string) string { return colorize(text, BMagenta) }
func BgCyan(text string) string    { return colorize(text, BCyan) }
func BgWhite(text string) string   { return colorize(text, BWhite) }
