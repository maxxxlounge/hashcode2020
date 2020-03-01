package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const Debug = false

type OutputLibrary struct {
	ID int
	Books []int
}

type Library struct {
	//BooksNr           int
	Index             int
	SignUpDay         int
	RegisteryDayLeft  int
	MaxBookScanPerDay int
	Books             []*Book
	GoodnessIndex       float64
	ScannedBook       []int
	TotalBookValue int
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

	var signUpDays int
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
				signUpDays, err = strconv.Atoi(scanPerDay)
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
				for _, p := range pieces {
					pr, err := strconv.Atoi(p)
					if err != nil {
						return nil, nil, 0, 0, 0, err
					}
					book := &AllBooks[pr]
					bb = append(bb, book)
				}
				score := 0
				sort.Slice(bb, func(i, j int) bool {
					return bb[i].Score > bb[j].Score
				})
				for k:=0; k<(DayForScanning-signUpDays) * bookPerDayNr && k < len(bb); k++ {
					score += bb[k].Score
				}
				l := Library{
					Index:             libraryIndex,
					SignUpDay:         signUpDays,
					RegisteryDayLeft:  signUpDays,
					MaxBookScanPerDay: bookPerDayNr,
					Books:             bb,
					GoodnessIndex: float64(score) / float64(signUpDays),
					TotalBookValue: score,
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

	booksIndexAlreadyScan := make(map[int]bool, 0)
	librariesReady := make([]*Library, 0)
	librariesReadyArr := make([]int, 0)

	isSignupActive := -1
	outputLibraries := make(map[int]OutputLibrary,0)
	print("%+v\n\n", Libraries)

	for i:= 0; i< DayForScanning; i++ {
		sort.Slice(Libraries, func(i, j int) bool {
			return Libraries[i].GoodnessIndex > Libraries[j].GoodnessIndex
		})

		//signup
		print("Day: %d\n", i)
		if i % 1000 == 0 {
			fmt.Printf("Day: %d/%d\n", i, DayForScanning)
		}
		if isSignupActive != -1 {
			if Libraries[isSignupActive].RegisteryDayLeft == 0 {
				librariesReady = append(librariesReady, &(Libraries[isSignupActive]))
				librariesReadyArr = append(librariesReadyArr,Libraries[isSignupActive].Index)
				print("\tSignup complete: %d\n", isSignupActive)
				isSignupActive = -1

			} else {
				Libraries[isSignupActive].RegisteryDayLeft--
				print("\tDay left: %d\n", Libraries[isSignupActive].RegisteryDayLeft)
			}
		}
		if isSignupActive == -1 {
			for j, library := range Libraries {
				if library.RegisteryDayLeft == library.SignUpDay && library.SignUpDay < (DayForScanning - i){
					print("\tSignup library: %d\n", library.Index)
					Libraries[j].RegisteryDayLeft--
					isSignupActive = j
					print("\tDay left: %d\n", Libraries[j].RegisteryDayLeft)
					break
				}
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

			counter := 0
			for _,book := range readyLibrary.Books {
				if !contains(booksIndexAlreadyScan,book.Index) && counter < readyLibrary.MaxBookScanPerDay {
					outLibrary.Books = append(outLibrary.Books,book.Index)
					booksIndexAlreadyScan[book.Index] = true
					counter++
				}
			}
			print("\t%+v\n", outLibrary)
			outputLibraries[readyLibrary.Index] = outLibrary
		}

		//recalculate
		for _, l := range Libraries {
			l.RecalculateGoodnessIndex(DayForScanning - i - 1, &booksIndexAlreadyScan)
		}

	}
	print("%+v\n", outputLibraries)
	ll := ""
	//ll := strconv.Itoa( len(outputLibraries) )+ "\n"
	n :=0
	for _, out := range outputLibraries {
		if len(out.Books) > 0 {
			ll += strconv.Itoa(out.ID) + " " + strconv.Itoa(len(out.Books)) + "\n"
			for _, b := range out.Books {
				ll += strconv.Itoa(b) + " "
			}
			ll += "\n"
			n++
		}
	}
	ll = strconv.Itoa( n )+ "\n" + ll
	ll = strings.Replace(ll, " \n", "\n", -1)
	err = writeSolution(filename, ll)

	return nil
}

func main() {
	var err error
	if len(os.Args) != 2 {
		log.Fatalf("usage: go run main <filename>")
	}

	filename := os.Args[1]

	err = ProcessFile("../input/" + filename + ".txt")
	if err != nil {
		err = errors.Wrap(err, "error processing a.txt")
		log.Fatal(err.Error())
	}
}

func contains(list map[int]bool, search int) bool {
	_, exists := list[search]
	return exists
}

func print(format string, a ...interface{}) {
	if Debug {
		fmt.Printf(format, a...)
	}
}

func (l *Library) RecalculateGoodnessIndex(daysForScanning int, alreadyRegisteredBooks *map[int]bool) {

	// solo non registrate
	if l.RegisteryDayLeft == l.SignUpDay {
		if l.SignUpDay > daysForScanning {
			l.GoodnessIndex = -99999
		} else {
			bookScore := 0
			c := 0
			for k := 0; c <= (daysForScanning-l.SignUpDay) * l.MaxBookScanPerDay && k < len(l.Books); k++ {
				if !contains(*alreadyRegisteredBooks, l.Books[k].Index) {
					bookScore += l.Books[k].Score
					c++
				}
			}
			l.GoodnessIndex = float64(bookScore)  / float64(l.SignUpDay)
			l.TotalBookValue = bookScore
		}
	}
}