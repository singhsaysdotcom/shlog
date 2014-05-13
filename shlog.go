// shlog - colorful shell logger for command line apps
package shlog

import (
	"fmt"
	"regexp"
)

type Color string

func (c Color) String() string {
	return string(c)
}

type Symbol string

func (s Symbol) String() string {
	return string(s)
}

const (
	// Colors
	Reset  Color = "\x1b[0m"
	Black  Color = "\x1b[30m"
	Red    Color = "\x1b[31m"
	Green  Color = "\x1b[32m"
	Orange Color = "\x1b[33m"
	Purple Color = "\x1b[34m"
	Pink   Color = "\x1b[35m"
	Cyan   Color = "\x1b[36m"
	White  Color = "\x1b[37m"
	Grey   Color = "\x1b[90m"

	// Symbols
	Arrow    Symbol = "\u27AF"
	ThumbsUp Symbol = "\U0001F44D"
)

var (
	escape_re = regexp.MustCompile("\x1b\\[\\d+m")
)

type Logger struct {
	Padding             int
	lastMessageLen      int
	MessagePrefixColor  Color
	MessageOkColor      Color
	MessageErrorColor   Color
	MessagePrefixSymbol Symbol
	DoneSymbol          Symbol
	MessageOkText       string
	MessageErrorText    string
	MessageDoneText     string
	StatusLeftDelim     string
	StatusRightDelim    string
}

func NewLogger() *Logger {
	return &Logger{
		Padding:             70,
		MessagePrefixColor:  Cyan,
		MessageOkColor:      Green,
		MessageErrorColor:   Red,
		MessagePrefixSymbol: Arrow,
		MessageOkText:       "ok",
		MessageErrorText:    "err",
		MessageDoneText:     "All done",
		DoneSymbol:          ThumbsUp,
		StatusLeftDelim:     "[",
		StatusRightDelim:    "]",
	}
}

func (l *Logger) Message(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Printf("%s %s\t%s%s", l.MessagePrefixColor, l.MessagePrefixSymbol, Reset, msg)
	l.lastMessageLen = len(escape_re.ReplaceAllString(msg, ""))
}

func (l *Logger) Status(color Color, text string) {
	fmt.Printf("%*s%s%s%s%s\n", l.Padding-l.lastMessageLen, l.StatusLeftDelim, color, text, Reset, l.StatusRightDelim)
}

func (l *Logger) Ok() {
	l.Status(l.MessageOkColor, l.MessageOkText)
}

func (l *Logger) Err() {
	l.Status(l.MessageErrorColor, l.MessageErrorText)
}

func (l *Logger) Done() {
	fmt.Printf("\n %s\t%s\n\n", l.DoneSymbol, l.MessageDoneText)
}
