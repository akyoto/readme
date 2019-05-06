package main

import (
	"os"
	"path"
	"path/filepath"

	"github.com/akyoto/color"
	"github.com/akyoto/readme/generator"
)

func main() {
	// Traverse directory
	err := filepath.Walk(".", func(fullPath string, info os.FileInfo, err error) error {
		if err != nil {
			color.Red(err.Error())
			return nil
		}

		if info == nil {
			color.Red("Invalid file path: %s", fullPath)
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if info.Name() == "README.src.md" {
			directory := filepath.Dir(fullPath)
			_ = generator.Run(fullPath, path.Join(directory, "README.md"))
		}

		return nil
	})

	if err != nil {
		color.Red(err.Error())
	}
}
