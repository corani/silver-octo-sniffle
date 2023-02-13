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

	lexerOut     string
	parserOut    string
	checkerOut   string
	generatorOut string
}

func do(srcName string, bs []byte, debug bool) (CompilationResult, error) {
	var buf bytes.Buffer

	tokens, err := lexer.Lex(reporter.NewReporter(&buf, debug), srcName, bs)
	if err != nil {
		return CompilationResult{
			lexerOut: buf.String(),
		}, err
	}

	buf.Reset()

	ast, err := parser.Parse(reporter.NewReporter(&buf, debug), tokens)
	if err != nil || buf.Len() > 0 {
		return CompilationResult{
			tokens:    tokens,
			parserOut: buf.String(),
		}, parser.ErrParsing
	}

	buf.Reset()

	checker.TypeCheck(reporter.NewReporter(&buf, debug), ast)
	if buf.Len() > 0 {
		return CompilationResult{
			tokens:     tokens,
			ast:        ast,
			checkerOut: buf.String(),
		}, checker.ErrChecking
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

	buf.Reset()

	generator.GenerateIR(reporter.NewReporter(&buf, debug), io.MultiWriter(&ir, out), ast)
	if buf.Len() > 0 {
		return CompilationResult{
			tokens:       tokens,
			ast:          ast,
			generatorOut: buf.String(),
		}, generator.ErrGenerating
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

	if result, err := do(*srcName, bs, false); err != nil {
		if errors.Is(err, lexer.ErrLexer) {
			fmt.Fprintln(os.Stderr, result.lexerOut)
		}

		if errors.Is(err, parser.ErrParsing) {
			fmt.Fprintln(os.Stderr, result.parserOut)
		}

		if errors.Is(err, checker.ErrChecking) {
			fmt.Fprintln(os.Stderr, result.checkerOut)
		}

		if errors.Is(err, generator.ErrGenerating) {
			fmt.Fprintln(os.Stderr, result.generatorOut)
		}

		panic(err)
	}
}
