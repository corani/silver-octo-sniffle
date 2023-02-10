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

	"github.com/corani/silver-octo-sniffle/ast"
	"github.com/google/go-cmp/cmp"
	"github.com/yuin/goldmark"
	mdast "github.com/yuin/goldmark/ast"
	mdtext "github.com/yuin/goldmark/text"
)

func getSpan(lines *mdtext.Segments) (int, int) {
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

func findSource(src []byte, node mdast.Node) []byte {
	if node.Kind() == mdast.KindFencedCodeBlock {
		code := node.(*mdast.FencedCodeBlock)
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
	root := parser.Parse(mdtext.NewReader(exp))

	source := findSource(exp, root)
	if source == nil {
		t.Fatal("couldn't find source")
	}

	return exp, source
}

func doTest(t *testing.T, srcName string, w io.Writer, bs []byte) {
	result, _ := do(srcName, bs)

	fmt.Fprintf(w, "# %s\n", srcName)

	fmt.Fprintf(w, "## Source\n")
	fmt.Fprintln(w, "```pascal")
	fmt.Fprint(w, string(bs))
	fmt.Fprintln(w, "```")

	if result.lexerOut != "" {
		fmt.Fprintln(w, "## Lexer errors")
		fmt.Fprintf(w, "```\n%s```\n", result.lexerOut)

		return
	}

	fmt.Fprintln(w, result.tokens)

	if result.parserOut != "" {
		fmt.Fprintln(w, "## Parser errors")
		fmt.Fprintf(w, "```\n%s```\n", result.parserOut)

		return
	}

	ast.PrintAST(w, result.ast)

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
		t.Log(result.ir)
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

	if info, err := os.Stat("bin"); err != nil {
		if err := os.Mkdir("bin", 0775); err != nil {
			t.Fatal(err)
		}
	} else if !info.IsDir() {
		t.Fatal("`bin` exists but is not a directory")
	}

	_ = filepath.Walk("test", func(path string, info fs.FileInfo, err error) error {
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
