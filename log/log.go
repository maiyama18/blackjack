package log

import (
	"fmt"
	"io"

	"github.com/mui87/blackjack/card"
)

type Logger struct {
	out io.Writer
}

func NewLogger(out io.Writer) *Logger {
	return &Logger{out: out}
}

func (l *Logger) Draw(name string, card card.Card, secret bool) {
	var cardStr string
	if secret {
		cardStr = "?"
	} else {
		cardStr = card.String()
	}
	_, _ = fmt.Fprintf(l.out, "[%s] draw %s", name, cardStr)
}
