package main

import (
	"fmt"
	"sort"
	"time"
)

type test struct {
	name string
	t    time.Time
}

func main() {
	tests := []*test{
		&test{"t1", time.Now()},
		&test{"t2", time.Now()},
	}

	fmt.Println(tests)

	sort.Slice(tests, func(i, j int) bool {
		return tests[i].t.Sub(tests[j].t) > 0
	})

	fmt.Println(tests)

	// t1 := time.Now()
	// t2 := time.Now()
	// fmt.Println(t1.Sub(t2))

	// // b := t1 < t2
}
