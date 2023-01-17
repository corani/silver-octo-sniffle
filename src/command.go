package main

import (
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

func execute(name string) error {
	cmd := exec.Command(name)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
