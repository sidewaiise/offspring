package genetic

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/sidewaiise/offspring/genes/phenotypes"
	"github.com/sidewaiise/offspring/genetic/genome"
)

type Chromosomer interface {
	Label() string
	Genes() *genome.Genes
	Traits() ([]string, error)

	// Genetic algo functions
	Fitness() float32
	Crossover() // Crossover between parents
	Mutate() error
	Offspring(spouse Chromosomer, g int, i int) Chromosomer
}

func NewParent(name string, genes genome.Genes, genome *genome.Genome) *Chromosome {
	h := Chromosome{
		Name:   name,
		genes:  genes,
		genome: genome,
	}

	return &h
}

func NewChild(parents []Chromosomer, genome *genome.Genome, g int, i int) Chromosomer {
	h := Chromosome{
		Name:    fmt.Sprintf("gen_%d_child_%d", g, i),
		parents: parents,
		genome:  genome,
	}

	h.Crossover()
	h.Mutate()

	return &h
}

type Chromosome struct {
	Name    string
	parents []Chromosomer
	genome  *genome.Genome
	genes   genome.Genes
}

// Label gets the genes
func (h *Chromosome) Label() string {
	return h.Name
}

// Genes gets the genes
func (h *Chromosome) Genes() *genome.Genes {
	return &h.genes
}

// Traits gets the active gene traits
func (h *Chromosome) Traits() ([]string, error) {
	traits := make([]string, 0)
	read := h.genes.Read(h.genome)
	for i, b := range h.genes {
		fmt.Printf("> Chromosome[%d] : % 08b\n", i, b)
	}

	for gene, active := range read {
		if active {
			g, err := h.genome.Gene(gene)
			if err != nil {
				return nil, err
			}

			d, err := g.GetDesc()
			if err != nil {
				return nil, err
			}

			traits = append(traits, d)
		}
	}
	return traits, nil
}

func (h *Chromosome) Offspring(spouse Chromosomer, g int, i int) Chromosomer {
	return NewChild([]Chromosomer{h, spouse}, h.genome, g, i)
}

// Fitness adds up for each gene that is active
func (h *Chromosome) Fitness() float32 {
	var score float32 = 0
	for gene, active := range h.genes.Read(h.genome) {
		if !active {
			continue
		}

		g, _ := h.genome.Gene(gene)
		s, _ := g.GetScore()

		score += s
	}
	return score
}

// Crossover randomises the order of genes then adds randomness to the crossover point.
func (h *Chromosome) Crossover() {
	if len(h.parents) == 0 {
		fmt.Println("No parents to crossover with")
		return
	}

	// Get list of gene keys
	h.genes = genome.NewGenes(&phenotypes.HumanAnecdotalPhenome, []uint64{})
	mixer := make([]uint64, len(*h.genome))
	for gene := range *h.genome {
		mixer = append(mixer, gene)
	}

	// Jumble them up
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(mixer), func(i, j int) {
		mixer[i], mixer[j] = mixer[j], mixer[i]
	})

	// Set random crossover point with a fluctuation range
	variationPercent := 0.1
	fluctuationRange := math.Round(float64(len(mixer)) * variationPercent)
	min := int(math.RoundToEven(float64(len(mixer))/2) - fluctuationRange)
	max := min + (int(fluctuationRange) * 2)
	crossoverPoint := rand.Intn(max-min) + min

	// Do the crossover for parents
	for i := 0; i < len(mixer); i++ {
		gene := mixer[i]
		if i < crossoverPoint {
			if h.parents[0].Genes().Has(gene) {
				h.genes.Set(gene)
			}
			continue
		}

		if h.parents[1].Genes().Has(gene) {
			h.genes.Set(gene)
		}
	}
}

func (h *Chromosome) Mutate() error {
	probabilityConstant := 0.01
	for gene := range h.genes.Read(h.genome) {
		if rand := math.Round(rand.Float64()*1000) / 1000; rand > probabilityConstant {
			continue
		}

		// Flip the gene if we hit our probability
		err := h.genes.Toggle(gene)
		if err != nil {
			return err
		}

		mutationAction := "activated"
		if !h.genes.Has(gene) {
			mutationAction = "deactivated"
		}

		g, err := h.genome.Gene(gene)
		if err != nil {
			return err
		}

		desc, err := g.GetDesc()
		if err != nil {
			return err
		}
		fmt.Printf("> [%s] Mutation %s gene for %s\n", h.Label(), mutationAction, desc)
	}
	return nil
}
