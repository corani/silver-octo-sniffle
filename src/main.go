package main

import (
	"bytes"
	"flag"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/corani/silver-octo-sniffle/ast"
	"github.com/corani/silver-octo-sniffle/check"
	"github.com/corani/silver-octo-sniffle/generate"
	"github.com/corani/silver-octo-sniffle/lex"
	"github.com/corani/silver-octo-sniffle/parse"
)

type CompilationResult struct {
	tokens lex.Tokens
	ast    ast.Node
	ir     string
	path   string
}

func do(srcName string, bs []byte) (CompilationResult, error) {
	tokens, err := lex.Tokenize(srcName, bs)
	if err != nil {
		return CompilationResult{}, err
	}

	ast, err := parse.Parse(tokens)
	if err != nil {
		return CompilationResult{
			tokens: tokens,
		}, err
	}

	if err := check.TypeCheck(ast); err != nil {
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

	if err := generate.GenerateIR(io.MultiWriter(&ir, out), ast); err != nil {
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

	if _, err := do(*srcName, bs); err != nil {
		panic(err)
	}
}
