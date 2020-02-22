package main

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type OutputLibrary struct {
	ID int
	Books []int
}

type Library struct {
	//BooksNr           int
	Index             int
	SighUpDay         int
	RegisteryDayLeft	int
	MaxBookScanPerDay int
	Books             []*Book
	MedianValue		float64
	ScannedBook []int
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
	var libraryIndex int
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
				libraryIndex = int(math.Floor(float64(i/2)) - 1)
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
				bb := make([]*Book, 0)
				medianScore := 0
				for _, p := range pieces {
					pr, err := strconv.Atoi(p)
					if err != nil {
						return nil, nil, 0, 0, 0, err
					}
					book := &AllBooks[pr]
					medianScore += book.Score
					bb = append(bb, book)
				}
				sort.Slice(bb, func(i, j int) bool {
					return bb[i].Score > bb[j].Score
				})
				l := Library{
					Index:	libraryIndex,
					SighUpDay:         scanPerDayNr,
					RegisteryDayLeft: 	scanPerDayNr,
					MaxBookScanPerDay: bookPerDayNr,
					Books:             bb,
					MedianValue: float64(medianScore)/float64(len(bb)),
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
	_, Libraries, _, _, DayForScanning, err := parseInputFromFile(filename)
	if err != nil {
		return err
	}
/*
	fmt.Printf("AllBooks %+v\n", AllBooks)
	fmt.Printf("Libraries %+v\n", Libraries)
	fmt.Printf("LibraryCount %+v\n", LibraryCount)
	fmt.Printf("BooksCount %+v\n", BooksCount)
	fmt.Printf("DayForScanning %+v\n", DayForScanning)
*/

	sort.Slice(Libraries, func(i, j int) bool {
		return Libraries[i].MedianValue > Libraries[j].MedianValue
	})

	booksIndexAlreadyScan := make([]int, 0)
	librariesReady := make([]*Library, 0)
	librariesReadyArr := make([]int, 0)

	isSignupActive := -1
	outputLibraries := make(map[int]OutputLibrary,0)

	for i:= 0; i< DayForScanning; i++ {
		//signup
		if isSignupActive == -1 {
			for j, library := range Libraries {
				if library.RegisteryDayLeft == library.MaxBookScanPerDay {
					Libraries[j].RegisteryDayLeft--
					isSignupActive = library.Index
				}
			}
		} else {
			if Libraries[isSignupActive].RegisteryDayLeft == 0 {
				librariesReady = append(librariesReady, &Libraries[isSignupActive])
				librariesReadyArr = append(librariesReadyArr,Libraries[isSignupActive].Index)
				isSignupActive = -1
			} else {
				Libraries[isSignupActive].RegisteryDayLeft--
			}
		}

		//libri
		for _, readyLibrary := range librariesReady {
			var outLibrary OutputLibrary
			outLibrary, ok := outputLibraries[readyLibrary.Index]
			if !ok {
				outLibrary = OutputLibrary{
					ID:readyLibrary.Index,
					Books:make([]int,0),
				}
			}

			for _,book := range readyLibrary.Books {
				counter := 0
				if !contains(booksIndexAlreadyScan,book.Index) && counter < readyLibrary.MaxBookScanPerDay {
					outLibrary.Books = append(outLibrary.Books,book.Index)
					booksIndexAlreadyScan = append(booksIndexAlreadyScan,book.Index)
					counter++
				}
			}
			outputLibraries[readyLibrary.Index] = outLibrary
		}

	}

	ll := strconv.Itoa( len(outputLibraries) )+ "\n"
	for _, out := range outputLibraries {
		ll += strconv.Itoa(out.ID) + " " + strconv.Itoa(len(out.Books)) + "\n"
		for _, b := range out.Books {
			ll += strconv.Itoa(b) + " "
		}
		ll += "\n"
	}

	ll = strings.Replace(ll, " \n", "\n", -1)
	err = writeSolution(filename, ll)
	/*
	var libsToSigh []Library
	remaingDay := DayForScanning
	for _, l := range Libraries {
		if remaingDay == 1 {
			break
		}
		if l.SighUpDay < remaingDay {
			libsToSigh = append(libsToSigh, l)
			remaingDay -= l.SighUpDay
		}
	}
	ll := strconv.Itoa(len(libsToSigh)) + "\n"
	for _, l := range libsToSigh {
		llbis := ""
		counter := 0
		for _, b := range l.Books {
			if !contains(booksIndexAlreadyScan, b.Index) {
				llbis += strconv.Itoa(b.Index) + " "
				counter++
				booksIndexAlreadyScan = append(booksIndexAlreadyScan, b.Index)
			}
		}
		if llbis != "" {
			ll = ll + strconv.Itoa(l.Index) + " " + strconv.Itoa(counter) + "\n" + llbis + "\n"
		}
	}
	ll = strings.Replace(ll, " \n", "\n", -1)
	err = writeSolution(filename, ll)
	 */
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

func contains(arr []int, str int) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
