package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/corani/silver-octo-sniffle/ast"
	"github.com/corani/silver-octo-sniffle/checker"
	"github.com/corani/silver-octo-sniffle/generator"
	"github.com/corani/silver-octo-sniffle/lexer"
	"github.com/corani/silver-octo-sniffle/parser"
	"github.com/corani/silver-octo-sniffle/reporter"
	"github.com/corani/silver-octo-sniffle/token"
)

type CompilationResult struct {
	tokens token.Tokens
	ast    ast.Node
	ir     string
	path   string

	lexerOut string
}

func do(srcName string, bs []byte) (CompilationResult, error) {
	var buf bytes.Buffer

	tokens, err := lexer.Lex(reporter.NewReporter(&buf), srcName, bs)
	if err != nil {
		return CompilationResult{
			lexerOut: buf.String(),
		}, err
	}

	ast, err := parser.Parse(tokens)
	if err != nil {
		return CompilationResult{
			tokens: tokens,
		}, err
	}

	if err := checker.TypeCheck(ast); err != nil {
		return CompilationResult{
			tokens: tokens,
			ast:    ast,
		}, err
	}

	baseName := strings.TrimSuffix(filepath.Base(srcName), filepath.Ext(srcName))

	dir, err := os.MkdirTemp("", "")
	if err != nil {
		return CompilationResult{
			tokens: tokens,
			ast:    ast,
		}, err
	}
	defer os.RemoveAll(dir)

	llName := filepath.Join(dir, baseName+".ll")
	outName := filepath.Join(filepath.Dir(filepath.Dir(srcName)), "bin", baseName)

	out, err := os.Create(llName)
	if err != nil {
		return CompilationResult{
			tokens: tokens,
			ast:    ast,
		}, err
	}
	defer out.Close()

	var ir bytes.Buffer

	if err := generator.GenerateIR(io.MultiWriter(&ir, out), ast); err != nil {
		return CompilationResult{
			tokens: tokens,
			ast:    ast,
		}, err
	}

	if err := compile(llName, outName); err != nil {
		return CompilationResult{
			tokens: tokens,
			ast:    ast,
			ir:     ir.String(),
		}, err
	}

	return CompilationResult{
		tokens: tokens,
		ast:    ast,
		ir:     ir.String(),
		path:   outName,
	}, nil
}

func main() {
	srcName := flag.String("src", "", "source file")
	flag.Parse()

	bs, err := os.ReadFile(*srcName)
	if err != nil {
		panic(err)
	}

	if result, err := do(*srcName, bs); err != nil {
		if errors.Is(err, lexer.ErrLexer) {
			fmt.Fprintln(os.Stderr, result.lexerOut)
		}

		panic(err)
	}
}
