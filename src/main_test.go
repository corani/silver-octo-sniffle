package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

func getSpan(lines *text.Segments) (int, int) {
	start := 0
	end := 0

	for i := 0; i < lines.Len(); i++ {
		line := lines.At(i)

		if start == 0 || start > line.Start {
			start = line.Start
		}

		if end == 0 || end < line.Stop {
			end = line.Stop
		}
	}

	return start, end
}

func findSource(src []byte, node ast.Node) []byte {
	if node.Kind() == ast.KindFencedCodeBlock {
		code := node.(*ast.FencedCodeBlock)
		start, end := getSpan(code.Lines())

		return src[start:end]
	}

	if node.ChildCount() == 0 {
		return nil
	}

	for child := node.FirstChild(); child.NextSibling() != nil; child = child.NextSibling() {
		if v := findSource(src, child); v != nil {
			return v
		}
	}

	return nil
}

func readGoldenTest(t *testing.T, path string) ([]byte, []byte) {
	t.Helper()

	exp, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	parser := goldmark.DefaultParser()
	root := parser.Parse(text.NewReader(exp))

	source := findSource(exp, root)
	if source == nil {
		t.Fatal("couldn't find source")
	}

	return exp, source
}

func doTest(t *testing.T, srcName string, w io.Writer, bs []byte) {
	t.Helper()

	result, err := do(srcName, bs)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Fprintf(w, "# %s\n", srcName)

	fmt.Fprintf(w, "## Source\n")
	fmt.Fprintln(w, "```pascal")
	fmt.Fprint(w, string(bs))
	fmt.Fprintln(w, "```")

	fmt.Fprintln(w, result.tokens)

	printAST(w, result.ast)

	fmt.Fprintln(w, "## IR")
	fmt.Fprintln(w, "```llvm")
	fmt.Fprint(w, result.ir)
	fmt.Fprintln(w, "```")

	fmt.Fprintln(w, "## Run")
	fmt.Fprintln(w, "```bash")

	absName, err := filepath.Abs(result.path)
	if err != nil {
		t.Fatal(err)
	}

	if err := execute(w, absName); err != nil {
		t.Fatal(err)
	}

	fmt.Fprintln(w, "```")
}

var update = flag.Bool("update", false, "update test cases")

func TestMain(t *testing.T) {
	t.Parallel()

	// NOTE(daniel): as we're writing the source filename to the expected output,
	// we need to change directory rather than doing a walk starting at `../test`.
	if err := os.Chdir(".."); err != nil {
		t.Fatal(err)
	}

	// TODO(daniel): perhaps we should have only a single markdown file per
	// test and extract the source code from the first code block.
	filepath.Walk("test", func(path string, info fs.FileInfo, err error) error {
		if err != nil || !strings.HasSuffix(info.Name(), ".md") {
			return err
		}

		t.Run(info.Name(), func(t *testing.T) {
			t.Parallel()

			exp, source := readGoldenTest(t, path)

			var out bytes.Buffer

			doTest(t, path, &out, source)

			if *update {
				if err := os.WriteFile(path, out.Bytes(), 0664); err != nil {
					t.Error(err)
				}
			} else {
				if diff := cmp.Diff(string(exp), out.String()); diff != "" {
					t.Error(diff)
				}
			}
		})

		return nil
	})
}
