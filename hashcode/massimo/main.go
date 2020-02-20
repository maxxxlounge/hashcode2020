package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Library struct {
	//BooksNr           int
	//Index int
	SighUpDay         int
	MaxBookScanPerDay int
	Books             []int
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

	defer file.Close()

	r := bufio.NewReader(file)

	i := 0
	var scanPerDayNr int
	var bookPerDayNr int

	for {
		var buffer bytes.Buffer

		rowb, _, err := r.ReadLine()
		if err != nil {
			log.Fatal(err)
			break
		}
		buffer.Write(rowb)
		row := buffer.String()
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
		case 1:
			pieces := strings.Split(row, " ")
			for j, p := range pieces {
				pi, err := strconv.Atoi(p)
				if err != nil {
					return nil, nil, 0, 0, 0, err
				}
				AllBooks = append(AllBooks, Book{Index: j, Score: pi, SentForScan: false})
			}
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
		i++
	}

	return

}

var ff []string

// Return points and errors
func elaborate() (int, []int, error) {
	panic("not implemented")
}

func writeSolution(filename string, solution []int) error {
	filename = strings.Replace(filename, ".in", ".out", -1)
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = outFile.WriteString(strconv.Itoa(len(solution)) + "\n")
	if err != nil {
		return err
	}
	for _, pt := range solution {
		_, err = outFile.WriteString(strconv.Itoa(pt) + " ")
		if err != nil {
			return err
		}
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

	/*point, solution, err = elaborate()
	if err != nil {
		return err
	}
	fmt.Printf("point:%v perc:%.2f \n", point, float64(len(solution)/len(p)))

	err = writeSolution(filename, solution)
	return err*/
	return nil
}

func main() {
	var err error

	err = ProcessFile("../input/b.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	/*
		err = ProcessFile("b_small.in")
		if err != nil {
			log.Fatal(err.Error())
		}
		err = ProcessFile("c_medium.in")
		if err != nil {
			log.Fatal(err.Error())
		}
		err = ProcessFile("d_quite_big.in")
		if err != nil {
			log.Fatal(err.Error())
		}
		err = ProcessFile("e_also_big.in")
		if err != nil {
			log.Fatal(err.Error())
		}*/

}
