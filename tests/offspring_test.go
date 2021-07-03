package main

import (
	"flag"
	"fmt"
	"testing"

	"github.com/sidewaiise/offspring/chromosomes"
	"github.com/sidewaiise/offspring/genes"
	"github.com/sidewaiise/offspring/genetic"
)

var (
	maxN        = flag.Int("maxN", 2, "Max number of children")
	generations = flag.Int("generations", 1, "Max number of generations")
)

func Test(t *testing.T) {
	flag.Parse()

	jamesGenes := genes.NewGenericHumanGenes(map[string]bool{
		"strong_muscles": true,
		"slow_muscles":   true,
		"strong_memory":  true,
		"distracted":     true,
		"focused":        true,
		"adaptive":       true,
		"introverted":    true,
		"healthy_haeart": true,
		"curiosity":      true,
		"nature_lover":   true,
		"urban_lover":    true,
		"scientific":     true,
		"creative":       true,
		"left_wing":      true,
		"capitalist":     true,
		"thrillseeker":   true,
		"explorer":       true,
		"good_looking":   true,
		"rambles":        true,
	})

	aymeGenes := genes.NewGenericHumanGenes(map[string]bool{
		"fast_muscles":   true,
		"weak_memory":    true,
		"distracted":     true,
		"rigid":          true,
		"extroverted":    true,
		"healthy_haeart": true,
		"curiosity":      true,
		"nature_lover":   true,
		"organised":      true,
		"thrillseeker":   true,
		"explorer":       true,
		"good_looking":   true,
	})

	p := genetic.Population{
		chromosomes.NewPhenotypeParent(
			"James",
			jamesGenes,
			&genes.GenericHumanGenome,
		),
		chromosomes.NewPhenotypeParent(
			"Ayme",
			aymeGenes,
			&genes.GenericHumanGenome,
		),
	}

	children := p.Reproduce(*generations, *maxN)

	fittest := children.GetFittest()

	fmt.Printf("Limmer Baby: %s\n\n", fittest.Label())
	for _, tr := range fittest.Traits() {
		fmt.Printf("\n--> %s", tr)
	}

	fmt.Printf("\n\nDue Jan 2022")
	fmt.Printf("\n\n---------------------\n\n")

	fmt.Print("All candidate summary:\n\n")

	for _, r := range children {
		fmt.Printf("Name: %s, Fitness: %f\n", r.Label(), r.Fitness())
	}

}
