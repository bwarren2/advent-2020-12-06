package advent20201206

import (
	"bufio"
	"os"
	"strings"
	// mapset "github.com/deckarep/golang-set"
)

// RecordsFromFile returns a channel that gives records from a file
func RecordsFromFile(filename string) <-chan string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	iterator := make(chan string)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	go func(scanner *bufio.Scanner) {
		buildString := new(strings.Builder)
		for scanner.Scan() {
			line := scanner.Text()
			buildString.WriteString(line)
			buildString.WriteString(" ")
			if line == "" {
				iterator <- buildString.String()
				buildString = new(strings.Builder)
			}
		}
		iterator <- buildString.String()
		close(iterator)
	}(scanner)
	return iterator
}

// GroupQuestionMap gets a map of answers
func GroupQuestionMap(filename string) (map[int64]map[rune]int64, map[int64]int64) {
	iterator := RecordsFromFile(filename)
	answerCounts := make(map[int64]map[rune]int64)
	groupTotal := make(map[int64]int64)
	var groupID int64
	for value := range iterator {
		// fmt.Println("Got: ", value)
		answerCounts[groupID] = make(map[rune]int64)
		for _, rune := range value {
			if rune != ' ' {
				answerCounts[groupID][rune]++
			}
		}
		groupTotal[groupID] = int64(len(strings.Fields(value)))
		// fmt.Println(groupTotal[groupID], strings.Split(strings.Trim(value, " "), " "))
		groupID++
	}
	return answerCounts, groupTotal
}

// Part1 answers part 1
func Part1(filename string) (count int) {
	groupAnswerMap, _ := GroupQuestionMap(filename)
	for _, answerMap := range groupAnswerMap {
		count += len(answerMap)
	}
	return
}

// Part2 answers part 2
func Part2(filename string) (count int) {
	groupAnswerMap, groupTotal := GroupQuestionMap(filename)
	// spew.Dump(groupAnswerMap, groupTotal)
	for groupID, answerMap := range groupAnswerMap {
		for _, ct := range answerMap {
			if ct == groupTotal[groupID] {
				count++
			}
		}
	}
	return
}
