package processor

import (
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ParseInput(filename string) (out *ReadyToProcess, err error) {
	data := ReadyToProcess{}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var scanPerDayNr int
	var bookPerDayNr int
	var libraryIndex int
	lines := strings.Split(string(f), "\n")

	for i, row := range lines {
		row = strings.Trim(row, " ")
		if row == "" {
			break
		}
		switch i {
		case 0:
			pieces := strings.Split(row, " ")
			data.BooksCount, err = strconv.Atoi(pieces[0])
			if err != nil {
				return nil, err
			}
			data.LibraryCount, err = strconv.Atoi(pieces[1])
			if err != nil {
				return nil, err
			}
			data.DayForScanning, err = strconv.Atoi(pieces[2])
			if err != nil {
				return nil, err
			}
			break
		case 1:
			pieces := strings.Split(row, " ")
			for j, p := range pieces {
				pi, err := strconv.Atoi(p)
				if err != nil {
					return nil, err
				}
				data.AllBooks = append(data.AllBooks, Book{Index: j, Score: pi, SentForScan: false})
			}
			break
		default:
			pieces := strings.Split(row, " ")
			if i%2 == 0 {
				libraryIndex = int(math.Floor(float64(i/2)) - 1)
				scanPerDay := pieces[1]
				scanPerDayNr, err = strconv.Atoi(scanPerDay)
				if err != nil {
					return nil, err
				}
				bookPerDay := pieces[2]
				bookPerDayNr, err = strconv.Atoi(bookPerDay)
				if err != nil {
					return nil, err
				}
			} else {
				bb := make([]*Book, 0)
				medianScore := 0
				for _, p := range pieces {
					pr, err := strconv.Atoi(p)
					if err != nil {
						return nil, err
					}
					book := &data.AllBooks[pr]
					medianScore += book.Score
					bb = append(bb, book)
				}
				sort.Slice(bb, func(i, j int) bool {
					return bb[i].Score > bb[j].Score
				})
				l := Library{
					Index:             libraryIndex,
					SighUpDay:         scanPerDayNr,
					RegisteryDayLeft:  scanPerDayNr,
					MaxBookScanPerDay: bookPerDayNr,
					Books:             bb,
					MedianValue:       float64(medianScore) / float64(len(bb)),
				}
				data.Libraries = append(data.Libraries, l)
			}
		}
	}
	return &data, nil
}
