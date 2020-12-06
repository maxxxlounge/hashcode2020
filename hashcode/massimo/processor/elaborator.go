package processor

import (
	"fmt"
	"sort"
)

const Debug = false

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
		bookScore := 0
		for k := 0; k < daysForScanning-l.SignUpDay && k < len(l.Books) && !contains(*alreadyRegisteredBooks, l.Books[k].Index); k++ {
			bookScore += l.Books[k].Score
		}
		l.GoodnessIndex = float64(bookScore) * float64(l.MaxBookScanPerDay) / float64(l.SignUpDay)
	}
}

func (d *ReadyToProcess) Elaborate() (Solution, error) {
	s := Solution{}
	booksIndexAlreadyScan := make(map[int]bool, 0)
	librariesReady := make([]*Library, 0)
	librariesReadyArr := make([]int, 0)

	isSignupActive := -1
	outputLibraries := make(map[int]OutputLibrary, 0)
	sort.Slice(d.Libraries, func(i, j int) bool {
		return d.Libraries[i].GoodnessIndex > d.Libraries[j].GoodnessIndex
	})
	print("%+v\n\n", s.Libraries)
	for i := 0; i < d.DayForScanning; i++ {
		//signup

		print("Day: %d\n", i)
		if i%1000 == 0 {
			fmt.Printf("Day: %d/%d\n", i, d.DayForScanning)
		}
		if isSignupActive == -1 {
			for j, library := range d.Libraries {
				if library.RegisteryDayLeft == library.SignUpDay && library.SignUpDay < (d.DayForScanning-i) {
					print("\tSignup library: %d\n", library.Index)
					d.Libraries[j].RegisteryDayLeft--
					isSignupActive = j
					print("\tDay left: %d\n", d.Libraries[j].RegisteryDayLeft)
					break
				}
			}
		} else {
			if d.Libraries[isSignupActive].RegisteryDayLeft == 0 {
				librariesReady = append(librariesReady, &(d.Libraries[isSignupActive]))
				librariesReadyArr = append(librariesReadyArr, d.Libraries[isSignupActive].Index)
				print("\tSignup complete: %d\n", isSignupActive)
				isSignupActive = -1

			} else {
				d.Libraries[isSignupActive].RegisteryDayLeft--
				print("\tDay left: %d\n", d.Libraries[isSignupActive].RegisteryDayLeft)
			}
		}

		//libri
		for _, readyLibrary := range librariesReady {
			var outLibrary OutputLibrary
			outLibrary, ok := outputLibraries[readyLibrary.Index]
			if !ok {
				outLibrary = OutputLibrary{
					ID:    readyLibrary.Index,
					Books: make([]int, 0),
				}
			}

			counter := 0
			for _, book := range readyLibrary.Books {
				if !contains(booksIndexAlreadyScan, book.Index) && counter < readyLibrary.MaxBookScanPerDay {
					outLibrary.Books = append(outLibrary.Books, book.Index)
					booksIndexAlreadyScan[book.Index] = true
					counter++
				}
			}
			print("\t%+v\n", outLibrary)
			outputLibraries[readyLibrary.Index] = outLibrary
		}

		//recalculate
		/*
			for _, l := range Libraries {
				l.RecalculateGoodnessIndex(DayForScanning - 1, &booksIndexAlreadyScan)
			}
		*/

	}
	s.Libraries = outputLibraries
	return s, nil
}
