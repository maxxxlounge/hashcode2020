package processor

import "sort"

func contains(arr []int, str int) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func (d *ReadyToProcess) Elaborate() (Solution, error) {
	s := Solution{}
	sort.Slice(d.Libraries, func(i, j int) bool {
		return d.Libraries[i].MedianValue > d.Libraries[j].MedianValue
	})

	booksIndexAlreadyScan := make([]int, 0)
	librariesReady := make([]*Library, 0)
	librariesReadyArr := make([]int, 0)

	isSignupActive := -1
	outputLibraries := make(map[int]OutputLibrary, 0)

	for i := 0; i < d.DayForScanning; i++ {
		//signup
		if isSignupActive == -1 {
			for j, library := range d.Libraries {
				if library.RegisteryDayLeft == library.MaxBookScanPerDay {
					d.Libraries[j].RegisteryDayLeft--
					isSignupActive = library.Index
				}
			}
		} else {
			if d.Libraries[isSignupActive].RegisteryDayLeft == 0 {
				librariesReady = append(librariesReady, &d.Libraries[isSignupActive])
				librariesReadyArr = append(librariesReadyArr, d.Libraries[isSignupActive].Index)
				isSignupActive = -1
			} else {
				d.Libraries[isSignupActive].RegisteryDayLeft--
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

			for _, book := range readyLibrary.Books {
				counter := 0
				if !contains(booksIndexAlreadyScan, book.Index) && counter < readyLibrary.MaxBookScanPerDay {
					outLibrary.Books = append(outLibrary.Books, book.Index)
					booksIndexAlreadyScan = append(booksIndexAlreadyScan, book.Index)
					counter++
				}
			}
			outputLibraries[readyLibrary.Index] = outLibrary
		}

	}
	s.Libraries = outputLibraries
	return s, nil
}
