package main

import (
	"io"
	"os"
	"os/exec"
)

func compile(source, target string) error {
	cmd := exec.Command("clang",
		"-o", target, source,
		"-Wno-override-module",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func execute(w io.Writer, name string) error {
	cmd := exec.Command(name)

	cmd.Stdout = w
	cmd.Stderr = w

	return cmd.Run()
}
