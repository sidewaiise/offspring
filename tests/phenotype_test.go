package main

import (
	"flag"
	"fmt"
	"testing"

	"github.com/sidewaiise/offspring/genes/phenotypes"
	"github.com/sidewaiise/offspring/genetic"
)

var (
	maxN        = flag.Int("maxN", 2, "Max number of children")
	generations = flag.Int("generations", 1, "Max number of generations")
)

func TestPhenotypes(t *testing.T) {
	flag.Parse()

	jamesGenes := phenotypes.NewHumanTraits([]uint64{
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

	aymeGenes := phenotypes.NewHumanTraits([]uint64{
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

	t.Run("Making a baby", func(tt *testing.T) {
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

		children, err := p.Reproduce(*generations, *maxN)

		if err != nil {
			tt.Fatalf("Reproduce error: %v", err)
		}

		fittest := children.Fittest()

		fmt.Printf("Limmer Baby: %s\n\n", fittest.Label())
		traits, err := fittest.Traits()
		if err != nil {
			tt.Fatal(err)
		}

		for _, tr := range traits {
			fmt.Printf("\n--> %s", tr)
		}

		fmt.Printf("\n\nDue Jan 2022")
		fmt.Printf("\n\n---------------------\n\n")

		fmt.Print("All candidate summary:\n\n")

		for _, r := range children {
			fmt.Printf("Name: %s, Fitness: %f\n", r.Label(), r.Fitness())
		}
	})

}
