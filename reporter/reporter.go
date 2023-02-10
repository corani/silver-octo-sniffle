package reporter

import (
	"fmt"
	"io"

	"github.com/corani/silver-octo-sniffle/token"
)

type Reporter struct {
	out io.Writer
}

func NewReporter(out io.Writer) *Reporter {
	return &Reporter{
		out: out,
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
