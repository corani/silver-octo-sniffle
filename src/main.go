package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	srcName := flag.String("src", "", "source file")
	outTokens := flag.Bool("tokens", false, "print tokens")
	outAST := flag.Bool("ast", false, "print AST")
	run := flag.Bool("run", false, "run output")
	flag.Parse()

	bs, err := os.ReadFile(*srcName)
	if err != nil {
		panic(err)
	}

	tokens, err := lex(*srcName, bs)
	if err != nil {
		panic(err)
	}

	if *outTokens {
		fmt.Println(tokens)
	}

	ast, err := parse(tokens)
	if err != nil {
		panic(err)
	}

	if *outAST {
		printAST(ast)
	}

	baseName := strings.TrimSuffix(filepath.Base(*srcName), filepath.Ext(*srcName))

	dir, err := os.MkdirTemp("", "")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	llName := filepath.Join(dir, baseName+".ll")
	outName := filepath.Join(filepath.Dir(filepath.Dir(*srcName)), "bin", baseName)

	out, err := os.Create(llName)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	fmt.Println("IR:")
	generateIR(io.MultiWriter(out, os.Stdout), ast)

	if err := compile(llName, outName); err != nil {
		panic(err)
	}

	if *run {
		fmt.Println("Run:")
		absName, err := filepath.Abs(outName)
		if err != nil {
			panic(err)
		}

		if err := execute(absName); err != nil {
			panic(err)
		}
	}
}
