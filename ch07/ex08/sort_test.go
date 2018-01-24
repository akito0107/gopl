package main

import (
	"sort"
	"testing"
)

func TestSelectableSort(t *testing.T) {
	t.Run("single column case (byArtist)", func(t *testing.T) {
		track := []*Track{
			{Artist: "3"},
			{Artist: "1"},
			{Artist: "2"},
		}
		s := &selectableSorter{t: track}
		s.Select(byArtist)

		sort.Sort(s)
		if track[0].Artist != "1" {
			t.Errorf("must be sorted Artist Order, actually :%s", track[0].Artist)
		}
		if track[1].Artist != "2" {
			t.Errorf("must be sorted Artist Order, actually :%s", track[1].Artist)
		}
		if track[2].Artist != "3" {
			t.Errorf("must be sorted Artist Order, actually :%s", track[2].Artist)
		}
	})
	t.Run("multi column case (byTitle => byArtist)", func(t *testing.T) {
		track := []*Track{
			{Title: "2", Artist: "2"},
			{Title: "1", Artist: "3"},
			{Title: "2", Artist: "1"},
			{Title: "1", Artist: "1"},
			{Title: "3", Artist: "2"},
			{Title: "3", Artist: "1"},
		}
		s := &selectableSorter{t: track}
		s.Select(byTitle)
		s.Select(byArtist)
		sort.Sort(s)

		if track[0].Title != "1" && track[0].Artist != "1" {
			t.Errorf("must be sorted Artist Order, actually :%+v", track[0])
		}
		if track[1].Title != "1" && track[1].Artist != "3" {
			t.Errorf("must be sorted Artist Order, actually :%+v", track[1])
		}
		if track[2].Title != "2" && track[2].Artist != "1" {
			t.Errorf("must be sorted Artist Order, actually :%+v", track[2])
		}
		if track[3].Title != "2" && track[3].Artist != "2" {
			t.Errorf("must be sorted Artist Order, actually :%+v", track[3])
		}
		if track[4].Title != "3" && track[4].Artist != "1" {
			t.Errorf("must be sorted Artist Order, actually :%+v", track[4])
		}
		if track[5].Title != "3" && track[5].Artist != "2" {
			t.Errorf("must be sorted Artist Order, actually :%+v", track[5])
		}
	})
}
