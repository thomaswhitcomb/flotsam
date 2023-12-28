package main

import (
	"testing"
)

func TestMakeAgeArray(t *testing.T) {
	ages := ageNameMap{5: []ageName{ageName{fname: "foo", lname: "bar"}}}
	ints := makeAgeArray(ages)
	if len(ints) != len(ages) {
		t.Errorf("Incorrect ages int[] array length: %d", len(ints))
	}
}
