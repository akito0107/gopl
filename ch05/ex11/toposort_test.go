package main

import "testing"

func Test_topoSort(t *testing.T) {
	var prereqs = map[string][]string{
		"a1": {"a2"},
		"a2": {
			"a3",
			"a4",
		},
		"a5": {"a2"},
	}
	seen := map[string]bool{}

	order, err := topoSort(prereqs)
	if err != nil {
		t.Error(err)
	}
	for _, o := range order {
		req, ok := prereqs[o]

		// prereqなし
		if !ok {
			seen[o] = true
			continue
		}

		for _, r := range req {
			if !seen[r] {
				t.Errorf("toposort invlaid prereq %s: ", r)
			}
		}
		seen[o] = true
	}
}
func Test_topoSortCyclic(t *testing.T) {
	var prereqs = map[string][]string{
		"a1": {"a2"},
		"a2": {
			"a3",
			"a4",
		},
		"a5": {"a2"},
		"a4": {"a2"},
	}

	_, err := topoSort(prereqs)
	if err == nil {
		t.Error("must be cyclic error")
	}
}
