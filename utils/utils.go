package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/k0kubun/pp/v3"
)

func ReadFromText(Path string) ([]int, []int, error) {
	file, err := os.Open(Path)
	pp.Println("Reading from %v", Path)
	if err != nil {
		pp.Println("Printing")
		return nil, nil, fmt.Errorf("cannot open file: %v", file)
	}

	defer file.Close()

	var nums1, nums2 []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// pp.Println(scanner.Text())
		var lineNums []int

		for _, numStr := range strings.Fields(scanner.Text()) {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, nil, fmt.Errorf("invalid number: %v", numStr)
			}

			lineNums = append(lineNums, num)
		}

		if len(lineNums) != 2 {
			return nil, nil, fmt.Errorf("expected 2 numbers in line, got %d", len(lineNums))
		}

		nums1 = append(nums1, lineNums[0])
		nums2 = append(nums2, lineNums[1])

	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %v", err)
	}

	return nums1, nums2, nil
}

func Distance(arr []int, arr1 []int) (int, error) {
	var distance int64

	for i := range arr {
		diff := arr[i] - arr1[i]
		pp.Printf("\n diff(%d, %d) = %d", arr[i], arr1[i], diff)
		if diff < 0 {
			diff = -diff
		}

		distance += int64(diff)
		pp.Println("Total Diff: %d", distance)
	}

	return int(distance), nil
}
