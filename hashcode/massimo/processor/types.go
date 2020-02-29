package processor

type OutputLibrary struct {
	ID    int
	Books []int
}

type Library struct {
	//BooksNr           int
	Index             int
	SignUpDay         int
	RegisteryDayLeft  int
	MaxBookScanPerDay int
	Books             []*Book
	MedianValue       float64
	ScannedBook       []int
	GoodnessIndex     float64
}

type Book struct {
	Index       int
	Score       int
	SentForScan bool
}

type Solution struct {
	Libraries map[int]OutputLibrary
}

type ReadyToProcess struct {
	AllBooks       []Book
	Libraries      []Library
	LibraryCount   int
	BooksCount     int
	DayForScanning int
}

func (l *Library) GetSumBookPoint() int {
	var s int
	for _, b := range l.Books {
		s += b.Score
	}
	return s
}
