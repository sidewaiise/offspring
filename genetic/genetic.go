package genetic

import (
	"fmt"
	"sort"

	"github.com/sidewaiise/offspring/chromosomes"
)

// Population is a collection
type Population []chromosomes.Chromosome

func (p Population) GetFittest() chromosomes.Chromosome {
	sort.Slice(p, func(i int, j int) bool {
		return p[i].Fitness() > p[j].Fitness()
	})
	return p[0]
}

// Reproduce is the fittest Chromosome over number of generations and n reproduction per parents
func (p Population) Reproduce(generations int, n int) Population {
	if len(p) < 2 {
		fmt.Errorf("Need at least 2 chromosomes in he population")
		return nil
	}

	if generations == 0 {
		return p
	}

	// Sort so fittest are in p[0] and p[1]
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

func (p Population) Traits() map[string][]string {
	traits := make(map[string][]string, 0)
	for _, member := range p {
		traits[member.Label()] = member.Traits()
	}
	return traits
}
