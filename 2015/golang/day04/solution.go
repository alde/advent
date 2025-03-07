package day04

import (
	"crypto/md5"
	"fmt"
	"strings"
	"time"

	"github.com/alde/advent/2015/golang/shared"
)

func Solve() {
	result := shared.Result{
		Title: "The Ideal Stocking Stuffer",
		Day:   4,
	}

	result.Parts = []shared.Part{
		part1(),
		part2(),
	}

	shared.PrettyPrint(result)
}

func part1() shared.Part {
	start := time.Now()
	result := firstNumber("ckczppom", 5)

	return shared.Part{
		Duration: time.Since(start),
		Result:   result,
	}
}

func part2() shared.Part {
	start := time.Now()
	result := firstNumber("ckczppom", 6)

	return shared.Part{
		Duration: time.Since(start),
		Result:   result,
	}
}

func firstNumber(secretKey string, targetZeros int) int {
	padding := 10_000
	prefix := strings.Repeat("0", targetZeros)
	for {
		candidate := fmt.Sprintf("%s%d", secretKey, padding)

		hash := md5.Sum([]byte(candidate))
		hex := fmt.Sprintf("%x", hash)
		if strings.HasPrefix(hex, prefix) {
			return padding
		}
		padding++
	}
}
