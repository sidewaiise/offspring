package genome

import (
	"fmt"
)

type Genome map[uint64]Gene

func (g *Genome) Gene(gene uint64) (*Gene, error) {
	if g == nil {
		return nil, fmt.Errorf("Genome not fulfilled yet. %d", gene)
	}

	got, ok := (*g)[gene]
	if !ok {
		return nil, fmt.Errorf("\nNo gene available with %d... %v\n", gene, got)
	}

	return &got, nil
}

func (g *Genome) Size() uint64 {
	return uint64(len(*g))
}
