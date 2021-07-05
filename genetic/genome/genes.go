package genome

import (
	"fmt"
	"math"
)

const (
	GeneByteSize = 8
)

func NewGenes(genome Genome) Genes {
	genes := make(Genes, 0)
	for i := 0; i < len(genome); i++ {
		if i%8 == 0 {
			genes = append(genes, 0)
		}
	}
	return genes
}

type Gene struct {
	Desc  string
	Score float32

	// Advanced usage
	Adjacency []uint64 // adjacent genes
	Exclusion []uint64 // genes that get excluded
	Priority  uint64
}

func (g *Gene) GetScore() (float32, error) {
	if g == nil {
		return 0, fmt.Errorf("Gene is nil")
	}
	return g.Score, nil
}

func (g *Gene) GetDesc() (string, error) {
	if g == nil {
		return "", fmt.Errorf("Gene is nil")
	}
	return g.Desc, nil
}

type Genes []byte

func (g *Genes) segment(gene uint64) *uint8 {
	if g == nil || len(*g) == 0 {
		panic("*Genes not initialized properly")
	}
	var segment = uint64(len(*g) - 1)
	if gene > 0 {
		segment -= uint64(math.Floor(float64(gene) / GeneByteSize))
	}
	return &(*g)[segment]
}

func (g *Genes) Set(gene uint64) error {
	if g == nil || len(*g) == 0 {
		return fmt.Errorf("Set(gene) : g == nil or (%v) length\n", *g)
	}

	segment := g.segment(gene)

	var flag uint8 = 1 << uint8(gene%GeneByteSize)

	*segment = *segment | flag

	return nil
}

func (g *Genes) Toggle(gene uint64) error {
	if g == nil || len(*g) == 0 {
		return fmt.Errorf("Toggle(gene) : g == nil or (%v) length\n", *g)
	}

	segment := g.segment(gene)

	var flag uint8 = 1 << uint8(gene%GeneByteSize)

	*segment = *segment ^ flag

	return nil
}

func (g *Genes) Has(gene uint64) bool {
	if g == nil || len(*g) == 0 {
		fmt.Printf("Has(gene) : g == nil or (%v) length\n", *g)
		return false
	}

	segment := g.segment(gene)

	var flag uint8 = 1 << uint8(gene%GeneByteSize)

	if has := *segment&flag != 0; !has {
		return false
	}

	return true
}

func (g *Genes) Read(genome *Genome) map[uint64]bool {
	result := make(map[uint64]bool, len(*g)*GeneByteSize)
	for gene := uint64(0); gene < genome.Size(); gene++ {
		if !g.Has(gene) {
			result[gene] = false
			continue
		}

		result[gene] = true
	}
	return result
}
