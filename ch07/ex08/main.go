package main

import "sort"

func main() {
	trackSort := make([]*Track, len(tracks))
	copy(trackSort, tracks)
	sorter := &selectableSorter{t: trackSort}
	sorter.Select(byTitle)
	sorter.Select(byArtist)
	sorter.Select(byAlbum)
	sort.Sort(sorter)
	printTracks(trackSort)

	trackStable := make([]*Track, len(tracks))
	copy(trackStable, tracks)
	stableSorter := &stableSorter{t: trackStable}
	stableSorter.Select(byTitle)
	stableSorter.Select(byArtist)
	stableSorter.Select(byAlbum)
	for stableSorter.HasNext() {
		sort.Stable(stableSorter)
	}
	printTracks(trackStable)
}
