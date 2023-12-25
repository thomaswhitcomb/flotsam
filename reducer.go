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
	for k, _ := range ages {
		agesonly = append(agesonly, k)
		n++
	}
	sort.Ints(agesonly)
	if n%2 == 0 {
		slot := int(math.Floor(float64(n)/2) - 1)

		avg := float64(agesonly[slot]+agesonly[slot+1]) / 2
		return avg, nil
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
	if who != nil {
		done <- jetsam.DoneChanMsg{
			Text: fmt.Sprintf(
				"Average: %f | Medium:%f for %s %s.",
				float64(average)/float64(n),
				age,
				who[0].fname,
				who[0].lname),
			Count: n,
		}
	} else {
		done <- jetsam.DoneChanMsg{
			Text: fmt.Sprintf(
				"Average: %f | Medium:%f ",
				float64(average)/float64(n),
				age),
			Count: n,
		}
	}
}
