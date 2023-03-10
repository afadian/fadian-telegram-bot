package util

import (
	"os"
	"path/filepath"
)

func AbsolutePath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}

	exe, _ := os.Executable()
	dir := filepath.Dir(exe)
	return filepath.Join(dir, path)
}
