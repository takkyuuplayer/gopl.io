package main

import (
	"reflect"
	"sort"
	"testing"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)

	return order
}

func TestTopoSort(t *testing.T) {
	want := []string{
		"intro to programming",
		"discrete math",
		"data structures",
		"algorithms",
		"linear algebra",
		"calculus",
		"formal languages",
		"computer organization",
		"compilers",
		"databases",
		"operating systems",
		"networks",
		"programming languages",
	}

	if !reflect.DeepEqual(topoSort(prereqs), want) {
		t.Errorf(`topoSort(prereqs) = %#v, want %#v`, topoSort(prereqs), want)
	}
}
