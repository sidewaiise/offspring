package chromosomes

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func NewPhenotypeParent(name string, genes Genes, genome *Genome) *Phenotype {
	h := Phenotype{
		Name:   name,
		genes:  genes,
		genome: genome,
	}

	return &h
}

func NewPhenotypeChild(parents []Chromosome, genome *Genome, i int) *Phenotype {
	h := Phenotype{
		Name:    parents[0].Label() + "_" + parents[1].Label() + "_child_" + fmt.Sprintf("%d", i),
		parents: parents,
		genome:  genome,
	}

	h.Crossover()
	h.Mutate()

	return &h
}

type Phenotype struct {
	Chromosome

	Name    string
	parents []Chromosome
	genome  *Genome
	genes   Genes
}

// Label gets the genes
func (h *Phenotype) Label() string {
	return h.Name
}

// Genes gets the genes
func (h *Phenotype) Genes() map[string]bool {
	return h.genes
}

// Traits gets the active gene traits
func (h *Phenotype) Traits() []string {
	var traits = make([]string, 0)
	for gene, active := range h.genes {
		if active {
			g, err := h.genome.Gene(gene)
			if err != nil {
				panic(err)
			}

			traits = append(traits, g.GetDesc())

		}
	}
	return traits
}

func (h *Phenotype) Offspring(spouse Chromosome, genome *Genome, i int) Chromosome {
	return NewPhenotypeChild([]Chromosome{h, spouse}, genome, i)
}

// Fitness adds up for each gene that is active
func (h *Phenotype) Fitness() float32 {
	var score float32 = 0
	for gene, active := range h.genes {
		if active {
			g, err := h.genome.Gene(gene)
			if err != err {
				panic(err)
			}
			score += g.Score
		}
	}
	return score
}

func (h *Phenotype) Crossover() {
	if len(h.parents) == 0 {
		fmt.Println("No parents to crossover with")
		return
	}

	// Get list of gene keys
	keys := make([]string, 0)
	h.genes = make(map[string]bool, 0)

	for key, _ := range *h.genome {
		keys = append(keys, key)
	}

	// Jumble them up
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	// Set random crossover point with a fluctuation range
	variationPercent := 0.05
	fluctuationRange := math.Round(float64(len(keys)) * variationPercent)
	min := int(math.RoundToEven(float64(len(keys))/2) - fluctuationRange)
	max := min + (int(fluctuationRange) * 2)
	crossoverPoint := rand.Intn(max-min) + min

	// Do the crossover for parents
	for i := 0; i < len(keys); i++ {
		gene := keys[i]
		if i < crossoverPoint {
			h.genes[gene] = h.parents[0].Genes()[gene]
			continue
		}

		h.genes[gene] = h.parents[1].Genes()[gene]
	}
}

func (h *Phenotype) Mutate() {
	probabilityConstant := 0.01
	for gene, active := range h.genes {
		if rand := math.Round(rand.Float64()*1000) / 1000; rand > probabilityConstant {
			continue
		}

		// Flip the gene if we hit our probability
		h.genes[gene] = !active
		g, err := h.genome.Gene(gene)
		if err != nil {
			panic(fmt.Sprintf("No gene found during mutation: ", gene))
		}

		if h.genes[gene] {
			fmt.Println("Mutation activated gene for " + g.GetDesc())
		} else {
			fmt.Println("Mutation deactivated gene for " + g.GetDesc())
		}
	}
}
