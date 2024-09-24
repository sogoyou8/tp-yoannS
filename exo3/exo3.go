package exo

import (
	"sort"
)

type Interval struct {
	start int
	end   int
}

func Ft_non_overlap(intervals [][]int) int {
	intervalSlice := make([]Interval, len(intervals))
	for i, interval := range intervals {
		intervalSlice[i] = Interval{interval[0], interval[1]}
	}

	sort.Slice(intervalSlice, func(i, j int) bool {
		return intervalSlice[i].start < intervalSlice[j].start
	})

	prevEnd := intervalSlice[0].end
	removeCount := 0

	for i := 1; i < len(intervalSlice); i++ {
		if intervalSlice[i].start < prevEnd {
			removeCount++
		} else {
			prevEnd = intervalSlice[i].end
		}
	}

	return removeCount
}
