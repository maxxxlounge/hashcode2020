package main

import (
	"bufio"
	"fmt"
	"github.com/maxxxlounge/hashcode2020/max/genetic"
	"log"
	"os"
	"strconv"
	"strings"
)

type Solution struct {
	pieces []int
	types  int
}

func parseInputFromFile(filename string) (int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, err
	}
	var maxSlice int
	var pp []int
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		row := scanner.Text()
		if i == 0 {
			pieces := strings.Split(row, " ")
			maxSlice, err = strconv.Atoi(pieces[0])
			if err != nil {
				return 0, nil, err
			}
			pl, err := strconv.Atoi(pieces[1])
			if err != nil {
				return 0, nil, err
			}
			pp = make([]int, pl)
		} else {
			slices := strings.Split(row, " ")
			for y, slice := range slices {
				intSlice, err := strconv.Atoi(slice)
				if err != nil {
					return 0, nil, err
				}
				pp[y] = intSlice
			}
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		return 0, nil, err
	}

	return maxSlice, pp, nil
}

var ff []string

// Return points and errors
func elaborate(maxSlice int, p []int) (int, []int, error) {
	var s []int
	pp := genetic.GeneratePopulation(maxSlice, 500, p)
	found := false
	for !found {
		pop := genetic.NaturalSelection(p)
	}

	//outputType := 0
	/*for i := len(p) - 1; i > 0; i-- {
		if pp+p[i] > maxSlice {
			continue
		}
		pp += p[i]
		s = append(s, i)
		outputType++
	}*/
	return pp, s, nil
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
	maxSlices, p, err := parseInputFromFile(filename)
	if err != nil {
		return err
	}
	point := 0
	var solution []int

	point, solution, err = elaborate(maxSlices, p)
	if err != nil {
		return err
	}
	fmt.Printf("point:%v perc:%.2f \n", point, float64(len(solution)/len(p)))

	err = writeSolution(filename, solution)
	return err
}

func main() {
	var err error

	err = ProcessFile("a_example.in")
	if err != nil {
		log.Fatal(err.Error())
	}
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
	}

}
