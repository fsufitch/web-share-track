package webtrack

import (
	"path/filepath"
)

var resource_dir string

func SetResourceDir(newdir string) {
	resource_dir = newdir
}

func GetResource(path ...string) string {
	full_path := append([]string{resource_dir}, path...)
	return filepath.Join(full_path...)
}
