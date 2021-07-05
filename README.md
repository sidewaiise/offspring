# Golang Genetic algorithm library

I was playing around with genetic algorithms and wrote this library one night. I thought it was a neat application.

## Maintaining a Genome

We already have a seed genome called `HumanAnecdotalPhenome` which is really designed to identify general personality traits of an individual. I do have plans to introduce more meaningful genomes in the future, but for now this is what we are working with.

You can create your own phenome by creating a package similar to this:

```
package myAwesomeGenome

const (
    MyTraitORGene1 uint64 = iota
    MyTraitORGene2
    ...
)

const MyAwesomeGenomeOrWhatever = genetic.Genome{
    MyTraitORGene1: {
        Desc:  "Describe what it is in plain language",
        Score: <some_acceptable_score>
    },
    ...
}
```

Now you have a genome, the standard fitness function will perform it's duties by using `Score` to build a fitness rank for each chromosome.

## Running the algo

to run the algorith, you need to begin with a `Population{}` and at least 2 parents (you can introduce as many as you like if starting position is dynamic).

Here's an eample of how you might begin:

```
p := genetic.Population{
    genetic.NewParent(
        "001",
        genome.NewGenes(&myAwesomeGenome.MyAwesomeGenomeOrWhatever, []uint64{
            MyTraitORGene1,
            MyTraitORGene2,
            ...
        }
        &phenotypes.HumanAnecdotalPhenome,
    ),
    ...
}
```

## Run Tests

Try out

```
go clean -testcache && go test ./tests -v -maxN=2 -generations=1
```