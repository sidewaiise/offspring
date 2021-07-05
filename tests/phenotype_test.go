package main

import (
	"flag"
	"fmt"
	"testing"
	"time"

	"github.com/sidewaiise/offspring/genes/phenotypes"
	"github.com/sidewaiise/offspring/genetic"
	"github.com/sidewaiise/offspring/genetic/genome"
)

var (
	maxN        = flag.Int("maxN", 2, "Max number of children")
	generations = flag.Int("generations", 1, "Max number of generations")
)

func TestPhenotypes(t *testing.T) {
	flag.Parse()

	jamesGenes := genome.NewGenes(&phenotypes.HumanAnecdotalPhenome, []uint64{
		phenotypes.StrongMusclesTrait,
		phenotypes.StrongMemoryTrait,
		phenotypes.DistractedTrait,
		phenotypes.FocusedTrait,
		phenotypes.AdaptiveTrait,
		phenotypes.IntrovertedTrait,
		phenotypes.GoodHeartTrait,
		phenotypes.CuriousTrait,
		phenotypes.NatureLoverTrait,
		phenotypes.CityLoverTrait,
		phenotypes.ScientificTrait,
		phenotypes.CreativityTrait,
		phenotypes.LeftWingPoliticsTrait,
		phenotypes.CapitalistTrait,
		phenotypes.ThrillseekerTrait,
		phenotypes.ExplorerTrait,
		phenotypes.GoodLookingTrait,
		phenotypes.RamblingSpeechTrait,
	})

	aymeGenes := genome.NewGenes(&phenotypes.HumanAnecdotalPhenome, []uint64{
		phenotypes.FastTwitchMusclesTrait,
		phenotypes.WeakMemoryTrait,
		phenotypes.StrongMemoryTrait,
		phenotypes.DistractedTrait,
		phenotypes.UnadaptiveTrait,
		phenotypes.ExtrovertedTrait,
		phenotypes.GoodHeartTrait,
		phenotypes.CuriousTrait,
		phenotypes.NatureLoverTrait,
		phenotypes.OrganisedTrait,
		phenotypes.ThrillseekerTrait,
		phenotypes.ExplorerTrait,
		phenotypes.GoodLookingTrait,
	})

	type testCase struct {
		children    int
		generations int
	}

	tcs := map[string]testCase{
		"making a baby": {
			children:    *maxN,
			generations: *generations,
		},
		"2 children 1 generation": {
			children:    2,
			generations: 1,
		},
		"5 children, 1 generation": {
			children:    5,
			generations: 1,
		},
		"50 children, 1 generation": {
			children:    50,
			generations: 1,
		},
		"500 children, 1 generation": {
			children:    500,
			generations: 1,
		},
		"2 children, 5 generation": {
			children:    2,
			generations: 5,
		},
		"2 children, 50 generation": {
			children:    2,
			generations: 50,
		},
		"2 children, 500 generation": {
			children:    2,
			generations: 500,
		},
		"5 children, 5 generation": {
			children:    5,
			generations: 5,
		},
		"50 children, 50 generation": {
			children:    50,
			generations: 50,
		},
		"100 children, 100 generation": {
			children:    100,
			generations: 100,
		},
	}

	for name, test := range tcs {
		t.Run(name, func(tt *testing.T) {
			start := time.Now()
			p := genetic.Population{
				genetic.NewParent(
					"James",
					jamesGenes,
					&phenotypes.HumanAnecdotalPhenome,
				),
				genetic.NewParent(
					"Ayme",
					aymeGenes,
					&phenotypes.HumanAnecdotalPhenome,
				),
			}

			children, err := p.Reproduce(test.generations, test.children)
			duration := time.Since(start)

			fmt.Println("---------------------------")
			if duration.Microseconds() <= 1 {
				fmt.Printf("Time taken: %d ns\n", duration.Nanoseconds())
			} else {
				fmt.Printf("Time taken: %d ms\n", duration.Milliseconds())
			}

			if err != nil {
				tt.Fatalf("Reproduce error: %v", err)
			}

			fittest := children.Fittest()

			fmt.Printf("\nLimmer Baby: %s\n\n", fittest.Label())
			traits, err := fittest.Traits()
			if err != nil {
				tt.Fatal(err)
			}

			for _, tr := range traits {
				fmt.Printf("\n--> %s", tr)
			}

			fmt.Printf("\n\nDue Jan 2022")
			fmt.Printf("\n\n---------------------\n\n")

			fmt.Printf("All candidate summary:\n\n")

			for _, r := range children {
				fmt.Printf("Name: %s, Fitness: %f\n", r.Label(), r.Fitness())
			}
		})
	}
}
