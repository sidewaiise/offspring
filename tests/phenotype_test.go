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
	surprise = flag.Bool("surprise", false, "Show the surprise")
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

	var tcs = map[string]testCase{
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
		"500000 children, 1 generation": {
			children:    500000,
			generations: 1,
		},
	}

	if *surprise {
		tcs = map[string]testCase{
			"making a baby": {
				children:    2,
				generations: 1,
			},
		}
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
			fmt.Printf("Time taken: %d us\n", duration.Microseconds())

			if err != nil {
				tt.Fatalf("Reproduce error: %v", err)
			}

			fittest := children.Fittest()
			traits, err := fittest.Traits()
			if err != nil {
				tt.Fatal(err)
			}

			if *surprise {
				fmt.Printf("\nLimmer Baby: %s\n\n", fittest.Label())
				for _, tr := range traits {
					fmt.Printf("\n--> %s", tr)
				}

				fmt.Printf("\n\nDue Jan 2022")
				fmt.Printf("\n\n---------------------\n\n")

				fmt.Printf("All candidate summary:\n\n")

				for _, r := range children {
					fmt.Printf("Name: %s, Fitness: %f\n", r.Label(), r.Fitness())
				}
			}
		})
	}
}
