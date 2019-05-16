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

func (l *Logger) Logf(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(l.out, format, a...)
}

func (l *Logger) Draw(name string, card card.Card, secret bool) {
	var cardStr string
	if secret {
		cardStr = "?"
	} else {
		cardStr = card.String()
	}
	_, _ = fmt.Fprintf(l.out, "[%s] draw %s\n", name, cardStr)
}

func (l *Logger) Hand(name string, hand *card.Hand) {
	_, _ = fmt.Fprintf(l.out, "[%s] hand -> %s\n", name, hand)
}

func (l *Logger) Burst(name string, score int) {
	_, _ = fmt.Fprintf(l.out, "[%s] burst with score %d\n", name, score)
}

func (l *Logger) Scores(pName, dName string, pScore, dScore int) {
	_, _ = fmt.Fprintf(l.out, "[%s]: score -> %d\n", pName, pScore)
	_, _ = fmt.Fprintf(l.out, "[%s]: score -> %d\n", dName, dScore)
}
