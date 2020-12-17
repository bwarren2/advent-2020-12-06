package advent20201206_test

import (
	"fmt"
	"testing"

	advent "github.com/bwarren2/advent20201206"
)

func TestRecordsFromFile(t *testing.T) {
	advent.RecordsFromFile("sample.txt")
}

func TestPart1(t *testing.T) {
	fmt.Println(advent.Part1("input.txt"))
}

func TestPart2(t *testing.T) {
	fmt.Println(advent.Part2("input.txt"))
}
