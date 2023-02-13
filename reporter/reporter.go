package reporter

import (
	"fmt"
	"io"

	"github.com/corani/silver-octo-sniffle/token"
)

type Reporter struct {
	out   io.Writer
	debug bool
}

func NewReporter(out io.Writer, debug bool) *Reporter {
	return &Reporter{
		out:   out,
		debug: debug,
	}
}

func (r *Reporter) Debugf(t token.Token, format string, args ...any) {
	if r.debug {
		str := fmt.Sprintf(format, args...)
		fmt.Fprintf(r.out, "%s:%d:%d: DEBUG: %s\n", t.File, t.Range.FromRow, t.Range.FromCol, str)
	}
}

func (r *Reporter) Infof(t token.Token, format string, args ...any) {
	str := fmt.Sprintf(format, args...)
	fmt.Fprintf(r.out, "%s:%d:%d: INFO : %s\n", t.File, t.Range.FromRow, t.Range.FromCol, str)
}

func (r *Reporter) Warnf(t token.Token, format string, args ...any) {
	str := fmt.Sprintf(format, args...)
	fmt.Fprintf(r.out, "%s:%d:%d: WARN : %s\n", t.File, t.Range.FromRow, t.Range.FromCol, str)
}

func (r *Reporter) Errorf(t token.Token, format string, args ...any) {
	str := fmt.Sprintf(format, args...)
	fmt.Fprintf(r.out, "%s:%d:%d: ERROR: %s\n", t.File, t.Range.FromRow, t.Range.FromCol, str)
}
