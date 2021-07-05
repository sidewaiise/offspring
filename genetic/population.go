package genetic

import (
	"fmt"
	"sort"
)

// Population is a collection of chomosomes
type Population []Chromosomer

// Fittest gets the fittest instance of the population
func (p Population) Fittest() Chromosomer {
	sort.Slice(p, func(i int, j int) bool {
		return p[i].Fitness() > p[j].Fitness()
	})
	return p[0]
}

// Reproduce is the fittest Chromosomer over number of generations and n reproduction per parents
func (p Population) Reproduce(generations int, n int) (Population, error) {
	if len(p) < 2 {
		return nil, fmt.Errorf("Need at least 2 chromosomes in he population")
	}

	if generations == 0 {
		return p, nil
	}

	sort.Slice(p, func(i int, j int) bool {
		return p[i].Fitness() > p[j].Fitness()
	})

	newPop := make(Population, 0)
	for i := 0; i < n; i++ {
		child := p[0].Offspring(p[1], generations, i)
		newPop = append(newPop, child)
	}

	return newPop.Reproduce(generations-1, n)
}

// Traits is the set of traits for all members of the population
func (p Population) Traits() map[string][]string {
	traits := make(map[string][]string, 0)
	for _, member := range p {
		t, _ := member.Traits()
		traits[member.Label()] = t
	}
	return traits
}
