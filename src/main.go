package main

import (
	"bytes"
	"flag"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func do(srcName string, bs []byte) (Tokens, Node, string, string) {
	tokens, err := lex(srcName, bs)
	if err != nil {
		panic(err)
	}

	ast, err := parse(tokens)
	if err != nil {
		panic(err)
	}

	typeCheck(ast)

	baseName := strings.TrimSuffix(filepath.Base(srcName), filepath.Ext(srcName))

	dir, err := os.MkdirTemp("", "")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	llName := filepath.Join(dir, baseName+".ll")
	outName := filepath.Join(filepath.Dir(filepath.Dir(srcName)), "bin", baseName)

	out, err := os.Create(llName)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	var ir bytes.Buffer

	generateIR(io.MultiWriter(&ir, out), ast)

	if err := compile(llName, outName); err != nil {
		panic(err)
	}

	return tokens, ast, ir.String(), outName
}

func main() {
	srcName := flag.String("src", "", "source file")
	flag.Parse()

	bs, err := os.ReadFile(*srcName)
	if err != nil {
		panic(err)
	}

	do(*srcName, bs)
}
