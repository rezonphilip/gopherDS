package recursion

import (
	"fmt"
	"os"
	"path/filepath"
)

func Walk(path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("errored while reading dir %s\n", path)
	}

	for _, entr := range entries {
		full := filepath.Join(path, entr.Name())
		if entr.IsDir() {
			Walk(full)
		} else {
			fmt.Printf("dir - %s - has file -%s\n", path, entr.Name())
		}
	}
}
