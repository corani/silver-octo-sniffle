package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func do(srcName string, w io.Writer, outTokens bool, outAST bool, run bool) {
	bs, err := os.ReadFile(srcName)
	if err != nil {
		panic(err)
	}

	tokens, err := lex(srcName, bs)
	if err != nil {
		panic(err)
	}

	if outTokens || outAST || run {
		fmt.Fprintf(w, "# %s\n", srcName)
		fmt.Fprintln(w, "```")
		fmt.Fprint(w, string(bs))
		fmt.Fprintln(w, "```")
	}

	if outTokens {
		fmt.Fprintln(w, tokens)
	}

	ast, err := parse(tokens)
	if err != nil {
		panic(err)
	}

	typeCheck(ast)

	if outAST {
		printAST(w, ast)
	}

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

	fmt.Fprintln(w, "## IR")
	fmt.Fprintln(w, "```llvm")
	generateIR(io.MultiWriter(out, w), ast)
	fmt.Fprintln(w, "```")

	if err := compile(llName, outName); err != nil {
		panic(err)
	}

	if run {
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
}

func main() {
	srcName := flag.String("src", "", "source file")
	outTokens := flag.Bool("tokens", false, "print tokens")
	outAST := flag.Bool("ast", false, "print AST")
	run := flag.Bool("run", false, "run output")
	flag.Parse()

	do(*srcName, os.Stdout, *outTokens, *outAST, *run)
}
