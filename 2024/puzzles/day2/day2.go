package day2

import (
	_ "embed"
	"strings"

	"github.com/KhachikAstoyan/advent-of-code/2024/utils"
)

type Day2 struct{}

//go:embed input.txt
var input string

func (d Day2) Name() string {
	return "Day 2: Red-Nosed Reports"
}

func (d Day2) Part1() uint64 {
	matrix := parseMatrix(input)
	safeReports := 0

	for _, report := range matrix {
		isReportDecreasing := report[0] > report[1]
		isSafe := true

		// we're going to check report[0] and report[1] twice
		// but it's not a big deal
		for i := 1; i < len(report); i++ {
			absDiff := utils.AbsDiff(report[i-1], report[i])
			isDecr := report[i-1] > report[i]

			if isReportDecreasing != isDecr || absDiff < 1 || absDiff > 3 {
				isSafe = false
				break
			}
		}

		if isSafe {
			safeReports++
		}
	}

	return uint64(safeReports)
}

func (d Day2) Part2() uint64 {
	return 0
}

func parseMatrix(input string) [][]uint64 {
	lines := strings.Split(input, "\n")
	matrix := make([][]uint64, len(lines))

	for i, line := range lines {
		parts := strings.Fields(line)
		parsedLine := make([]uint64, len(parts))

		for j, part := range parts {
			parsedLine[j] = uint64(utils.ParseInt(part, 10, 64))
		}

		matrix[i] = parsedLine
	}

	return matrix
}
