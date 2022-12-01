package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
    "bufio"
	"strconv"
	"sort"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")


	defer input.Close()
	scanner := bufio.NewScanner(input)

	at := 0;
	elfMap := map[int]int{}
	elfMap[at] = 0
	for scanner.Scan() {
		valString := scanner.Text()

		if valString == "" {
			at++
			elfMap[at] = 0
			continue
		} else {
			valInt,_ := strconv.Atoi(valString)
			elfMap[at] += valInt
		}
	}

	sums := []int{}
	for _,val := range elfMap {
		sums = append(sums, val)
	}

	sort.Ints(sums)

	fmt.Printf("Part 1: %d\n", sums[len(sums) - 1])
	fmt.Printf("Part 2: %d\n", sums[len(sums) - 1] + sums[len(sums) - 2] + sums[len(sums) - 3])
}
