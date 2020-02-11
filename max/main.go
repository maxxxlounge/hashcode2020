package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pizzas map[int]int


func parseInputFromFile(filename string) (int, int,map[int]int, error){
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, nil, err
	}
	var maxSlice, pizzaTypes int
	defer file.Close()
	p := make (map[int]int)
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		row := scanner.Text()
		if i == 0 {
			pieces := strings.Split(row, " ")
			maxSlice, err = strconv.Atoi(pieces[0])
			if err != nil {
				return 0, 0, nil, err
			}
			pizzaTypes, err = strconv.Atoi(pieces[1])
			if err != nil {
				return 0, 0, nil, err
			}
		} else {
			slices := strings.Split(row, " ")
			ID := 0
			for _, slice := range slices{
				intSlice, err := strconv.Atoi(slice)
				if err != nil {
					return 0, 0, nil, err
				}
				p[ID]=intSlice
				ID++
			}
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, nil, err
	}

	return maxSlice, pizzaTypes, p, nil
}



func main(){
	filename := "e_also_big.in"
	var maxSlices,_ int
	var p map[int]int
	maxSlices,_,p,err := parseInputFromFile(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	var solution []int
	total := maxSlices
	outputType := 0
	for k,v := range p{
		if total - v <= 0 {
			continue
		}
		total -= v
		solution = append(solution,k)
		outputType++
	}

	filename = strings.Replace(filename,".in",".out",-1)
	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	_, err = outFile.WriteString(strconv.Itoa(outputType) + "\n")
	if err != nil {
		log.Fatal(err)
	}
	for _,pt := range solution {
		_, err = outFile.WriteString(strconv.Itoa(pt) + " ")
		if err != nil {
			log.Fatal(err)
		}
	}

}
