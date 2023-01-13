package main

import (
	"flag"
	"fmt"
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

	baseName := strings.TrimSuffix(*srcName, filepath.Ext(*srcName))
	llName := baseName + ".ll"
	outName := baseName + ".out"

	out, err := os.Create(llName)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	generateIR(out, ast)

	if err := compile(llName, outName); err != nil {
		panic(err)
	}

	if *run {
		absName, err := filepath.Abs(outName)
		if err != nil {
			panic(err)
		}

		if err := execute(absName); err != nil {
			panic(err)
		}
	}
}
