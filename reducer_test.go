package main

import (
	"strconv"
	"testing"

	"github.com/thomaswhitcomb/jetsam"
)

func atof(s string) float64 {
	val, _ := strconv.ParseFloat(s, 64)
	return val
}

func TestMakeAgeArray(t *testing.T) {
	ages := ageNameMap{5: []ageName{ageName{fname: "foo", lname: "bar"}}}
	ints := makeAgeArray(ages)
	if len(ints) != len(ages) {
		t.Errorf("Incorrect ages int[] array length: %d", len(ints))
	}
}

func TestIsEven(t *testing.T) {
	if isEven(42) != true {
		t.Errorf("You can't even get a simple TRUE boolean function working")
	}
	if isEven(17) != false {
		t.Errorf("You can't even get a simple FALSE boolean function working")
	}
}

func TestMedianOdd(t *testing.T) {
	ages := ageNameMap{
		1: []ageName{ageName{fname: "fname1", lname: "lname1"}},
		2: []ageName{ageName{fname: "fname2", lname: "lname2"}},
		3: []ageName{ageName{fname: "fname3", lname: "lname3"}},
	}
	m, who := calculateMedian(ages)
	if m != 2 {
		t.Errorf("Bad median calculation: %f\n", m)
	}
	x := ageName{fname: "fname2", lname: "lname2"}
	if who[0] != x {
		t.Errorf("Bad median people selection: %v\n", who)
	}
}
func TestMedianEven(t *testing.T) {
	ages := ageNameMap{
		1: []ageName{ageName{fname: "fname1", lname: "lname1"}},
		2: []ageName{ageName{fname: "fname2", lname: "lname2"}},
		3: []ageName{ageName{fname: "fname3", lname: "lname3"}},
		4: []ageName{ageName{fname: "fname4", lname: "lname4"}},
	}
	m, who := calculateMedian(ages)
	if m != 2.5 {
		t.Errorf("Bad median calculation: %f\n", m)
	}
	x := ageName{fname: "fname2", lname: "lname2"}
	y := ageName{fname: "fname3", lname: "lname3"}
	if who[0] != x || who[1] != y {
		t.Errorf("Bad median people selection: %v\n", who)
	}
}

func TestParseLine(t *testing.T) {
	age, an := parseLine("dinah,dog,14")
	if age != 14 {
		t.Errorf("Bad age parse: %d\n", age)
	}
	x := ageName{fname: "dinah", lname: "dog"}
	if an != x {
		t.Errorf("Bad age name parse: %v\n", an)
	}

}
func TestParseLineBadAge(t *testing.T) {
	age, an := parseLine("dinah,dog,14x")
	if age != -1 {
		t.Errorf("Expecting -1 for age, got: %d\n", age)
	}
	x := ageName{}
	if an != x {
		t.Errorf("Expecting empty agename, got: %v\n", an)
	}

}
func TestParseLineBadColumnCount1(t *testing.T) {
	age, an := parseLine("dinah,dog,14,abc")
	if age != -1 {
		t.Errorf("Expecting -1 for age, got: %d\n", age)
	}
	x := ageName{}
	if an != x {
		t.Errorf("Expecting empty agename, got: %v\n", an)
	}

}
func TestParseLineBadColumnCount2(t *testing.T) {
	age, an := parseLine("dinah,dog")
	if age != -1 {
		t.Errorf("Expecting -1 for age, got: %d\n", age)
	}
	x := ageName{}
	if an != x {
		t.Errorf("Expecting empty agename, got: %v\n", an)
	}

}
func TestReduceHappy(t *testing.T) {
	in := make(chan string, 3)
	done := make(chan jetsam.DoneChanMsg, 1)

	in <- "dinah,dog,14"
	close(in)
	reduce(in, done)
	msg := <-done
	if atof(msg.Results["average"]) != 14 {
		t.Errorf("Expecting average of 14, got %f\n", atof(msg.Results["average"]))
	}
	if atof(msg.Results["median"]) != 14 {
		t.Errorf("Expecting average of 14, got %f\n", atof(msg.Results["average"]))
	}
	if msg.NumberProcessed != 1 {
		t.Errorf("Expecting NumberProcessed of 1, got: %d\n", msg.NumberProcessed)
	}
}
func TestReduceSad(t *testing.T) {
	in := make(chan string, 3)
	done := make(chan jetsam.DoneChanMsg, 1)

	close(in)
	reduce(in, done)
	msg := <-done
	if msg.NumberProcessed != 0 {
		t.Errorf("Expecting NumberProcessed of 1, got: %d\n", msg.NumberProcessed)
	}
}
