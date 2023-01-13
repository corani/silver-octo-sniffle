package main

import (
	"fmt"
	"os"
	"os/exec"
)

func compile(source, target string) error {
	fmt.Printf("[CMD] clang -o %v %v\n", target, source)

	cmd := exec.Command("clang", "-o", target, source)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func execute(name string) error {
	fmt.Printf("[CMD] %v\n", name)

	cmd := exec.Command(name)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
