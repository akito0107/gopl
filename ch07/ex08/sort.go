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

type selectableSorter struct {
	t     []*Track
	order []compare
}

func (s *selectableSorter) Len() int           { return len(s.t) }
func (s *selectableSorter) Swap(i, j int)      { s.t[i], s.t[j] = s.t[j], s.t[i] }
func (s *selectableSorter) Less(i, j int) bool { return s.less(s.t[i], s.t[j]) }

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

type stableSorter struct {
	idx   int
	t     []*Track
	order []compare
}

func (s *stableSorter) Len() int           { return len(s.t) }
func (s *stableSorter) Swap(i, j int)      { s.t[i], s.t[j] = s.t[j], s.t[i] }
func (s *stableSorter) Less(i, j int) bool { return s.less(s.t[i], s.t[j]) }

func (s *stableSorter) Select(c compare) {
	s.order = append(s.order, c)
	s.idx += 1
}

func (s *stableSorter) Clear() {
	s.order = []compare{}
	s.idx = 0
}

func (s *stableSorter) less(x, y *Track) bool {
	if len(s.order) <= 0 {
		panic(s)
	}
	c := s.order[0]
	r := c(x, y)
	if r > 0 {
		return true
	} else {
		return false
	}
}

func (s *stableSorter) HasNext() bool {
	s.idx--
	s.order = s.order[1:]
	return s.idx > 0
}
