package main

type compare func(x, y *Track) int

var byTitle compare = func(x, y *Track) int {
	if x.Title == y.Title {
		return 0
	} else if x.Title < y.Title {
		return 1
	} else {
		return -1
	}
}

var byArtist compare = func(x, y *Track) int {
	if x.Artist == y.Artist {
		return 0
	} else if x.Artist < y.Artist {
		return 1
	} else {
		return -1
	}
}

var byAlbum compare = func(x, y *Track) int {
	if x.Album == y.Album {
		return 0
	} else if x.Album < y.Album {
		return 1
	} else {
		return -1
	}
}
var byYear compare = func(x, y *Track) int {
	if x.Year == y.Year {
		return 0
	} else if x.Year < y.Year {
		return 1
	} else {
		return -1
	}
}

type selectableSorter struct {
	Tracks []*Track
	order  []compare
}

func (s *selectableSorter) Len() int           { return len(s.Tracks) }
func (s *selectableSorter) Swap(i, j int)      { s.Tracks[i], s.Tracks[j] = s.Tracks[j], s.Tracks[i] }
func (s *selectableSorter) Less(i, j int) bool { return s.less(s.Tracks[i], s.Tracks[j]) }

func (s *selectableSorter) less(x, y *Track) bool {
	for _, c := range s.order {
		r := c(x, y)
		if r == 0 {
			continue
		} else if r > 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

func (s *selectableSorter) Select(c compare) {
	s.order = append(s.order, c)
}

func (s *selectableSorter) Clear() {
	s.order = []compare{}
}
