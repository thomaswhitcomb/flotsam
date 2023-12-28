package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/thomaswhitcomb/jetsam"
)

type ageName struct {
	fname string
	lname string
}

type ageNameMap map[int][]ageName

func isEven(n int) bool {
	return n%2 == 0
}
func makeAgeArray(ages ageNameMap) []int {
	agesonly := []int{}
	for k, v := range ages {
		for i := 0; i < len(v); i++ {
			agesonly = append(agesonly, k)
		}
	}
	return agesonly
}

func calculateMedian(ages ageNameMap) (float64, []ageName) {
	// produce an integer array with an age entry for every person
	agesonly := makeAgeArray(ages)
	// sort the integer array.
	sort.Ints(agesonly)

	n := len(agesonly)
	//If the integer array has an even number
	// than compute the median from the two middle
	// integers
	if isEven(n) {
		mid := int(math.Floor(float64(n)/2) - 1)
		median := float64(agesonly[mid]+agesonly[mid+1]) / 2
		names := ages[agesonly[mid]]
		if agesonly[mid] != agesonly[mid+1] {
			for _, name := range ages[agesonly[mid+1]] {
				names = append(names, name)
			}
		}
		return median, names
	} else {
		mid := int(math.Floor(float64(n) / 2.0))
		return float64(agesonly[mid]), ages[agesonly[mid]]
	}
}
func parseLine(line string) (int, ageName) {
	parts := strings.Split(line, ",")
	if len(parts) != 3 {
		log.Printf("Less than 3: %s\n", line)
		return -1, ageName{}
	}
	age, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil {
		log.Printf(
			"Age conversion failed. |%s| %v\n",
			strings.TrimSpace(parts[2]), err)
		return -1, ageName{}
	}
	an := ageName{
		fname: strings.TrimSpace(parts[0]),
		lname: strings.TrimSpace(parts[1]),
	}
	return age, an
}
func reduce(in <-chan string, done chan<- jetsam.DoneChanMsg) {
	n := 0
	tot := 0
	ages := map[int][]ageName{}
	for {
		line, ok := <-in
		if ok {
			age, an := parseLine(line)
			if age == -1 {
				continue
			}
			if val, ok := ages[age]; ok {
				ages[age] = append(val, an)
			} else {
				ages[age] = []ageName{an}
			}
			tot = tot + age

			n++
		} else {
			// Queue empty, stop
			break
		}
	}
	if len(ages) == 0 {
		done <- jetsam.DoneChanMsg{
			Results:         jetsam.DoneParms{"status": "no content found"},
			NumberProcessed: 0,
		}
		return
	}
	age, who := calculateMedian(ages)
	names := who[0].fname + " " + who[0].lname
	for i := 1; i < len(who); i++ {
		names = names + ", " + who[i].fname + " " + who[i].lname
	}
	done <- jetsam.DoneChanMsg{
		Results: jetsam.DoneParms{
			"median":  fmt.Sprintf("%f", age),
			"average": fmt.Sprintf("%f", float64(tot)/float64(n)),
			"names":   names,
			"count":   fmt.Sprintf("%d", len(who)),
		},
		NumberProcessed: n,
	}
}
