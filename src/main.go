package main

import (
	"bytes"
	"flag"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type CompilationResult struct {
	tokens Tokens
	ast    Node
	ir     string
	path   string
}

func do(srcName string, bs []byte) (CompilationResult, error) {
	tokens, err := lex(srcName, bs)
	if err != nil {
		return CompilationResult{}, err
	}

	ast, err := parse(tokens)
	if err != nil {
		return CompilationResult{
			tokens: tokens,
		}, err
	}

	if err := typeCheck(ast); err != nil {
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

	if err := generateIR(io.MultiWriter(&ir, out), ast); err != nil {
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
