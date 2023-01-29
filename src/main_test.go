package main

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMain(t *testing.T) {
	t.Parallel()

	// NOTE(daniel): as we're writing the source filename to the expected output,
	// we need to change directory rather than doing a walk starting at `../test`.
	if err := os.Chdir(".."); err != nil {
		t.Fatal(err)
	}

	filepath.Walk("test", func(path string, info fs.FileInfo, err error) error {
		if err != nil || !strings.HasSuffix(info.Name(), ".in") {
			return err
		}

		t.Run(info.Name(), func(t *testing.T) {
			t.Parallel()

			// TODO(daniel): this could be more robust.
			exp, err := os.ReadFile(strings.ReplaceAll(path, ".in", ".exp"))
			if err != nil {
				t.Fatal(err)
			}

			var out bytes.Buffer

			do(path, &out, true, true, true)

			if diff := cmp.Diff(string(exp), out.String()); diff != "" {
				t.Error(diff)
			}
		})

		return nil
	})
}
