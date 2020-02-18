package genetic

import "math/rand"

type Gene struct {
	Value  int
	Index  int
	Active bool
}

type Organism struct {
	DNA     []Gene
	Fitness int
}

type Population []Organism

func (o *Organism) CalculateFitness(target int) {
	var fitness int
	for _, i := range o.DNA {
		if i.Active {
			fitness += i.Value
			if fitness > target {
				fitness = 0
				break
			}
		}
	}
	o.Fitness = fitness
}

func (o *Organism) Mutate() {
	for _, g := range o.DNA {
		r := rand.Intn(100)
		if r > 95 {
			g.Active = !g.Active
		}
	}
}

func CreateOrganism(dna []int) Organism {
	o := Organism{}
	for i, d := range dna {
		o.DNA[i] = Gene{
			Value:  d,
			Index:  i,
			Active: (rand.Int()%2 == 0),
		}
	}
	return o
}

func createPool(population Population, target int, maxFitness int) {
	pool := []Organism{}
	// create a pool for next generation
	for i := 0; i < len(population); i++ {
		population[i].CalculateFitness(target)
		num := int((population[i].Fitness / maxFitness) * 100)
		for n := 0; n < num; n++ {
			pool = append(pool, population[i])
		}
	}
	return
}

// perform natural selection to create the next generation
func NaturalSelection(pool Population, population Population, target int) Population {
	next := Population{}

	for i := 0; i < len(population); i++ {
		r1, r2 := rand.Intn(len(pool)), rand.Intn(len(pool))
		a := pool[r1]
		b := pool[r2]

		child := Generate(a, b)
		child.Mutate()
		child.CalculateFitness(target)
		next[i] = child
	}
	return next
}

func Generate(om Organism, of Organism) Organism {
	o := Organism{}
	for i, g := range om.DNA {
		active := g.Active
		if i%2 == 0 {
			active = of.DNA[i].Active
		}
		o.DNA[i] = Gene{
			Value:  g.Index,
			Index:  i,
			Active: active,
		}
	}
	return o
}

func (pp Population) GetBest() Organism {
	f := 0
	var o Organism
	for _, p := range pp {
		if p.Fitness > f {
			o = p
			f = p.Fitness
		}
	}
	return o
}

func GeneratePopulation(target, organismCount int, ss []int) Population {
	p := Population{}
	for i := 0; i < organismCount; i++ {
		o := CreateOrganism(ss)
		p = append(p, o)
		o.CalculateFitness(target)
	}
	return p
}
