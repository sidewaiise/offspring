package chromosomes

import "fmt"

type Gene struct {
	Desc  string
	Score float32
}

func (g *Gene) GetScore() float32 {
	if g == nil {
		panic("Gene is nil")
	}
	return g.Score
}

func (g *Gene) GetDesc() string {
	if g == nil {
		panic("Gene is nil")
	}
	return g.Desc
}

type Genes map[string]bool

type Genome map[string]Gene

func (g *Genome) Gene(name string) (*Gene, error) {
	if g == nil {
		return nil, fmt.Errorf("Genome not fulfilled yet. %s", name)
	}
	gene, ok := (*g)[name]
	if !ok {
		return nil, fmt.Errorf("No gene named %s", name)
	}
	return &gene, nil
}

type Chromosome interface {
	Label() string
	Genes() map[string]bool
	Traits() []string
	Offspring(spouse Chromosome, genome *Genome, i int) Chromosome

	// Genetic algo functions
	Fitness() float32
	Crossover() // Crossover between parents
	Mutate()
}
