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
	"github.com/russross/blackfriday/v2"
)

func readGoldenTest(t *testing.T, path string) ([]byte, []byte) {
	t.Helper()

	exp, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	parser := blackfriday.New()

	var source string

	root := parser.Parse(exp)
	root.Walk(func(node *blackfriday.Node, _ bool) blackfriday.WalkStatus {
		switch node.Type {
		case blackfriday.Heading:
			if node.HeadingID == "Source" {
				return blackfriday.GoToNext
			}

			return blackfriday.SkipChildren
		case blackfriday.Document, blackfriday.Paragraph:
			return blackfriday.GoToNext
		case blackfriday.Code:
			source = string(node.Literal)
			source = strings.TrimPrefix(source, "\n")
			source = strings.TrimSuffix(source, "\n")
			source = strings.TrimPrefix(source, "```")
			source = strings.TrimSuffix(source, "```")
			source = source + "\n"

			return blackfriday.Terminate
		default:
			return blackfriday.SkipChildren
		}
	})

	return exp, []byte(source)
}

func doTest(t *testing.T, srcName string, w io.Writer, bs []byte) {
	t.Helper()

	tokens, ast, ir, outName := do(srcName, bs)

	fmt.Fprintf(w, "# %s\n", srcName)

	fmt.Fprintf(w, "## Source\n")
	fmt.Fprintln(w, "```")
	fmt.Fprint(w, string(bs))
	fmt.Fprintln(w, "```")

	fmt.Fprintln(w, tokens)

	printAST(w, ast)

	fmt.Fprintln(w, "## IR")
	fmt.Fprintln(w, "```llvm")
	fmt.Fprint(w, ir)
	fmt.Fprintln(w, "```")

	fmt.Fprintln(w, "## Run")
	fmt.Fprintln(w, "```bash")

	absName, err := filepath.Abs(outName)
	if err != nil {
		panic(err)
	}

	if err := execute(w, absName); err != nil {
		panic(err)
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
