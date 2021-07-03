package main

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/sidewaiise/offspring/chromosomes"
	"github.com/sidewaiise/offspring/genes"
	"github.com/sidewaiise/offspring/genetic"
)

func Test(t *testing.T) {
	jamesGenes := genes.NewGenericHumanGenes(map[string]bool{
		"strong_muscles":       true,
		"strong_memory":        true,
		"distracted":           true,
		"focused":              true,
		"adaptive":             true,
		"introverted":          true,
		"gender_chromosome_XX": true,
		"good_looking":         true,
		"rambles":              true,
	})

	aymeGenes := genes.NewGenericHumanGenes(map[string]bool{
		"fast_muscles": true,
		"weak_memory":  true,
		"distracted":   true,
		"rigid":        true,
		"extroverted":  true,
		"explorer":     true,
		"good_looking": true,
	})

	p := genetic.Population{
		chromosomes.NewPhenotypeParent(
			"James",
			jamesGenes,
			&genes.GenericHuman,
		),
		chromosomes.NewPhenotypeParent(
			"Ayme",
			aymeGenes,
			&genes.GenericHuman,
		),
	}

	var result = make(genetic.Population, 0)
	children := p.Reproduce(1, int(math.Round(rand.Float64()*10)))
	for _, c := range children {
		result = append(result, c)
	}

	fittest := result.GetFittest()

	fmt.Printf("Limmer Baby: %s\n\n", fittest.Label())
	for _, tr := range fittest.Traits() {
		fmt.Printf("\n--> %s", tr)
	}

	fmt.Printf("\n\nDue Jan 2022")
	fmt.Printf("\n\n---------------------\n\n")

	fmt.Print("All candidate summary:\n\n")

	for _, r := range result {
		fmt.Printf("Name: %s, Fitness: %f\n", r.Label(), r.Fitness())
	}

}
