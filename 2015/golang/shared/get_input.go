package shared

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func InputMustExist(path string, year int, day int) {
	if _, err := os.Stat(path); err == nil {
		return
	}
	slog.Info("input not found - attempting to download", "year", year, "day", day)
	data := getInput(year, day)
	writeFile(path, data)
	if _, err := os.Stat(path); err != nil {
		slog.Error("failed to download input file")
	}
}

func getInput(year int, day int) []byte {
	sessionCookie, ok := os.LookupEnv("ADVENT_SESSION")
	if !ok {
		fmt.Print("ADVENT_SESSION not set - please input session cookie: ")
		fmt.Scan(&sessionCookie)
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	resp, err := client.Do(req)
	if err != nil {
		panic("failed getting input data")
	}

	body, _ := io.ReadAll(resp.Body)

	return body
}

func writeFile(path string, data []byte) {
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		panic(err)
	}

}
