package shared

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}

type LineGenerator chan string

func ReadLines(path string) LineGenerator {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	lines := make(LineGenerator)

	go func() {
		defer close(lines)
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading file:", err)
		}
		file.Close()
	}()

	return lines
}
