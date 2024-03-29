package generator

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Run reads the template file and writes the generated file to the output path.
func Run(templatePath string, outputPath string) error {
	absTemplatePath, err := filepath.Abs(templatePath)

	if err != nil {
		return err
	}

	directory := filepath.Dir(absTemplatePath)
	name := filepath.Base(directory)

	// Find remote URL and relative path
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	cmd.Dir = directory
	output, err := cmd.Output()

	if err != nil {
		return err
	}

	remoteOriginURL := strings.TrimSpace(string(output))
	relativePath := strings.TrimSuffix(remoteOriginURL, ".git")
	relativePath = strings.TrimPrefix(relativePath, "git@")
	relativePath = strings.TrimPrefix(relativePath, "https://")
	relativePath = strings.TrimPrefix(relativePath, "github.com:")
	relativePath = strings.TrimPrefix(relativePath, "github.com/")

	// Header
	header := `
		[![Godoc][godoc-image]][godoc-url]
		[![Report][report-image]][report-url]
		[![Tests][tests-image]][tests-url]
		[![Coverage][coverage-image]][coverage-url]
		[![Sponsor][sponsor-image]][sponsor-url]
	`
	header = strings.TrimSpace(header)
	header = strings.ReplaceAll(header, "\t", "")

	sponsors := []Person{
		// {
		// 	Name:        "Max Räche",
		// 	ImageSource: "https://avatars3.githubusercontent.com/u/39157397?s=70&v=4",
		// 	Link:        "https://github.com/yukinamida",
		// },
		// {
		// 	Name:        "Cedric Fung",
		// 	ImageSource: "https://avatars3.githubusercontent.com/u/2269238?s=70&v=4",
		// 	Link:        "https://github.com/cedricfung",
		// },
		// {
		// 	Name:        "Scott Rayapoullé",
		// 	ImageSource: "https://avatars3.githubusercontent.com/u/11772084?s=70&v=4",
		// 	Link:        "https://github.com/soulcramer",
		// },
		// {
		// 	Name:        "Eduard Urbach",
		// 	ImageSource: "https://avatars3.githubusercontent.com/u/438936?s=70&v=4",
		// 	Link:        "https://eduardurbach.com",
		// },
	}

	tableHeader := strings.Builder{}
	tableSeparator := strings.Builder{}
	tableFooter := strings.Builder{}

	for index, sponsor := range sponsors {
		fmt.Fprintf(&tableHeader, "[![%s](%s)](%s)", sponsor.Name, sponsor.ImageSource, sponsor.Link)
		tableSeparator.WriteString("---")
		fmt.Fprintf(&tableFooter, "[%s](%s)", sponsor.Name, sponsor.Link)

		if index != len(sponsors)-1 {
			tableHeader.WriteString(" | ")
			tableSeparator.WriteString(" | ")
			tableFooter.WriteString(" | ")
		}
	}

	// Footer
	footer := fmt.Sprintf(`
		[godoc-image]: https://godoc.org/github.com/{relativePath}?status.svg
		[godoc-url]: https://godoc.org/github.com/{relativePath}
		[report-image]: https://goreportcard.com/badge/github.com/{relativePath}
		[report-url]: https://goreportcard.com/report/github.com/{relativePath}
		[tests-image]: https://cloud.drone.io/api/badges/{relativePath}/status.svg
		[tests-url]: https://cloud.drone.io/{relativePath}
		[coverage-image]: https://codecov.io/gh/{relativePath}/graph/badge.svg
		[coverage-url]: https://codecov.io/gh/{relativePath}
		[sponsor-image]: https://img.shields.io/badge/github-donate-green.svg
		[sponsor-url]: https://github.com/users/akyoto/sponsorship
	`)

	footer = strings.TrimSpace(footer)
	footer = strings.ReplaceAll(footer, "\t", "")
	footer = strings.ReplaceAll(footer, "{relativePath}", relativePath)

	// Install
	install := fmt.Sprintf("## Installation\n\n```shell\ngo get -u github.com/%s/...\n```", relativePath)

	snippets := map[string]string{
		"name":       name,
		"go:header":  header,
		"go:footer":  footer,
		"go:install": install,
	}

	inFile, err := os.Open(templatePath)

	if err != nil {
		return err
	}

	defer inFile.Close()

	outFile, err := os.Create(outputPath)

	if err != nil {
		return err
	}

	defer outFile.Close()
	bufferedOutFile := bufio.NewWriter(outFile)
	defer bufferedOutFile.Flush()
	buffer := make([]byte, 4096)
	inSnippetName := false
	snippetName := strings.Builder{}

	for {
		n, err := inFile.Read(buffer)

		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			return nil
		}

		for i := 0; i < n; i++ {
			char := buffer[i]

			if char == '{' {
				inSnippetName = true
				snippetName.Reset()
				continue
			}

			if char == '}' {
				inSnippetName = false
				code := snippets[snippetName.String()]
				_, _ = bufferedOutFile.WriteString(code)
				continue
			}

			if inSnippetName {
				snippetName.WriteByte(char)
				continue
			}

			_ = bufferedOutFile.WriteByte(char)
		}
	}
}
