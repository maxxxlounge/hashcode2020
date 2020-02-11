package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pizza struct {
	ID int
	Slices int
}

type Order struct {
	Pizzas []*Pizza
	Best []*Pizza
}

func (o *Order) AddPizza(pizza *Pizza) {
	o.Pizzas = append(o.Pizzas, pizza)
}

func (o *Order) RemoveLastPizza() {
	o.Pizzas = o.Pizzas[:len(o.Pizzas)-1]
}

func (o *Order) MarkAsBest() {
	o.Best = nil
	o.Best = append(o.Pizzas[:0:0], o.Pizzas...)
}

func (o *Order) GetScore() int {
	score := 0
	for _, pizza := range o.Pizzas {
		score += pizza.Slices
	}
	return score
}

func (o *Order) GetOutput() string {
	output := ""
	for _, pizza := range o.Best {
		output += strconv.Itoa(pizza.Slices) + " "
	}
	return output
}

func parseInputFromFile(filename string) ([]*Pizza, int, error){
	file, err := os.Open(filename)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	var pizzas []*Pizza
	pax := 0
	for scanner.Scan() {
		row := scanner.Text()
		if i == 0 {
			pieces := strings.Split(row, " ")
			pax, err = strconv.Atoi(pieces[0])
			if err != nil {
				return nil, 0, err
			}
			size, err := strconv.Atoi(pieces[1])
			if err != nil {
				return nil, 0, err
			}
			pizzas = make([]*Pizza, 0, size)
		} else {
			slices := strings.Split(row, " ")
			ID := 0
			for _, slice := range slices{
				intSlice, err := strconv.Atoi(slice)
				if err != nil {
					return nil, 0, err
				}
				pizzas = append(pizzas, &Pizza{
					ID:     ID,
					Slices: intSlice,
				})
				ID++
			}
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		return nil, 0, err
	}

	return pizzas, pax,  nil
}

func main() {

	if len(os.Args) != 2 {
		log.Fatal("missing argument. Usage: go run main.go <filename>")
	}
	name := os.Args[1]
	// parse input
	pizzas, pax, err := parseInputFromFile(name + ".in")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Pax => %d", pax)
	//fmt.Println()
	for _, pizza := range pizzas {
		fmt.Printf("%+v", pizza)
	}

	// elaborate


	orders := make([]Order, 0)
	outer:
		for j:= 1; j<len(pizzas); j++ {
			fmt.Printf("%d/%d", j, len(pizzas))
			fmt.Println()
			o := Order{
				Pizzas: make([]*Pizza, 0),
			}
			i := len(pizzas) -j
			max := 0
			fmt.Println()
			for i >=0 {
				o.AddPizza(pizzas[i])
				if o.GetScore() == pax {
					max = o.GetScore()
					o.MarkAsBest()
					fmt.Println("GOOOOL")
					fmt.Printf("Score: %d (diff %d)", max, pax-max)
					orders = append(orders, o)
					break outer
				} else if o.GetScore() > pax {
					o.RemoveLastPizza()
					//fmt.Println("removed pizza")
				} else if o.GetScore() > max {
					//fmt.Println("new max")
					max = o.GetScore()
					o.MarkAsBest()
				}
				fmt.Printf("Score: %d (diff %d)", max, pax-max)
				fmt.Println()
				i--
			}
			orders = append(orders, o)
		}

	// choose output
	maxScore := 0
	bestIndex := -1
	for index, o := range orders {
		if o.GetScore() > maxScore {
			maxScore = o.GetScore()
			bestIndex = index
		}
	}
	fmt.Printf("Best score: %d (diff %d)", maxScore, pax-maxScore)
	fmt.Println()
	// write output
	bestOrder := orders[bestIndex]
	outFile, err := os.Create(name + ".out")
	if err != nil {
		log.Fatal(err)
	}
	_, err = outFile.WriteString(strconv.Itoa(len(bestOrder.Best)) + "\n")
	if err != nil {
		log.Fatal(err)
	}
	_, err = outFile.WriteString(bestOrder.GetOutput())
	if err != nil {
		log.Fatal(err)
	}

}
