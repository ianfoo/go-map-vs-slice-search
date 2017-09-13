package main

import (
	"fmt"
	"testing"
)

var sizes = []int{5, 10, 20, 50, 500, 2500}

type testCase struct {
	desc      string
	searchFor int
}

func BenchmarkMapSearch(b *testing.B) {
	for _, s := range sizes {
		m := generateMap(s)
		cases := makeTestCases(s)
		for _, tc := range cases {
			b.Run(tc.desc, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = mapSearch(tc.searchFor, m)
				}
			})
		}
	}
}

func BenchmarkSliceSearch(b *testing.B) {
	for _, s := range sizes {
		sl := generateSlice(s)
		cases := makeTestCases(s)
		for _, tc := range cases {
			b.Run(tc.desc, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = sliceSearch(tc.searchFor, sl)
				}
			})
		}
	}
}

func desc(desc string, count int) string {
	return fmt.Sprintf("%s (%d elements)", desc, count)
}

func makeTestCases(count int) []testCase {
	return []testCase{
		{desc: desc("beginning", count), searchFor: 1},
		{desc: desc("middle", count), searchFor: count / 2},
		{desc: desc("end", count), searchFor: count},
		{desc: desc("notfound", count), searchFor: count * 2},
	}
}

func generateMap(count int) map[int]struct{} {
	m := make(map[int]struct{}, count)
	for i := 0; i < count; i++ {
		m[i+1] = struct{}{}
	}
	return m
}

func generateSlice(count int) []int {
	s := make([]int, 0, count)
	for i := 0; i < count; i++ {
		s = append(s, i+1)
	}
	return s
}
