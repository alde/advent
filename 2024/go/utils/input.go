package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func SplitInt(input string) ([]int, error) {
	var result []int
	strs := strings.Split(input, ",")
	for _, s := range strs {
		val, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		result = append(result, val)
	}
	return result, nil
}

type LineData struct {
	Line  string
	Error error
}

func LazyReadLines(filename string) <-chan *LineData {
	ch := make(chan *LineData)

	go func() {
		defer close(ch)

		// Open the file
		file, err := os.Open(filename)
		if err != nil {
			// Send the error to the channel and close it
			ch <- &LineData{"", err}
			return
		}
		defer file.Close()

		// Create a scanner to read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ch <- &LineData{scanner.Text(), nil}
		}

		// Check for any error that occurred while scanning
		if err := scanner.Err(); err != nil {
			ch <- &LineData{"", err}
		}
	}()

	return ch
}

func ConsumeAllInput(linesChan <-chan *LineData) ([]string, error) {
	input := []string{}
	for line := range linesChan {
		if line.Error != nil {
			return nil, line.Error
		}
		input = append(input, line.Line)
	}
	return input, nil
}
