package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Library struct {
	//BooksNr           int
	Index             int
	SighUpDay         int
	MaxBookScanPerDay int
	Books             []int
	BookToScan        []int
	BookPoints        int
	AvgSighUp         int
}

type Book struct {
	Index       int
	Score       int
	SentForScan bool
}

func parseInputFromFile(filename string) (AllBooks []Book, Libraries []Library, LibraryCount, BooksCount, DayForScanning int, err error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, 0, 0, 0, err
	}

	out, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, nil, 0, 0, 0, err
	}

	defer file.Close()

	var scanPerDayNr int
	var bookPerDayNr int

	lines := strings.Split(string(out), "\n")

	for i, row := range lines {
		row = strings.Trim(row, " ")
		if row == "" {
			break
		}
		switch i {
		case 0:
			pieces := strings.Split(row, " ")
			BooksCount, err = strconv.Atoi(pieces[0])
			if err != nil {
				return nil, nil, 0, 0, 0, err
			}
			LibraryCount, err = strconv.Atoi(pieces[1])
			if err != nil {
				return nil, nil, 0, 0, 0, err
			}
			DayForScanning, err = strconv.Atoi(pieces[2])
			if err != nil {
				return nil, nil, 0, 0, 0, err
			}
			break
		case 1:
			pieces := strings.Split(row, " ")
			for j, p := range pieces {
				pi, err := strconv.Atoi(p)
				if err != nil {
					return nil, nil, 0, 0, 0, err
				}
				AllBooks = append(AllBooks, Book{Index: j, Score: pi, SentForScan: false})
			}
			break
		default:
			pieces := strings.Split(row, " ")

			if i%2 == 0 {
				scanPerDay := pieces[1]
				scanPerDayNr, err = strconv.Atoi(scanPerDay)
				if err != nil {
					return nil, nil, 0, 0, 0, err
				}
				bookPerDay := pieces[2]
				bookPerDayNr, err = strconv.Atoi(bookPerDay)
				if err != nil {
					return nil, nil, 0, 0, 0, err
				}
			} else {
				bb := []int{}
				for _, p := range pieces {
					pr, err := strconv.Atoi(p)
					if err != nil {
						return nil, nil, 0, 0, 0, err
					}
					bb = append(bb, pr)
				}
				l := Library{
					SighUpDay:         scanPerDayNr,
					MaxBookScanPerDay: bookPerDayNr,
					Books:             bb,
				}
				Libraries = append(Libraries, l)
			}
		}
	}
	return
}

func writeSolution(filename string, solution string) error {
	filename = strings.Replace(filename, ".txt", ".txt.out", -1)
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = outFile.Write([]byte(solution))
	if err != nil {
		return err
	}
	return nil
}

func ProcessFile(filename string) error {
	AllBooks, Libraries, LibraryCount, BooksCount, DayForScanning, err := parseInputFromFile(filename)
	if err != nil {
		return err
	}

	fmt.Printf("AllBooks %+v\n", AllBooks)
	fmt.Printf("Libraries %+v\n", Libraries)
	fmt.Printf("LibraryCount %+v\n", LibraryCount)
	fmt.Printf("BooksCount %+v\n", BooksCount)
	fmt.Printf("DayForScanning %+v\n", DayForScanning)

	var libsToSigh []Library
	remaingDay := DayForScanning

	for li := 0; li < len(Libraries); li++ {
		l := Libraries[li]
		if remaingDay == 1 {
			break
		}
		if l.SighUpDay < remaingDay {
			l.Index = li
			libsToSigh = append(libsToSigh, l)
			remaingDay -= l.SighUpDay
		}
	}

	//sighuptimeavg := sighuptimesum / len(Libraries)

	scannedBook := make(map[int]bool)
	libtosighNotEmpty := []Library{}
	sort.Slice(libsToSigh, func(i, j int) bool {
		return (libsToSigh[i].MaxBookScanPerDay / libsToSigh[i].SighUpDay) > (libsToSigh[j].MaxBookScanPerDay / libsToSigh[j].SighUpDay)
	})
	for _, l := range libsToSigh {
		l.BookToScan = []int{}
		for i := 0; i < len(l.Books); i++ {
			if scannedBook[l.Books[i]] == false {
				l.BookToScan = append(l.BookToScan, l.Books[i])
				scannedBook[l.Books[i]] = true
			}
		}
		sort.Slice(l.BookToScan, func(i, j int) bool {
			return l.BookToScan[i] > l.BookToScan[j]
		})
		if len(l.BookToScan) == 0 {
			continue
		}
		libtosighNotEmpty = append(libtosighNotEmpty, l)
	}

	ll := ""
	ll += strconv.Itoa(len(libtosighNotEmpty)) + "\n"
	for _, l := range libtosighNotEmpty {
		ll += strconv.Itoa(l.Index) + " " + strconv.Itoa(len(l.BookToScan)) + "\n"
		for bi, b := range l.BookToScan {
			ll += strconv.Itoa(b)
			if bi < len(l.BookToScan)-1 {
				ll += " "
			}
		}
		ll += "\n"
	}
	err = writeSolution(filename, ll)

	return nil
}

func main() {
	var err error

	err = ProcessFile("../input/a.txt")
	if err != nil {
		err = errors.Wrap(err, "error processing a.txt")
		log.Fatal(err.Error())
	}
	err = ProcessFile("../input/b.txt")
	if err != nil {
		err = errors.Wrap(err, "error processing b.txt")
		log.Fatal(err.Error())
	}
	err = ProcessFile("../input/c.txt")
	if err != nil {
		err = errors.Wrap(err, "error processing c.txt")
		log.Fatal(err.Error())
	}
	err = ProcessFile("../input/d.txt")
	if err != nil {
		err = errors.Wrap(err, "error processing d.txt")
		log.Fatal(err.Error())
	}
	err = ProcessFile("../input/e.txt")
	if err != nil {
		err = errors.Wrap(err, "error processing e.txt")
		log.Fatal(err.Error())
	}
	err = ProcessFile("../input/f.txt")
	if err != nil {
		err = errors.Wrap(err, "error processing f.txt")
		log.Fatal(err.Error())
	}

}
