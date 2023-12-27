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

func findMedium(ages ageNameMap) (float64, []ageName) {
	agesonly := []int{}
	n := 0
	for k, v := range ages {
		for i := 0; i < len(v); i++ {
			agesonly = append(agesonly, k)
			n++
		}
	}
	sort.Ints(agesonly)
	if n%2 == 0 {
		slot := int(math.Floor(float64(n)/2) - 1)
		median := float64(agesonly[slot]+agesonly[slot+1]) / 2
		names := ages[agesonly[slot]]
		for _, name := range ages[agesonly[slot+1]] {
			names = append(names, name)
		}
		return median, names
	} else {
		slot := int(math.Floor(float64(n) / 2.0))
		return float64(agesonly[slot]), ages[agesonly[slot]]
	}
}
func reduce(in <-chan string, done chan<- jetsam.DoneChanMsg) {
	n := 0
	average := 0
	ages := map[int][]ageName{}
	for {
		line, ok := <-in
		if ok {
			parts := strings.Split(line, ",")
			if len(parts) != 3 {
				log.Printf("Less than 3: %s\n", line)
				continue
			}
			age, err := strconv.Atoi(strings.TrimSpace(parts[2]))
			if err != nil {
				log.Printf("Age conversion failed. |%s| %v\n", strings.TrimSpace(parts[2]), err)
				continue
			}
			an := ageName{
				fname: strings.TrimSpace(parts[0]),
				lname: strings.TrimSpace(parts[1]),
			}
			if val, ok := ages[age]; ok {
				ages[age] = append(val, an)
			} else {
				ages[age] = []ageName{an}
			}
			average = average + age

			n++
		} else {
			fmt.Printf("Reducing input queue empty. Stopping\n")
			break
		}
	}
	if len(ages) == 0 {
		done <- jetsam.DoneChanMsg{
			Text:  "no content found",
			Count: 0,
		}
		return
	}
	age, who := findMedium(ages)
	names := who[0].fname + " " + who[0].lname
	for i := 1; i < len(who); i++ {
		names = names + ", " + who[i].fname + " " + who[i].lname
	}
	done <- jetsam.DoneChanMsg{
		Text: fmt.Sprintf(
			"Average: %f | Medium:%f for %s.",
			float64(average)/float64(n),
			age,
			names),
		Count: n,
	}
}
